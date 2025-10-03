package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Push(args []string) {
	settings := readConfig()
	repoPath := settings.Path

	gitRepoPath := repoPath

	isDirty := checkIfGitDirty(gitRepoPath)

	var execCmd *exec.Cmd
	var err error
	if isDirty {
		addAndCommit(gitRepoPath, "auto-commit before push")
	}

	branch, err := getCurrentBranch()
	if err != nil {
		log.Fatalf("Error getting current branch: %v", err)
	}

	execCmd = exec.Command("git", "push", "-u", "origin", branch)

	execCmd.Dir = gitRepoPath
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	fmt.Println("Pushing changes to remote...")

	err = execCmd.Run()
	if err != nil {
		log.Fatalf("Error pushing changes: %v", err)
	}
}

func addAndCommit(path, message string) {
	execCmd := exec.Command("git", "add", ".")

	execCmd.Dir = path

	err := execCmd.Run()
	if err != nil {
		log.Fatalf("Error adding changes: %v", err)
	}

	execCmd = exec.Command("git", "commit", "-m", message)
	execCmd.Dir = path

	err = execCmd.Run()
	if err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}
}

func checkIfGitDirty(path string) bool {
	cmd := "git"
	args := []string{"diff", "--quiet"}

	command := exec.Command(cmd, args...)
	command.Dir = path

	err := command.Run()
	return err != nil
}

func getCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	branch := strings.TrimSpace(string(output))
	return branch, nil
}
