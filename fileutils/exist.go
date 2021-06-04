package fileutils

import "os"

// IsExist returns a boolean indicating whether a file or directory exist.
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}
