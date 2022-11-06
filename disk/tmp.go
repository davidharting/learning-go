package disk

import "os"

type tmpDisk struct {
	*FilesystemDisk
}

func (*tmpDisk) Close() error {
	// TODO: delete rootdirectory
	return nil
}

func NewTmpDisk() *tmpDisk {
	tmp, _ := os.MkdirTemp("/tmp", "tmp-disk")
	d := tmpDisk{FilesystemDisk: &FilesystemDisk{rootDirectory: tmp}}
	return &d
}
