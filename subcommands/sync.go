package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func Sync(args []string, quiet bool) {
	settings := ReadConfig()
	repoPath := settings.Path

	if CheckIfGitDirty(repoPath) {
		fmt.Println("Repository is dirty, This can cause conflicts in your history and merge conflicts which will need to be done manually")
		log.Fatalf("Please commit or stash your changes before syncing.")
	}

	GitPull(repoPath, quiet)

	// check if symlinks exist and create them if they don't
	dotfiles := GetFilesList(repoPath)

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}

	homeDir := currentUser.HomeDir
	for _, file := range dotfiles {
		sourcePath := filepath.Join(repoPath, file)
		targetPath := filepath.Join(homeDir, file)

		if !CheckIfSymlinkExists(targetPath) {
			err := os.Symlink(sourcePath, targetPath)
			if err != nil {
				log.Printf("Error creating symlink for %s: %v", file, err)
			}
		}
	}
}

func CheckIfSymlinkExists(path string) bool {
	info, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}
	return info != ""
}
