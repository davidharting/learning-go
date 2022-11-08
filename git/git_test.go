package git

import (
	"os"
	"testing"

	libgit "github.com/libgit2/git2go/v34"
	"github.com/stretchr/testify/assert"
)

func TestOpenRepository(t *testing.T) {
	_, teardown := setupTestGitRepo(t)
	defer teardown()
}

// It will be much easier to do debug the setup here if I have more capability on Filesystem.
// Specifically, I want to list the directory and I want a PutMultiple method
func setupTestGitRepo(t *testing.T) (Repo, func()) {
	tmp, err := os.MkdirTemp("/tmp", "git_test_")
	assert.NoError(t, err, "Unable to create temporary directory")

	_, err = libgit.InitRepository(tmp, true)
	assert.NoError(t, err, "Unable to initialize repo into tmp folder (%v)\n", tmp)

	repo, err := NewRepo(tmp)
	assert.NoError(t, err, "Unable to create test Repo for tmp folder (%v)\n", tmp)

	teardownTestGitRepo := func() {
		os.RemoveAll(tmp)
	}
	return repo, teardownTestGitRepo
}

// If I can't clone from a bundle, why not init and copy in files!
// Then I have full control too over what files to initialize with
