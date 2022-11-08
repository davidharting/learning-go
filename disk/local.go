package disk

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
)

type Local struct {
	rootDirectory string
}

func (f *Local) get(path string) (string, *DiskGetError) {
	fullPath := getFullPath(f, path)
	content, err := os.ReadFile(fullPath)

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

func getFullPath(f *Local, relativePath string) string {
	return filepath.Join(f.rootDirectory, relativePath)
}

func (f *Local) put(path string, contents string) error {
	fullPath := getFullPath(f, path)

	// TODO: What modes to use?
	os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	err := os.WriteFile(fullPath, bytes.NewBufferString(contents).Bytes(), fs.ModePerm)
	return err
}