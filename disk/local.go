package disk

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
)

type local struct {
	RootDirectory string
}

func NewLocal(rootDirectory string) *local {
	return &local{RootDirectory: rootDirectory}
}

func (f *local) get(path string) (string, *DiskGetError) {
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

func getFullPath(f *local, relativePath string) string {
	return filepath.Join(f.RootDirectory, relativePath)
}

func (l *local) put(path string, contents string) error {
	fullPath := getFullPath(l, path)

	// TODO: What modes to use?
	os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
	err := os.WriteFile(fullPath, bytes.NewBufferString(contents).Bytes(), fs.ModePerm)
	return err
}

func (l *local) putMany(files []file) []filePutError {
	errors := make([]filePutError, 0)

	for _, file := range files {
		err := l.put(file.relativePath, file.contents)
		if err != nil {
			errors = append(errors, filePutError{relativePath: file.relativePath, originalErorr: err})
		}
	}

	return errors
}

func (f *local) ListAll() []FileInfo {
	files := make([]FileInfo, 0)

	filepath.WalkDir(f.RootDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		info, err := d.Info()

		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		relativePath, _ := filepath.Rel(f.RootDirectory, path)

		files = append(files, FileInfo{RelativePath: relativePath, SizeInBytes: info.Size()})
		return nil
	})

	return files
}
