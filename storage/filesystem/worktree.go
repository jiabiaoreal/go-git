package filesystem

import (
	"gopkg.in/src-d/go-git.v4/plumbing/format/worktree"
	"gopkg.in/src-d/go-git.v4/storage/filesystem/internal/dotgit"
)

// WorktreeStorage managers worktree info
type WorktreeStorage struct {
	dir *dotgit.DotGit
}

// ListWorktrees list all created worktree except the default
func (wts WorktreeStorage) ListWorktrees() ([]worktree.Info, error) {

	return nil, nil
}

// SwitchToWorktree switch to the given worktree
// if error happens, origin repo not changed
func (wts WorktreeStorage) SwitchToWorktree(name string) error {

	return nil
}

// SetWorktree init a  new worktree
func (wts WorktreeStorage) SetWorktree(wt worktree.Info) error {

	return nil
}

// Remove  removes a worktree
func (wts WorktreeStorage) Remove(name string) error {

	return nil
}
