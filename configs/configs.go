package configs

import (
	"fmt"
	"os"
)

const (
	hostEnvVar    = "Host"
	connEnvVar    = "Conn"
	migrateEnvVar = "Migrate"
)

type Config struct {
	Host    string
	Conn    string
	Migrate string
}

func NewConfig() (*Config, error) {
	err := lookUp(hostEnvVar)
	err = lookUp(connEnvVar)
	err = lookUp(migrateEnvVar)
	return &Config{
		Host:    os.Getenv(hostEnvVar),
		Conn:    os.Getenv(connEnvVar),
		Migrate: os.Getenv(migrateEnvVar),
	}, err
}

func lookUp(key string) error {
	_, ok := os.LookupEnv(key)
	if !ok {
		return fmt.Errorf("no such key in .env: %s", key)
	}
	return nil
}
