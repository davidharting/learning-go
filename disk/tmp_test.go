package disk

import (
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
