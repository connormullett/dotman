package subcommands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/connormullett/dotman/util"
)

func Doctor(args []string, fix bool) {
	settings := util.ReadConfig()

	path := settings.Path

	managedFiles := GetFilesList(path)

	currentUser, err := user.Current()

	if err != nil {
		log.Fatal("Failed to get current user:", err)
	}

	homeDir := currentUser.HomeDir

	var brokenLinks []string
	for _, file := range managedFiles {
		dotfilePath := filepath.Join(homeDir, file)
		fmt.Println("Checking:", dotfilePath)
		if isBrokenSymlink(dotfilePath) {
			brokenLinks = append(brokenLinks, dotfilePath)
		}
	}

	if len(brokenLinks) == 0 {
		fmt.Println("All symlinks are healthy!")
	} else {
		fmt.Println("Broken symlinks found:")
		for _, link := range brokenLinks {
			fmt.Println(" -", link)
		}
	}

	if !fix {
		return
	}

	for _, link := range brokenLinks {
		err := os.Remove(link)
		if err != nil {
			log.Printf("Failed to remove broken symlink %s: %v", link, err)
			continue
		}

		actualFile := filepath.Join(path, filepath.Base(link))

		err = os.Symlink(actualFile, link)
		if err != nil {
			log.Printf("Failed to recreate symlink %s -> %s: %v", link, actualFile, err)
		}
	}
}

func isBrokenSymlink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		fmt.Println("failed to lstat file", path, err)
		return false
	}

	if fi.Mode()&os.ModeSymlink == 0 {
		// Not a symlink
		return false
	}

	// It's a symbolic link, check if target exists
	target, err := os.Readlink(path)
	if err != nil {
		fmt.Println("failed to readlink target of", path, err)
		return false
	}
	// If target is not absolute, join with directory of symlink
	if !filepath.IsAbs(target) {
		target = filepath.Join(filepath.Dir(path), target)
	}
	_, err = os.Stat(target)
	if os.IsNotExist(err) {
		return true // Target doesn't exist, broken
	}
	if err != nil {
		// Check for "too many levels of symbolic links" error (circular symlink)
		if errors.Is(err, syscall.ELOOP) {
			fmt.Println("circular symlink detected at", path)
			return true // Circular symlinks are broken
		}
		fmt.Println("failed to stat resolved target of", path, err)
		return false
	}
	return false // Target exists, so it's not broken
}
