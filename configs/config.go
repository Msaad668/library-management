// configs/config.go
package configs

import (
	"fmt"
	"os"
)

// Config holds the configuration values
type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
}

// GetConfig reads environment variables and returns a Config struct
func GetConfig() *Config {
    return &Config{
        DBHost:     getEnvironmentVar("DB_HOST", "localhost"),
        DBUser:     getEnvironmentVar("DB_USER", "postgres"),
        DBPassword: getEnvironmentVar("DB_PASSWORD", "react123"),
        DBName:     getEnvironmentVar("DB_NAME", "library_db"),
        DBPort:     getEnvironmentVar("DB_PORT", "5432"),
    }
}

func getEnvironmentVar(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// GetDSN constructs the Data Source Name for PostgreSQL connection
func (c *Config) GetDSN() string {
    return fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
    )
}