package magestaff

import (
	"github.com/go-git/go-git/v5"
	"github.com/whilp/git-urls"
	"log"
	"strings"
)

func GetFullRepoName() (repoName string, err error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return
	}
	cfg, err := repo.Config()
	repoURL := cfg.Remotes["origin"].URLs[0]
	url, _ := giturls.Parse(repoURL)
	repoName = strings.Trim(url.Path, ".git")
	return
}

func GetRepoName() (repoName string, err error) {
	repo, err := GetFullRepoName()
	if err != nil {
		return
	}
	repoSplit := strings.SplitN(repo, "/", 2)
	repoName = repoSplit[1]
	return
}

func GetOrganisationName() (organisation string, err error) {
	repo, err := GetFullRepoName()
	if err != nil {
		return
	}
	repoSplit := strings.SplitN(repo, "/", 2)
	organisation = repoSplit[0]
	return
}

func GetGRepoBranch() (branch string, err error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return
	}
	ref, err := repo.Reference("HEAD", true)
	hash := ref.Hash()
	log.Println(hash)
	branch = ref.Name().Short()
	return
}

func GetGitSha(short bool) (hash string, err error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return
	}
	ref, err := repo.Reference("HEAD", true)
	hash = ref.Hash().String()
	if short {
		hash = hash[:7]
	}
	return
}
