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
