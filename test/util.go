package test

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MakeFile - make a file containing the string contents and return the files
//   name, or an error if the file can not be made
func MakeFile(baseName string, fileType string, contents string) string {
	var f *os.File
	var filePath string
	var err error
	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err == nil {
		fileName := fmt.Sprintf("%s%X.%s", baseName, b, fileType)
		filePath = filepath.Join(os.TempDir(), fileName)
		f, err = os.Create(filePath)
		if err == nil {
			_, err = io.WriteString(f, contents)
		}
	}
	return filePath
}
