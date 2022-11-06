package disk

import "os"

type tmpDisk struct {
	*FilesystemDisk
}

func (d *tmpDisk) Close() error {
	return os.RemoveAll(d.rootDirectory)
}

func NewTmpDisk() *tmpDisk {
	tmp, _ := os.MkdirTemp("/tmp", "tmp-disk")
	d := tmpDisk{FilesystemDisk: &FilesystemDisk{rootDirectory: tmp}}
	return &d
}
