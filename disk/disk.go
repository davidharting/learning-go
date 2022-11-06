package disk

import "fmt"

// TODO: Create DiskReader, DiskWriter, and DiskWalker interfaces

type Disk interface {
	get(string) (string, *DiskGetError)
	put(string, string) error
}

type DiskGetErrorCode string

const (
	NotFound     DiskGetErrorCode = "NotFound"
	NotPermitted DiskGetErrorCode = "NotPermitted"
	Unknown      DiskGetErrorCode = "Unknown"
)

type DiskGetError struct {
	code DiskGetErrorCode
	path string
}

func (e DiskGetError) Error() string {
	return fmt.Sprintf("Unable to read file %v, due to error code %v", e.path, e.code)
}
