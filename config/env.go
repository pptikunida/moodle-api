package config

import (
	"os"
)

func GetEnv(key, defaultVal string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal
	}

	return value
}
