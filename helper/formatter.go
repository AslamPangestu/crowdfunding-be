package helper

import "path/filepath"

//https://freshman.tech/snippets/go/filename-no-extension/
func RemoveFileExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
