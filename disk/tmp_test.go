package disk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpDiskGetMissingFile(t *testing.T) {
	disk := NewTmpDisk()
	defer disk.Close()

	contents, err := disk.get("path/to/missing/file")
	assert.Empty(t, contents)
	assert.Error(t, err)
	assert.Equal(t, err.code, NotFound)
}

func TestClosingDeletesTmpDir(t *testing.T) {
	disk := NewTmpDisk()
	_, err := os.ReadDir(disk.rootDirectory)
	assert.NoError(t, err)

	disk.Close()
	_, err = os.ReadDir(disk.rootDirectory)
	assert.Error(t, err)
}

func TestPutCreatesFile(t *testing.T) {
	disk := NewTmpDisk()
	defer disk.Close()

	_, err := disk.get("path/abc.txt")
	assert.Error(t, err, "File should not exist yet.")

	putErr := disk.put("path/abc.txt", "hello")
	assert.NoError(t, putErr, "Unable to put file")

	got, secondGetErr := disk.get("path/abc.txt")
	assert.Nil(t, secondGetErr, "Got an error retrieving file contents after putting them")
	assert.Equal(t, "hello", got)
}

func TestPutUpdatesFile(t *testing.T) {}
