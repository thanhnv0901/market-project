package utils

import (
	"path"
	"runtime"
)

// GetDirCurrent ..
func GetDirCurrent() (string, int, bool) {
	// Get current directory path
	_, filename, line, isOk := runtime.Caller(1)
	pathDir := path.Dir(filename)

	return pathDir, line, isOk
}
