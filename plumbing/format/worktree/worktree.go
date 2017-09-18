package worktree

import (
	"errors"

	"gopkg.in/src-d/go-git.v4/plumbing"
)

var (
	ErrInfoNil       = errors.New("Info is nil")
	ErrHeadNil       = errors.New("Info head is nil")
	ErrWorktreeEmpty = errors.New("Info worktree empty")
)

// Info information of a  worktree
type Info struct {
	Name string
	Wt   string
	HEAD *plumbing.Reference
}

// Validate check if info is valid
func (i *Info) Validate() error {
	if i == nil {
		return ErrInfoNil
	}

	if i.HEAD == nil {
		return ErrHeadNil
	}

	if i.Wt == "" {
		return ErrWorktreeEmpty
	}

	return nil
}
