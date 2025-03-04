package modules

import (
	"fmt"
	"os"
	"path/filepath"
)

// Import загружает внешний модуль
func Import(moduleName string) (interface{}, error) {
	path := filepath.Join("modules", moduleName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("module %s not found", moduleName)
	}
	// Логика загрузки модуля
	return nil, nil
}
