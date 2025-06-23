package lib

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/kayceeDev/caspa-events/pkg/config"
)

// SentryConfig holds configuration for Sentry initialization
type SentryConfig struct {
	DSN              string
	Environment      string
	Release          string
	ServerName       string
	Debug            bool
	TracesSampleRate float64
	AttachStacktrace bool
}

// NewDefaultSentryConfig creates a default Sentry configuration
func NewDefaultSentryConfig(dsn string) SentryConfig {
	env := config.EnvVars()

	return SentryConfig{
		DSN:              dsn,
		Environment:      env.APP_ENV,
		Release:          env.APP_VERSION,
		ServerName:       "goldnlilies_server",
		Debug:            env.APP_ENV != "production",
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
	}
}

func InitializeSentry(config SentryConfig) error {
	if config.Environment == "development" {
		return nil
	}

	// Validate required fields
	if config.DSN == "" {
		fmt.Println("Sentry DSN not provided, skipping initialization")
		return nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.DSN,
		Environment:      config.Environment,
		Release:          config.Release,
		ServerName:       config.ServerName,
		Debug:            config.Debug,
		TracesSampleRate: config.TracesSampleRate,
		AttachStacktrace: config.AttachStacktrace,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if event == nil {
				return nil
			}
			// Add runtime information
			event.Tags["go_version"] = runtime.Version()
			event.Tags["num_goroutines"] = fmt.Sprintf("%d", runtime.NumGoroutine())

			// Add request information if available
			if hint.Context != nil && hint != nil {
				if c, ok := hint.Context.(*gin.Context); ok {
					event.Request = extractRequestData(c.Request)
				}
			}
			return event
		},
		EnableTracing: true,
	})

	if err != nil {
		log.Printf("Sentry initialization failed: %v\n", err)
		return err
	}

	return nil
}

func CaptureError(err error, c *gin.Context, tags map[string]string, args ...interface{}) {
	env := config.EnvVars()

	if env.APP_ENV == "local" || err == nil {
		return
	}

	hub := sentry.CurrentHub().Clone()
	hub.ConfigureScope(func(scope *sentry.Scope) {
		if c != nil && c.Request != nil {
			// Set the raw request
			scope.SetRequest(c.Request)
			scope.SetUser(extractUserData(c))

			// Get request body from context
			requestContext := make(map[string]interface{})

			// Add basic request info
			requestContext["url"] = c.Request.URL.String()
			requestContext["method"] = c.Request.Method
			requestContext["path"] = c.FullPath()

			// Add query parameters
			if len(c.Request.URL.RawQuery) > 0 {
				requestContext["query_params"] = c.Request.URL.Query()
			}

			// Add request body if available
			if bodyData, exists := c.Get(RequestBodyKey); exists {
				requestContext["body"] = bodyData
			}

			// Add headers (excluding sensitive data)
			headers := make(map[string]string)
			for k, v := range c.Request.Header {
				if !isSensitiveHeader(k) && len(v) > 0 {
					headers[k] = v[0]
				}
			}
			requestContext["headers"] = headers

			scope.SetContext("request", requestContext)
		}

		// Add custom tags
		for k, v := range tags {
			scope.SetTag(k, v)
		}

		// Add performance metrics
		addPerformanceMetrics(scope)
	})

	hub.CaptureException(err)
}

func isSensitiveHeader(header string) bool {
	sensitiveHeaders := map[string]bool{
		"Authorization":    true,
		"Cookie":           true,
		"Set-Cookie":       true,
		"X-CSRF-Token":     true,
		"X-Authentication": true,
		"X-API-Key":        true,
		"X-Access-Token":   true,
		"Private-Token":    true,
		"Session-Token":    true,
		"Refresh-Token":    true,
	}
	return sensitiveHeaders[header]
}

func addPerformanceMetrics(scope *sentry.Scope) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	scope.SetExtra("performance_metrics", map[string]interface{}{
		"allocated_memory_mb": m.Alloc / 1024 / 1024,
		"total_allocated_mb":  m.TotalAlloc / 1024 / 1024,
		"system_memory_mb":    m.Sys / 1024 / 1024,
		"gc_cycles":           m.NumGC,
		"goroutines":          runtime.NumGoroutine(),
	})
}

func extractRequestData(r *http.Request) *sentry.Request {
	if r == nil {
		return nil
	}

	headers := make(map[string]string)
	for k, v := range r.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	data := make(map[string]string)

	// Add query parameters
	for k, v := range r.URL.Query() {
		if len(v) > 0 {
			data[fmt.Sprintf("query_%s", k)] = v[0]
		}
	}

	return &sentry.Request{
		URL:         r.URL.String(),
		Method:      r.Method,
		Headers:     headers,
		Cookies:     r.Header.Get("Cookie"),
		QueryString: r.URL.RawQuery,
	}
}

func extractUserData(c *gin.Context) sentry.User {
	user := sentry.User{
		IPAddress: c.ClientIP(),
	}

	// Add user ID if available from your auth context
	if uid, exists := c.Get("user_id"); exists {
		user.ID = fmt.Sprint(uid)
	}

	return user
}
