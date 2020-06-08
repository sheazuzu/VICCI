package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	ErrNoEnvFile   = errors.New("no .env file in the provided directory")
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func loadEnv(folder string) error {

	defaultEnvFile := fmt.Sprintf("%s.env", folder)
	if err := loadConfigFile(defaultEnvFile); err != nil {
		return err
	}
	if !fileExists(defaultEnvFile) {
		return ErrNoEnvFile
	}
	return nil
}

func loadConfigFile(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil && fileExists(fileName) {
		return err
	}
	return nil
}


func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
