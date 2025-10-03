package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Push(args []string, force bool) {
	settings := ReadConfig()
	repoPath := settings.Path

	gitRepoPath := repoPath

	isDirty := CheckIfGitDirty(gitRepoPath)

	var execCmd *exec.Cmd
	var err error
	if isDirty {
		AddAndCommit(gitRepoPath, "auto-commit before push")
	}

	branch, err := GetCurrentBranch()
	if err != nil {
		log.Fatalf("Error getting current branch: %v", err)
	}

	gitArgs := []string{"push", "-u", "origin", branch}
	if force {
		gitArgs = append(gitArgs, "--force")
	}

	execCmd = exec.Command("git", gitArgs...)

	execCmd.Dir = gitRepoPath
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr

	fmt.Println("Pushing changes to remote...")

	err = execCmd.Run()
	if err != nil {
		log.Fatalf("Error pushing changes: %v", err)
	}
}
