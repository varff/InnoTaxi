package helper

import "os"

func GetEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		if defaultValue == "" {
			panic("Environment variable isn't set")
		}
		return defaultValue
	}
	return value
}
