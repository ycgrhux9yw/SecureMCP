package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	ServerPort    int
	ScanTimeout   int
	ReportDir     string
	LogLevel      string
	EnableMetrics bool
}

// Load loads the configuration from environment variables.
// Defaults have been adjusted for local development use:
//   - SERVER_PORT: 9090 (avoid conflict with common local services on 8080)
//   - SCAN_TIMEOUT: 60s (give scans more time to complete)
//   - LOG_LEVEL: debug (more verbose output during development)
func Load() *Config {
	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "9090"))
	timeout, _ := strconv.Atoi(getEnv("SCAN_TIMEOUT", "60"))
	enableMetrics, _ := strconv.ParseBool(getEnv("ENABLE_METRICS", "false"))

	return &Config{
		ServerPort:    port,
		ScanTimeout:   timeout,
		ReportDir:     getEnv("REPORT_DIR", "reports"),
		LogLevel:      getEnv("LOG_LEVEL", "debug"),
		EnableMetrics: enableMetrics,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
