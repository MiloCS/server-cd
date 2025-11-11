package cd

import (
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/storage/memory"
)

type repo struct {
	CommitHash    string
	TopLevelFiles []string
}

func lightweightCheckout(repoURL string, branch string) (repo, error) {
	checkedOut, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.ReferenceName("refs/heads/" + branch),
		SingleBranch:  true,
		Depth:         1,
	})
	if err != nil {
		return repo{}, err
	}

	ref, err := checkedOut.Head()
	if err != nil {
		return repo{}, err
	}

	files, err := checkedOut.Worktree()
	if err != nil {
		return repo{}, err
	}

	var topLevelFiles []string
	entries, err := files.Filesystem.ReadDir(".")
	if err != nil {
		return repo{}, err
	}
	for _, entry := range entries {
		topLevelFiles = append(topLevelFiles, entry.Name())
	}

	return repo{
		CommitHash:    ref.Hash().String(),
		TopLevelFiles: topLevelFiles,
	}, nil
}

func repoChanged(repoURL string, branch string, commitHash string) (bool, error) {
	repo, err := lightweightCheckout(repoURL, branch)
	if err != nil {
		return false, err
	}

	return repo.CommitHash != commitHash, nil
}
