package utils

import (
	"os"
)

// FileExists check if path exists
func FileExists(path string) bool {
	exists := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exists = false
	}

	return exists
}
