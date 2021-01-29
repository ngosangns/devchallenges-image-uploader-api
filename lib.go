package main

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func CreateFile(filePath string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(filePath), 0770); err != nil {
		return nil, err
	}
	return os.Create(filePath)
}

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v", fn, line, err)
}

func ErrorUploadingResponse() string {
	return "Error while uploading..."
}
