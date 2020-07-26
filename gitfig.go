package gitfig

import (
  "encoding/json"
  "gopkg.in/yaml.v2"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/storage/memory"
)

func LoadYaml(filename string, v interface{}, opts *FigOptions) error {
  contents, err := LoadFile(filename, opts)
  if err != nil {
    return err
  }
  bs := []byte(contents)
  err = yaml.Unmarshal(bs, v)
  if err != nil {
    return err
  }
  return nil
}

func LoadJson(filename string, v interface{}, opts *FigOptions) error {
  contents, err := LoadFile(filename, opts)
  if err != nil {
    return err
  }

  bs := []byte(contents)
  err = json.Unmarshal(bs, v)
  if err != nil {
    return err
  }
  return nil
}

func LoadFile(filename string, opts *FigOptions) (string, error) {

  r, err := git.Clone(memory.NewStorage(), nil, opts.GetCloneOptions())
  if err != nil {
    return "", err
  }

  ref, err := r.Head()
  if err != nil {
    return "", err
  }

  commit, err := r.CommitObject(ref.Hash())
  if err != nil {
    return "", err
  }

  f, err := commit.File(filename)
  if err != nil {
    return "", err
  }
  return f.Contents()

  // tree, err := commit.Tree()
  // if err != nil {
  //   return "", err
  // }

  // f, err := tree.File(filename)
  // if err != nil {
  //   return "", err
  // }
}
