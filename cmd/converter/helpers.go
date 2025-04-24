package main
import (
    "os"
)
// WriteFile writes content to a file
func WriteFile(path, content string) error {
    return os.WriteFile(path, []byte(content), 0644)
}
