package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/davidharting/learngo/disk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type readOnlySuite struct {
	suite.Suite
	repo     *Repo
	teardown func() error
}

func (s *readOnlySuite) SetupSuite() {
	repo, teardown := setupTestGitRepo(s.T())
	s.repo = &repo
	s.teardown = teardown
}

func (s *readOnlySuite) TearDownSuite() {
	err := s.teardown()
	assert.NoError(s.T(), err)
}

func TestReadOnlySuite(t *testing.T) {
	suite.Run(t, new(readOnlySuite))
}

func (s *readOnlySuite) TestOpenRepository() {
	// If this runs at all we good
	assert.True(s.T(), true)
}

func setupTestGitRepo(t *testing.T) (Repo, func() error) {
	tmp := disk.NewTmp()

	cloneCmd := exec.Command("git", "clone", getGitBundleAbsPath(t), ".")
	cloneCmd.Dir = tmp.RootDirectory
	err := cloneCmd.Run()
	assert.NoErrorf(t, err, "Unable to clone from bundle into tmp directory %v", tmp.RootDirectory)

	repo, err := NewRepo(tmp.RootDirectory)
	assert.NoError(t, err, "Unable to create test Repo for tmp folder (%v)\n", tmp.RootDirectory)

	teardownTestGitRepo := func() error {
		return tmp.Close()
	}

	return repo, teardownTestGitRepo
}

func getGitBundleAbsPath(t *testing.T) string {
	workDir, err := os.Getwd()
	assert.NoError(t, err)
	return filepath.Join(workDir, "testdata/jaffle_shop_metrics_bundle.pack")
}
