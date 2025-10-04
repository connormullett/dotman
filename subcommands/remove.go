package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/connormullett/dotman/util"
)

func Remove(args []string) {
	settings := util.ReadConfig()

	dotfileName := args[0]

	path := settings.Path

	// Path to the file in the dotman repository
	dotfilePath := filepath.Join(path, dotfileName)

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("Error getting current user: %v", err)
	}

	homeDir := currentUser.HomeDir

	// Path to the symlink in home directory
	symlinkPath := filepath.Join(homeDir, dotfileName)

	// Check if file exists in dotman repository
	if _, err := os.Stat(dotfilePath); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist in dotman repository at %s", dotfileName, dotfilePath)
	}

	// Check if symlink exists in home directory
	fi, err := os.Lstat(symlinkPath)
	if err == nil {
		// File/symlink exists, verify it's a symlink
		if fi.Mode()&os.ModeSymlink == 0 {
			log.Fatalf("File at %s exists but is not a symlink. Aborting to avoid data loss.", symlinkPath)
		}

		// Verify the symlink points to the dotman repository
		target, err := os.Readlink(symlinkPath)
		if err != nil {
			log.Fatalf("Error reading symlink: %v", err)
		}

		// Resolve to absolute path if relative
		if !filepath.IsAbs(target) {
			target = filepath.Join(filepath.Dir(symlinkPath), target)
		}

		absTarget, _ := filepath.Abs(target)
		absDotfilePath, _ := filepath.Abs(dotfilePath)

		if absTarget != absDotfilePath {
			log.Fatalf("Symlink at %s points to %s, not to dotman repository at %s. Aborting.", symlinkPath, target, dotfilePath)
		}

		// Remove the symlink
		fmt.Printf("Removing symlink at %s\n", symlinkPath)
		err = os.Remove(symlinkPath)
		if err != nil {
			log.Fatalf("Error removing symlink: %v", err)
		}
	} else if !os.IsNotExist(err) {
		log.Fatalf("Error checking symlink: %v", err)
	}

	// Move the file from dotman repository back to home directory
	fmt.Printf("Moving %s back to %s\n", dotfilePath, symlinkPath)
	err = os.Rename(dotfilePath, symlinkPath)
	if err != nil {
		log.Fatalf("Error moving file back to original location: %v", err)
	}

	util.AddAndCommit(path, "Removed "+dotfileName)

	if settings.ManagedFiles != nil {
		for i, file := range settings.ManagedFiles {
			if filepath.Base(file) == dotfileName {
				// Remove from ManagedFiles slice
				settings.ManagedFiles = append(settings.ManagedFiles[:i], settings.ManagedFiles[i+1:]...)
				break
			}
		}
	}
	util.WriteConfig(settings)

	fmt.Printf("Successfully removed %s from dotman management\n", dotfileName)
}
