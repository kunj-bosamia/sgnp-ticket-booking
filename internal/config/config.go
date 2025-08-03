package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port           string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPass         string
	DBName         string
	JWTSecret      string
	DBPoolMaxConns int
	DBPoolMinConns int
	VerifyToken    string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		Port:           getEnv("PORT", "8080"),
		DBHost:         getEnv("DB_HOST", "localhost"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPass:         getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", "sgnp"),
		JWTSecret:      getEnv("JWT_SECRET", "devsecret"),
		DBPoolMaxConns: getEnvAsInt("DB_POOL_MAX_CONNS", 20),
		DBPoolMinConns: getEnvAsInt("DB_POOL_MIN_CONNS", 5),
		VerifyToken:    mustGetEnv("VERIFY_TOKEN"),
	}
}

func getEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Env %s not found. Using default: %s", key, defaultVal)
		return defaultVal
	}
	return val
}

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return val
}

func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		log.Printf("Env %s not found. Using default: %d", key, defaultVal)
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		log.Printf("Invalid int for %s: %v. Using default: %d", key, err, defaultVal)
		return defaultVal
	}
	return val
}
