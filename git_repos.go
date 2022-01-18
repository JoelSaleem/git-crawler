package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type GitRepo struct {
	path      string
	name      string
	nameLower string
}

type repos []GitRepo

func (r repos) FindGitRepos(rootPath string) (repos, error) {
	fileSys := os.DirFS(rootPath)

	fs.WalkDir(fileSys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if len(path) > 4 {
			s := path[len(path)-4:]
			if s == ".git" {
				pwd := rootPath + "/" + path
				pwd = pwd[:len(pwd)-5]

				r = append(r, GitRepo{path: pwd})
			}
		}
		return nil
	})

	return r, nil
}

func (r repos) GetRepoNames() (repos, error) {
	if len(r) == 0 {
		return make(repos, 0), fmt.Errorf("no repos found")
	}

	out := make(repos, len(r))

	for i, repo := range r {
		splitPath := strings.Split(repo.path, "/")
		name := splitPath[len(splitPath)-1]
		newRepo := GitRepo{path: repo.path, name: name, nameLower: strings.ToLower(name)}
		out[i] = newRepo
	}

	return out, nil
}

func (r repos) FilterNames(input string) (repos, error) {
	filteredNames := repos{}
	for _, repo := range r {
		if strings.Contains(repo.nameLower, strings.ToLower(input)) {
			filteredNames = append(filteredNames, repo)
		}
	}

	return filteredNames, nil
}

func (r repos) PPrint() {
	for _, v := range r {
		fmt.Println(v.name)
	}
}
