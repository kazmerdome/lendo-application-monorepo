package config

import "syscall"

// GetEnvString gets the environment variable by key.
// if the key does not exist, it returns the fallback value.
func GetEnvString(key, fallback string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}

	return fallback
}
