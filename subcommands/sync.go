package subcommands

import (
	"fmt"
	"log"
)

func Sync(args []string, quiet bool) {
	settings := ReadConfig()
	repoPath := settings.Path

	if CheckIfGitDirty(repoPath) {
		fmt.Println("Repository is dirty, This can cause conflicts in your history and merge conflicts which will need to be done manually")
		log.Fatalf("Please commit or stash your changes before syncing.")
	}

	GitPull(repoPath, quiet)
}
