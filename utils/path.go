package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func GetDesktopPath() (string, error) {
	var desktopPath string
	homeDir, err := os.UserHomeDir()
	if err != nil {
			return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	switch runtime.GOOS {
	case "windows":
			desktopPath = filepath.Join(homeDir, "Desktop")
	case "darwin":
			desktopPath = filepath.Join(homeDir, "Desktop")
	case "linux":
			desktopPath = filepath.Join(homeDir, "Desktop")
	default:
			return "", fmt.Errorf("unsupported platform")
	}

	return desktopPath, nil
}