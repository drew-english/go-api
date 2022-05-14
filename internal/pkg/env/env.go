package env

import "os"

func GetEnvDefault(key, fallback string) string {
	if envVar, ok := os.LookupEnv(key); ok {
		return envVar
	}

	return fallback
}
