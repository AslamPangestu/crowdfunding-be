package helper

import (
	"os"
	"path/filepath"
)

//https://freshman.tech/snippets/go/filename-no-extension/
func RemoveFileExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func GenerateURL(location string, filename string) string {
	baseURL := os.Getenv("BACKEND_URL")
	baseDir := os.Getenv("STORAGE_URL")
	return baseURL + "/" + baseDir + "/" + location + "/" + filename
}

func GeneratePath(location string, filename string) string {
	baseDir := os.Getenv("STORAGE_PATH")
	return baseDir + "/" + location + "/" + filename
}
