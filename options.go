package gitfig

import (
  "os"

  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/plumbing"
  "github.com/go-git/go-git/v5/plumbing/transport"
)

type FigOptions struct {
  // add tag as option?
  Repo, Branch string
  Auth transport.AuthMethod
  Debug bool
}

func (opts *FigOptions) GetCloneOptions() *git.CloneOptions {
  cloneOpt := &git.CloneOptions{
    URL: opts.Repo,
    SingleBranch: true,
    Depth: 1,
  }
  if opts.Branch == "" {
    cloneOpt.ReferenceName = plumbing.NewBranchReferenceName("master")
  } else {
    cloneOpt.ReferenceName = plumbing.NewBranchReferenceName(opts.Branch)
  }

  if opts.Auth != nil {
    cloneOpt.Auth = opts.Auth
  }

  if opts.Debug {
    cloneOpt.Progress = os.Stdout
  }

  return cloneOpt
}
