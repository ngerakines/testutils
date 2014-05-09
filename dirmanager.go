package testutils

import (
	"code.google.com/p/go-uuid/uuid"
	"os"
	"path/filepath"
)

type directoryManager struct {
	Path string
}

func (dm *directoryManager) Close() {
	os.RemoveAll(dm.Path)
}

func NewDirectoryManager() *directoryManager {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	cacheDirectory := filepath.Join(pwd, uuid.New())
	os.MkdirAll(cacheDirectory, 00777)
	return &directoryManager{cacheDirectory}
}
