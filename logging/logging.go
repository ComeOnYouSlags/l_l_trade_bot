package logging

import (
	"fmt"
	"os"
	"path/filepath"
	h "tradebot/helpers"
)

func CreateLogFile(logDir string) (string, error) {
	filename := h.GetTimestamp() + ".log"
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
