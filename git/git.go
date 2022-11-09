package git

// Note this package requires libgit2 to be installed on the system. I wish it was packaged in!
import (
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

func NewRepo(path string) (Repo, error) {
	libgitRepo, err := libgit.OpenRepository(path)
	if err != nil {
		return nil, err
	}

	return &repo{libgitRepo: libgitRepo, path: path}, nil
}

type Repo interface {
	CurrentBranch() (string, error)
	LatestCommit() (string, error)
	GetPath() string
}

func (r *repo) CurrentBranch() (string, error) {
	head, err := r.libgitRepo.Head()
	if err != nil {
		return "", err
	}
	return head.Name(), nil
}

func (r *repo) LatestCommit() (string, error) {
	head, _ := r.libgitRepo.Head()
	ref, _ := head.Resolve()
	annotatedCommit, _ := r.libgitRepo.AnnotatedCommitFromRef(ref)
	oid := annotatedCommit.Id()
	commit, _ := r.libgitRepo.LookupCommit(oid)
	return commit.Message(), nil
}

func (r *repo) GetPath() string {
	return r.path
}
