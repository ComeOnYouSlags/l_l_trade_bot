package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GetCurrentDateString() string {
	return time.Now().Format("1970-01-01")
}

func GetCurrentTimeString() string {
	return time.Now().Format("23:15:01")
}

func CreateLogFile(logDir string) (string, error) {
	filename := GetCurrentDateString() + GetCurrentTimeString() + ".log"
	fullPath := filepath.Join(logDir, filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	// Create the file
	file, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Verify file was created successfully
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return "", fmt.Errorf("file was not created successfully")
	}

	return fullPath, nil
}
