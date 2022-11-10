package git

// Note this package requires libgit2 to be installed on the system. I wish it was packaged in!
import (
	"fmt"

	libgit "github.com/libgit2/git2go/v34"
)

func MostRecentCommit() (string, error) {
	repo, _ := libgit.OpenRepository(".")
	head, _ := repo.Head()
	ref, _ := head.Resolve()
	annotatedCommit, _ := repo.AnnotatedCommitFromRef(ref)
	oid := annotatedCommit.Id()
	commit, _ := repo.LookupCommit(oid)
	return commit.Message(), nil
}

type repo struct {
	libgitRepo *libgit.Repository
	path       string
}

func NewRepo(path string) (*repo, error) {
	libgitRepo, err := libgit.OpenRepository(path)
	if err != nil {
		return nil, err
	}

	return &repo{libgitRepo: libgitRepo, path: path}, nil
}

type Repo interface {
	CurrentBranch() (string, error)
	LatestCommit() (*commit, error)
	GetPath() string
}

func (r *repo) CurrentBranch() (string, error) {
	head, err := r.libgitRepo.Head()
	if err != nil {
		return "", err
	}
	branch, err := head.Branch().Name()
	if err != nil {
		return "", err
	}

	return branch, nil
}

type commit struct {
	Author  string
	Message string
}

// Get the latest commit on the current branch
func (r *repo) LatestCommit() (*commit, error) {
	head, _ := r.libgitRepo.Head()
	ref, _ := head.Resolve()
	annotatedCommit, _ := r.libgitRepo.AnnotatedCommitFromRef(ref)
	oid := annotatedCommit.Id()
	c, err := r.libgitRepo.LookupCommit(oid)

	if err != nil {
		return nil, err
	}

	com := &commit{Author: fmt.Sprintf("%v <%v>", c.Author().Name, c.Author().Email), Message: c.Message()}

	return com, nil
}

func (r *repo) GetPath() string {
	return r.path
}
