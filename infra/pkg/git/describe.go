package git

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func Describe() (string, error) {
	// Get repository
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	repo, err := git.PlainOpen(path.Join(dir, "../"))
	if err != nil {
		return "", err
	}

	// Get all tags
	tags := make(map[plumbing.Hash]*plumbing.Reference)
	tagsIter, err := repo.Tags()
	if err != nil {
		return "", err
	}

	err = tagsIter.ForEach(func(t *plumbing.Reference) error {
		tags[t.Hash()] = t
		return nil
	})
	if err != nil {
		return "", err
	}

	// Find commit history from HEAD
	head, err := repo.Head()
	if err != nil {
		return "", err
	}

	commitsIter, err := repo.Log(&git.LogOptions{
		From:  head.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	if err != nil {
		return "", err
	}

	// Search the tag
	var tag *plumbing.Reference
	var count int
	err = commitsIter.ForEach(func(c *object.Commit) error {
		if t, ok := tags[c.Hash]; ok {
			tag = t
		}
		if tag != nil {
			return storer.ErrStop
		}
		count++
		return nil
	})
	if count == 0 {
		return fmt.Sprint(tag.Name().Short()), nil
	} else {
		return fmt.Sprintf("%v-%v-g%v",
			tag.Name().Short(),
			count,
			head.Hash().String()[0:7],
		), nil
	}
}
