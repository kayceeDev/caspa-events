package domains

type Config struct {
	GIN_MODE           string `env:"GIN_MODE" required:"true"`
	APP_ENV            string `env:"APP_ENV" required:"true"`
	APP_VERSION        string `env:"APP_VERSION" required:"true"`
	CLIENT_URL         string `env:"CLIENT_URL" required:"true"`
	PORT               string `env:"PORT" required:"true"`
	APP_ORIGINS        string `env:"APP_ORIGINS" required:"true"`
	APP_ORIGIN_REGEX   string `env:"APP_ORIGIN_REGEX" required:"true"`
	USE_REGEX_CORS     string `env:"USE_REGEX_CORS" required:"true"`
	CONNECTION_TIMEOUT string `env:"CONNECTION_TIMEOUT" required:"true"`
	STAGING_API_HOST   string `env:"STAGING_API_HOST"`

	DB_HOST     string `env:"DB_HOST" required:"true"`
	DB_PORT     string `env:"DB_PORT" required:"true"`
	DB_USER     string `env:"DB_USER" required:"true"`
	DB_PASSWORD string `env:"DB_PASSWORD" required:"true"`
	DB_NAME     string `env:"DB_NAME" required:"true"`
	DB_SSLMODE  string `env:"DB_SSLMODE" required:"true"`

	REDIS_ADDR     string `env:"REDIS_ADDR" required:"true"`
	REDIS_PASSWORD string `env:"REDIS_PASSWORD" required:"true"`
	REDIS_USERNAME string `env:"REDIS_USERNAME" required:"true"`
	REDIS_TLS      string `env:"REDIS_TLS" required:"true"`

	JWT_SECRET        string `env:"JWT_SECRET" required:"true"`
	ACCESS_TOKEN_TTL  string `env:"ACCESS_TOKEN_TTL" required:"true"`
	REFRESH_TOKEN_TTL string `env:"REFRESH_TOKEN_TTL" required:"true"`




	EMAIL_TEMPLATE_BASE_URL     string `env:"EMAIL_TEMPLATE_BASE_URL" required:"true"`
	EMAIL_TEMPLATE_ACCESS_TOKEN string `env:"EMAIL_TEMPLATE_ACCESS_TOKEN" required:"true"`

}
