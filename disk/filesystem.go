package disk

import (
	"bytes"
	"os"
	"path/filepath"
)

type FilesystemDisk struct {
	rootDirectory string
}

func (f *FilesystemDisk) get(path string) (string, *DiskGetError) {
	fullPath := getFullPath(f, path)
	content, err := os.ReadFile(fullPath) // Todo: Concat path with root directory

	if os.IsPermission(err) {
		return "", &DiskGetError{path: path, code: NotPermitted}
	}
	if os.IsNotExist(err) {
		return "", &DiskGetError{path: path, code: NotFound}
	}
	if err != nil {
		return "", &DiskGetError{path: path, code: Unknown}
	}

	return bytes.NewBuffer(content).String(), nil
}

// func (f *FilesystemDisk) put(path string) error {
// 	return nil
// }

func getFullPath(f *FilesystemDisk, relativePath string) string {
	return filepath.Join(f.rootDirectory, relativePath)
}
