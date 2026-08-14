package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var AllowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".svg":  true,
	".webp": true,
}

// IsAllowedImage checks if the file extension is an allowed image format
func IsAllowedImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return AllowedExtensions[ext]
}

// SaveUploadedFile saves an uploaded file to the specified directory with a unique timestamped name without needing gin.Context
func SaveUploadedFile(file *multipart.FileHeader, targetDir string) (string, error) {
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open source file: %w", err)
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	uniqueFilename := fmt.Sprintf("icon_%d%s", time.Now().UnixNano(), ext)
	dstPath := filepath.Join(targetDir, uniqueFilename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file content: %w", err)
	}

	return uniqueFilename, nil
}

// DeleteFile removes a file from local disk if it exists
func DeleteFile(filePath string) error {
	if filePath == "" {
		return nil
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil // File does not exist, nothing to delete
	}
	return os.Remove(filePath)
}
