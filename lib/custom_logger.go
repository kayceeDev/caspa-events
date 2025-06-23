package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomLoggerInterface interface {
	Info(...interface{})
	Warn(...interface{})
	Debug(...interface{})
	Error(...interface{})
	Panic(...interface{})
}

type CustomLogger struct {
	info  *log.Logger
	warn  *log.Logger
	debug *log.Logger
	error *log.Logger
	panic *log.Logger
}

func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		info:  slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo),
		warn:  slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelWarn),
		debug: slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelDebug),
		error: slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, nil), slog.LevelError),
		panic: slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, nil), slog.LevelError),
	}
}

func (cl *CustomLogger) Info(args ...interface{}) {
	fmt.Println("\n")
	cl.info.Println(args...)

}

func (cl *CustomLogger) Warn(args ...interface{}) {
	fmt.Println("\n")
	cl.warn.Println(args...)

}

func (cl *CustomLogger) Debug(args ...interface{}) {
	fmt.Println("\n")
	cl.debug.Println(args...)

}

func (cl *CustomLogger) Error(args ...interface{}) {

	fmt.Println("\n")
	cl.error.Println(args...)

	tags, err := getTags(args...)

	// Capture the error with Sentry
	if c, exists := args[0].(*gin.Context); exists {
		CaptureError(err, c, tags, args...)
	} else {
		CaptureError(err, nil, tags, args...)
	}

}
func (cl *CustomLogger) Panic(args ...interface{}) {

	// Log the panic
	fmt.Println("\n")
	cl.panic.Println(args...)
	tags, err := getTags(args...)

	// Check if we have a request body in the args
	if c, exists := args[0].(*gin.Context); exists {
		CaptureError(err, c, tags, args...)
	} else {
		CaptureError(err, nil, tags, args...)
	}

	// Now panic with the original arguments
	fmt.Println("\n")
	cl.panic.Panicln(args...)

}

func getTags(args ...interface{}) (map[string]string, error) {
	tags := make(map[string]string)
	var err error
	var message string

	// Extract message and error from args
	for _, arg := range args {
		switch v := arg.(type) {
		case error:
			err = v
			// Convert error to string value
			tags["error"] = v.Error()
		case string:
			message = v
			tags["message"] = v
		case *gin.Context:
			// Skip gin.Context as it's handled separately
			continue
		default:
			// For other types, try JSON marshal first
			if bodyBytes, err := json.Marshal(v); err == nil {
				tags["value"] = string(bodyBytes)
			} else {
				// Fallback to detailed string representation
				tags["value"] = fmt.Sprintf("%+v", v)
			}

		}
	}

	// If no error was found but we have a message, create an error from the message
	if err == nil && message != "" {
		err = errors.New(message)
	}

	// Clean up any pointer references and ensure valid JSON strings
	for k, v := range tags {
		if strings.Contains(v, "0x") {
			// Try to get the underlying value
			cleanValue := strings.ReplaceAll(v, "0x", "ptr_")
			tags[k] = cleanValue
		}

		// Ensure the value is valid JSON string if it looks like JSON
		if strings.HasPrefix(v, "{") || strings.HasPrefix(v, "[") {
			var js interface{}
			if json.Unmarshal([]byte(v), &js) == nil {
				if prettyJSON, err := json.MarshalIndent(js, "", "  "); err == nil {
					tags[k] = string(prettyJSON)
				}
			}
		}
	}

	return tags, err
}
