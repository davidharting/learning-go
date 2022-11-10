package disk

import "os"

type tmp struct {
	*local
}

// Removes the temporary directory associated with the tmpDisk
func (d *tmp) Close() error {
	return os.RemoveAll(d.RootDirectory)
}

// Creates a temporary directory with a random name and a tmpDisk object to encapsulate it
func NewTmp() *tmp {
	tmpDir, _ := os.MkdirTemp("/tmp", "tmp-disk")
	d := tmp{local: &local{RootDirectory: tmpDir}}
	return &d
}
