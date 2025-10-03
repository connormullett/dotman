package subcommands

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GitPull(path string) {
	execCmd := CreateCommand("git", []string{"pull"}, path)

	err := execCmd.Run()
	if err != nil {
		log.Fatalf("Error pulling changes: %v", err)
	}
}

func AddAndCommit(path, message string) {
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

func CheckIfGitDirty(path string) bool {
	cmd := "git"
	args := []string{"diff", "--quiet"}

	command := exec.Command(cmd, args...)
	command.Dir = path

	err := command.Run()
	return err != nil
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	branch := strings.TrimSpace(string(output))
	return branch, nil
}

func GitAdd(path, file string) {
	commandName := "git"
	cmdArgs := []string{"add", file}

	cmd := CreateCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}

func CommitNewFile(path, fileName string) {
	commandName := "git"
	cmdArgs := []string{"commit", "-m", fmt.Sprintf("Add %s", fileName)}

	cmd := CreateCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}

func CreateCommand(commandName string, cmdArgs []string, path string) *exec.Cmd {
	cmd := exec.Command(commandName, cmdArgs...)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func AddRemote(path, remote string) {
	commandName := "git"
	cmdArgs := []string{"remote", "add", "origin", remote}

	cmd := CreateCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}

func InitRepository(path string) {
	commandName := "git"
	cmdArgs := []string{"init"}

	cmd := CreateCommand(commandName, cmdArgs, path)

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed in %s: %v\n", path, err)
	}
}
