package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadConfig загружает конфигурацию из файла
func LoadConfig(filename string, config interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(config)
}
