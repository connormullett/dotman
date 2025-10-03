package subcommands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Add(args []string) {
	settings := ReadConfig()

	dotfilePath := args[0]

	_, err := os.Stat(dotfilePath)
	if os.IsNotExist(err) {
		fmt.Printf("Dotfile %s does not exist\n", dotfilePath)
		return
	}

	// just the file name
	fileName := filepath.Base(dotfilePath)

	// path to ~/.dotman
	repoPath := settings.Path

	// path to ~/.dotman/.git
	gitRepoPath := filepath.Join(repoPath, ".git")

	_, err = os.Stat(gitRepoPath)

	if os.IsNotExist(err) {
		fmt.Println("Dotman repository does not exist. Please run 'dotman init' first")
		return
	}

	// path to move file to (~/.dotman + fileName)
	targetPath := filepath.Join(repoPath, fileName)

	// get absolute paths to ensure symlink works correctly
	absDotfilePath, err := filepath.Abs(dotfilePath)
	if err != nil {
		log.Fatalf("Error getting absolute path for dotfile: %v", err)
	}

	absTargetPath, err := filepath.Abs(targetPath)
	if err != nil {
		log.Fatalf("Error getting absolute target path: %v", err)
	}

	fmt.Println("Target path in repository:", absTargetPath)
	fmt.Println("Dotfile path:", absDotfilePath)

	// move file to dotman repository
	err = os.Rename(absDotfilePath, absTargetPath)
	if err != nil {
		log.Fatalf("Error moving file: %v", err)
	}

	// create symlink at the original location pointing to the new location in dotman repo
	err = os.Symlink(absTargetPath, absDotfilePath)
	if err != nil {
		log.Fatalf("Error creating symlink: %v", err)
	}

	// git add
	GitAdd(repoPath, targetPath)

	// create commit
	CommitNewFile(repoPath, fileName)
}
