package cd

import "github.com/go-git/go-git/v6"

func repoChanged(repoPath string, branch string, commitHash string) (bool, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return false, err
	}

	ref, err := repo.Reference(git.ReferenceName("refs/heads/"+branch), true)
	if err != nil {
		return false, err
	}

	if ref.Hash().String() != commitHash {
		return true, nil
	}

	return false, nil
}

