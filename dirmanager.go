package testutils

import (
	"code.google.com/p/go-uuid/uuid"
	"os"
	"path/filepath"
)

type DirectoryManager struct {
	Path string
}

func (dm *DirectoryManager) Close() {
	os.RemoveAll(dm.Path)
}

func NewDirectoryManager() *DirectoryManager {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	cacheDirectory := filepath.Join(pwd, uuid.New())
	err = os.MkdirAll(cacheDirectory, 0777)
	if err != nil {
		panic(err.Error())
	}
	return &DirectoryManager{cacheDirectory}
}
