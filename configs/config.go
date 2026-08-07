package configs

import "os"

type Config struct {
	AppName   string
	Port      string
	JwtSecret string
	Dsn       string
}

func LoadConfig() Config {
	return Config{
		AppName:   getEnv("APP_NAME", "queue-kiosk"),
		Port:      getEnv("APP_PORT", "8080"),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
		Dsn:       getEnv("DB_DSN", "sqlite://queue-kiosk.db"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
