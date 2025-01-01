package data

import "os"

// Get if path exists
func PathExists(path string) bool {
	_, error := os.Stat(path)
  return !os.IsNotExist(error)
}

