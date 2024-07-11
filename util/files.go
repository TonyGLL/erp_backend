package util

import "os"

func SaveToFile(filename string, file []byte) error {
	return os.WriteFile(filename, file, 0666)
}
