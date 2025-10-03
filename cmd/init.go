package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Init(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: dotman init <remote>")
		return
	}

	remote := args[0]
	fmt.Println("Initializing dotman repository with remote:", remote)

	settings := readConfig()
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

	executeInitRepository(repoPath)

	executeAddRemote(repoPath, remote)
}

func createCommand(commandName string, cmdArgs []string, path string) *exec.Cmd {
	cmd := exec.Command(commandName, cmdArgs...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func executeAddRemote(path, remote string) {
	commandName := "git"
	cmdArgs := []string{"remote", "add", "origin", remote}

	cmd := createCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}

func executeInitRepository(path string) {
	commandName := "git"
	cmdArgs := []string{"init"}

	cmd := createCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}
