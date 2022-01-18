package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	rootPath := os.Getenv("ROOT_PATH")
	if rootPath == "" {
		log.Fatalln("no root path set")
	}

	fmt.Println("Search for a git repo:")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("could not parse repo", err)
	}
	in = strings.Trim(in, " \n")

	r := repos{}
	r, _ = r.FindGitRepos(rootPath)
	r, _ = r.GetRepoNames()

	filteredNames, _ := r.FilterNames(in)

	for {
		filteredNames.PPrint()
		if len(filteredNames) == 0 {
			fmt.Println("Oops! Couldn't find any repos with that name")
			return
		}

		if len(filteredNames) == 1 {
			fmt.Printf("found repo: %s. Launching now...\n", filteredNames[0].name)
			cmd := exec.Command("code", filteredNames[0].path)
			cmd.Run()
			return
		}

		fmt.Println("\nRefine your search.")
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("could not parse repo", err)
		}
		in = strings.Trim(in, " \n")

		filteredNames, _ = filteredNames.FilterNames(in)
	}
}
