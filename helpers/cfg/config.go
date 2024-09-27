package cfg

import (
	"fmt"
	"os"
)

type Config map[string]string

func (m Config) Get(key string) (string, error) {
	value, exists := m[key]
	if exists {
		return value, nil
	}

	return "", fmt.Errorf("[Error]: Key %s does not exist", key)
}

func (m Config) GetEnv(key string) string {
	val := os.Getenv(key)
	if val != "" {
		return val
	}
	val, err := m.Get(key)
	if err != nil {
		return ""
	}

	return val
}
