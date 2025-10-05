package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"slices"

	"github.com/connormullett/dotman/util"
)

func CheckIfSymlinkExists(path string) bool {
	info, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}
	return info != ""
}

func Sync(args []string, quiet bool) {
	settings := util.ReadConfig()
	repoPath := settings.Path

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}
	homeDir := currentUser.HomeDir

	// abort if repo is dirty
	if util.IsRepoDirty(repoPath) {
		fmt.Println("Repository is dirty, This can cause conflicts in your history and merge conflicts which will need to be done manually")
		log.Fatalf("Please commit or stash your changes before syncing.")
	}

	// pull changes
	util.Pull(repoPath, quiet)

	// add unmanaged files (new to this client)
	dir, err := os.ReadDir(repoPath)
	if err != nil {
		log.Fatalf("Error reading repository directory: %v", err)
	}

	// check if file is already managed, add if not
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		path := filepath.Join(homeDir, entry.Name())
		if !slices.Contains(settings.ManagedFiles, path) {
			settings.ManagedFiles = append(settings.ManagedFiles, path)
		}
	}

	util.WriteConfig(settings)

	// create symlinks
	dotfiles := GetFilesList(repoPath)

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
			repoFilePath := filepath.Join(repoPath, filepath.Base(file))
			err := os.Symlink(repoFilePath, targetPath)
			if err != nil {
				log.Printf("Error creating symlink for %s: %v", file, err)
			}
		}
	}
}
