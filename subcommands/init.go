package subcommands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/connormullett/dotman/util"
)

func Init(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: dotman init <remote>")
		return
	}

	remote := args[0]
	fmt.Println("Initializing dotman repository with remote:", remote)

	settings := util.ReadConfig()
	repoPath := settings.Path

	err := os.MkdirAll(repoPath, 0755)
	if err != nil {
		panic(err)
	}

	gitRepo := filepath.Join(repoPath, ".git")
	fmt.Println("Checking for existing git repository at", gitRepo)

	info, err := os.Stat(gitRepo)

	exists := true
	if err != nil {
		exists = false
	}
	if info != nil && !info.IsDir() {
		exists = false
	}

	if exists {
		println("dotman repository already initialized at", repoPath)
		return
	}

	util.InitRepository(repoPath)

	util.AddRemote(repoPath, remote)
}
