package disk

import "os"

type tmpDisk struct {
	*Local
}

// Removes the temporary directory associated with the tmpDisk
func (d *tmpDisk) Close() error {
	return os.RemoveAll(d.rootDirectory)
}

// Creates a temporary directory with a random name and a tmpDisk object to encapsulate it
func NewTmpDisk() *tmpDisk {
	tmp, _ := os.MkdirTemp("/tmp", "tmp-disk")
	d := tmpDisk{Local: &Local{rootDirectory: tmp}}
	return &d
}
