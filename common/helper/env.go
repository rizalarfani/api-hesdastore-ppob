package helper

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Errorf("missing required env: %s", key))
	}

	return v
}

func GetEnvInt(key string) int {
	v := GetEnv(key)
	i, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Errorf("invalid int for %s: %q", key, v))
	}
	return i
}
