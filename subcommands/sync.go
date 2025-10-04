package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/connormullett/dotman/util"
)

func Sync(args []string, quiet bool) {
	settings := util.ReadConfig()
	repoPath := settings.Path

	if util.IsRepoDirty(repoPath) {
		fmt.Println("Repository is dirty, This can cause conflicts in your history and merge conflicts which will need to be done manually")
		log.Fatalf("Please commit or stash your changes before syncing.")
	}

	util.Pull(repoPath, quiet)

	// check if symlinks exist and create them if they don't
	dotfiles := GetFilesList(repoPath)

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}

	homeDir := currentUser.HomeDir
	for _, file := range dotfiles {
		targetPath := filepath.Join(homeDir, filepath.Base(file))

		// check if dotfile with the same name exists in home directory
		info, err := os.Lstat(targetPath)

		// file exists
		if err == nil {
			isSymlink := info.Mode()&os.ModeSymlink == os.ModeSymlink

			// file already exists and is not a symlink
			if !isSymlink {
				// if it exists, back it up before creating the symlink
				err := os.Rename(targetPath, targetPath+".backup")
				if err != nil {
					log.Printf("Error backing up existing file %s: %v", targetPath, err)
					continue
				}

				if !quiet {
					fmt.Printf("Backed up existing file %s to %s.backup\n", targetPath, targetPath)
				}
			} else {
				// symlink already exists, skip
				if !quiet {
					fmt.Printf("Symlink already exists for %s\n", file)
				}
				continue
			}
		}

		if !CheckIfSymlinkExists(targetPath) {
			fmt.Println("source: ", file)
			fmt.Println("target: ", targetPath)
			err := os.Symlink(file, targetPath)
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
