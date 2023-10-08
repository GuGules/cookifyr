package utils

import (
	"fmt"
	"os"
)

func GetEnvOr(key string, fallback string) string {
	value, found := os.LookupEnv(key)
	if found {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) (string, error) {
	value, found := os.LookupEnv(key)
	if found {
		return value, nil
	}
	return "", fmt.Errorf("env variable %s not found", key)
}
