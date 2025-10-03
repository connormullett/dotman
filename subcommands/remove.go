package subcommands

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func Remove(args []string) {
	settings := ReadConfig()

	dotfileName := args[0]

	path := settings.Path

	dotfilePath := filepath.Join(path, dotfileName)

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}

	homeDir := currentUser.HomeDir

	symlinkPath := filepath.Join(homeDir, dotfileName)

	_, err = os.Lstat(symlinkPath)
	if os.IsExist(err) {
		err = os.Remove(symlinkPath)
		if err != nil {
			log.Fatalf("Error removing symlink: %v", err)
		}
	}

	err = os.Rename(dotfilePath, symlinkPath)
	if err != nil {
		log.Fatalf("Error moving file back to original location: %v", err)
	}

	AddAndCommit(path, "Removed "+dotfileName)
}
