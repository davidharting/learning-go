package disk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestTmpDiskGetMissingFile(t *testing.T) {
	disk := NewTmp()
	defer disk.Close()

	contents, err := disk.get("path/to/missing/file")
	assert.Empty(t, contents)
	assert.Error(t, err)
	assert.Equal(t, err.code, NotFound)
}

func TestClosingDeletesTmpDir(t *testing.T) {
	disk := NewTmp()
	_, err := os.ReadDir(disk.RootDirectory)
	assert.NoError(t, err)

	disk.Close()
	_, err = os.ReadDir(disk.RootDirectory)
	assert.Error(t, err)
}

func TestPutCreatesFile(t *testing.T) {
	disk := NewTmp()
	defer disk.Close()

	_, err := disk.get("path/abc.txt")
	assert.Error(t, err, "File should not exist yet.")

	putErr := disk.put("path/abc.txt", "hello")
	assert.NoError(t, putErr, "Unable to put file")

	got, secondGetErr := disk.get("path/abc.txt")
	assert.Nil(t, secondGetErr, "Got an error retrieving file contents after putting them")
	assert.Equal(t, "hello", got)
}

func TestPutUpdatesFile(t *testing.T) {
	disk := NewTmp()
	defer disk.Close()

	putErr := disk.put("path/to/file.txt", "Hello, World")
	assert.NoError(t, putErr, "Problem doing put to create")
	putErr = disk.put("path/to/file.txt", "Goodbye, Friends")
	assert.NoError(t, putErr, "Problem doing put to update")

	contents, getErr := disk.get("path/to/file.txt")
	assert.Nil(t, getErr)
	assert.Equal(t, contents, "Goodbye, Friends")
}

// func TestLeadingSlashesDontMatter
// func TestMultipleLevels of nesting
// func TestPut into path where some but not all of the ancestors exist

type ListAllSuite struct {
	suite.Suite
}

func TestListAllTestSuite(t *testing.T) {
	suite.Run(t, new(ListAllSuite))
}

func (s *ListAllSuite) TestEmptyDirectoryReturnsEmptyList() {
	disk := NewTmp()
	defer disk.Close()

	assert.Empty(s.T(), disk.ListAll())
}

func (s *ListAllSuite) TestListsFilesInDirectoryRecursively() {
	disk := NewTmp()
	defer disk.Close()

	disk.put("abc.txt", "hello")
	disk.put("/one/two/def.txt", "1234")
	disk.put("/one/two/ghi.txt", "cool")

	files := disk.ListAll()

	assert.Len(s.T(), files, 3, "Got the wrong number of files back.")

	expectedPaths := []string{"abc.txt", "one/two/def.txt", "one/two/ghi.txt"}
	gotPaths := make([]string, 0)

	for _, f := range files {
		gotPaths = append(gotPaths, f.relativePath)
	}

	assert.Equal(s.T(), expectedPaths, gotPaths)
}

func TestPutManyWorksWithEmptySlice(t *testing.T) {
	disk := NewTmp()
	defer disk.Close()

	errors := disk.putMany(make([]file, 0))
	assert.Empty(t, errors)
}

func TestPutMany(t *testing.T) {
	disk := NewTmp()
	defer disk.Close()

	files := [3]file{
		{relativePath: "models/staging/stg_customers.sql", contents: "select 1"},
		{relativePath: "readme.md", contents: "# Project"},
		{relativePath: "models/customers.sql", contents: "select 2"},
	}

	errors := disk.putMany(files[0:3])
	assert.Empty(t, errors)

	list := disk.ListAll()
	assert.Len(t, list, 3)

	paths := make([]string, 0)
	for _, fileInfo := range list {
		paths = append(paths, fileInfo.relativePath)
	}
	assert.Contains(t, paths, "models/staging/stg_customers.sql")
	assert.Contains(t, paths, "readme.md")
	assert.Contains(t, paths, "models/customers.sql")
}
