package disk

import "fmt"

type Disk interface {
	DiskReader
	DiskWriter
}

type DiskReader interface {
	get(string) (string, DiskGetError)
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

func (e *DiskGetError) Error() string {
	return fmt.Sprintf("Unable to read file %v, due to error code %v", e.path, e.code)
}

func (e *DiskGetError) String() string {
	return e.Error()
}

type DiskWriter interface {
	put(string, string) error
}

type DiskLister interface {
	listAll() []file
	// list - but take a depth and a starting subdirectory
}

type file struct {
	relativePath string
	sizeInBytes  int64
}
