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
		Port:      getEnv("APP_PORT", "6272"),
		JwtSecret: getEnv("JWT_SECRET", "secret"),
		Dsn:       getEnv("DB_DSN", "mongodb+srv://admin:1@cluster0.re2oqf1.mongodb.net/?appName=Cluster0"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
