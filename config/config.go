package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Server   string
	Port     int
	Username string
	Password string
	KeyPath  string
}

func NewServer() *ServerConfig {
	return &ServerConfig{
		Server:   getEnv("SERVER", ""),
		Port:     getPortEnv("PORT", 22),
		Username: getEnv("USERNAME", "root"),
		KeyPath:  getEnv("KEY_PATH", "/Users/bujikuh/.ssh/id_rsa"),
		Password: getEnv("PASSWORD", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getPortEnv(port string, defaultVal int) int {
	valueStr := getEnv(port, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func NewDataBase() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     getDbEnv("PG_HOST", ""),
		Port:     getDbPortEnv("PG_PORT", 5432),
		Username: getDbEnv("PG_USER", ""),
		Password: getDbEnv("PG_PASS", ""),
		DbName:   getDbEnv("PG_DB_NAME", ""),
	}
}

func getDbEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getDbPortEnv(port string, defaultVal int) int {
	valueStr := getEnv(port, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}