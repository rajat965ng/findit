package service

import (
	"findings/model"
	"fmt"
	"os"
	"regexp"

	"github.com/go-git/go-git/v5"
)

type IGitService interface {
	GitClone() (string, error)
	GrepText(patterns []string) ([]string, error)
	CleanUp() error
}

type gitService struct {
	repository *model.Repository
	target     string
}

func NewGitService(repository *model.Repository, target string) *gitService {
	return &gitService{repository: repository, target: target}
}

func (svc *gitService) GitClone() (string, error) {
	r, err := git.PlainClone(svc.target+"/"+svc.repository.Name, false, &git.CloneOptions{
		URL:               svc.repository.Url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	if err != nil {
		return "", err
	}

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		return "", err
	}
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return "", err
	}
	fmt.Println(commit)
	return svc.target + "/" + svc.repository.Name, nil
}

func (svc *gitService) GrepText(patterns []string) ([]string, error) {
	regexps := make([]*regexp.Regexp, 0)
	for _, j := range patterns {
		regexps = append(regexps, regexp.MustCompile(j))
	}
	r, _ := git.PlainOpen(svc.target + "/" + svc.repository.Name)
	worktree, _ := r.Worktree()
	grepResults, err := worktree.Grep(&git.GrepOptions{
		Patterns:    regexps,
		InvertMatch: false,
	})
	if err != nil {
		return nil, err
	}
	matches := make([]string, 0)
	for _, j := range grepResults {
		matches = append(matches, fmt.Sprintf("%s:%d:%s", j.FileName, j.LineNumber, j.Content))
	}
	return matches, nil
}

func (svc *gitService) CleanUp() error {
	return os.RemoveAll(svc.target + "/" + svc.repository.Name)
}
