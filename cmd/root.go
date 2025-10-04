package cmd

import (
	"github.com/connormullett/dotman/subcommands"

	"github.com/spf13/cobra"
)

func newInitSubCommand() *cobra.Command {
	var initSubCommand = &cobra.Command{
		Use:   "init",
		Short: "Initialize dotman repository",
		Long:  `Initialize dotman repository`,
		Run: func(cmd *cobra.Command, args []string) {
			subcommands.Init(args)
		},
	}
	return initSubCommand
}

func newAddSubCommand() *cobra.Command {
	var addSubCommand = &cobra.Command{
		Use:   "add",
		Short: "Add a new dot file to dotman repository",
		Long:  `Add a new dot file to dotman repository`,
		Run: func(cmd *cobra.Command, args []string) {
			subcommands.Add(args)
		},
	}
	return addSubCommand
}

func newRemoveSubCommand() *cobra.Command {
	var removeSubCommand = &cobra.Command{
		Use:   "remove",
		Short: "Remove a dot file from dotman repository",
		Long:  `Remove a dot file from dotman repository`,
		Run: func(cmd *cobra.Command, args []string) {
			subcommands.Remove(args)
		},
	}
	return removeSubCommand
}

func newPushSubCommand() *cobra.Command {
	var pushSubCommand = &cobra.Command{
		Use:   "push",
		Short: "Push changes to remote repository",
		Long:  `Push changes to remote repository`,
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			subcommands.Push(args, force)
		},
	}
	pushSubCommand.Flags().BoolP("force", "f", false, "Force push changes to remote repository")
	return pushSubCommand
}

func newSyncSubCommand() *cobra.Command {
	var syncSubCommand = &cobra.Command{
		Use:   "sync",
		Short: "Sync dotman repository with remote. It's a good idea to execute \nthis in your shell's startup file to ensure your dotfiles are always up to date.",
		Long:  `Sync dotman repository with remote. It's a good idea to execute \nthis in your shell's startup file to ensure your dotfiles are always up to date.`,
		Run: func(cmd *cobra.Command, args []string) {
			quiet, _ := cmd.Flags().GetBool("quiet")
			subcommands.Sync(args, quiet)
		},
	}
	syncSubCommand.Flags().BoolP("quiet", "q", false, "Suppress output messages")
	return syncSubCommand
}

func newListSubCommand() *cobra.Command {
	var listSubCommand = &cobra.Command{
		Use:   "list",
		Short: "List all managed dot files",
		Long:  `List all managed dot files`,
		Run: func(cmd *cobra.Command, args []string) {
			subcommands.List(args)
		},
	}
	return listSubCommand
}

func newDoctorSubCommand() *cobra.Command {
	var doctorSubCommand = &cobra.Command{
		Use:   "doctor",
		Short: "Check the health of your dotman setup",
		Long:  `Check the health of your dotman setup`,
		Run: func(cmd *cobra.Command, args []string) {
			fix, _ := cmd.Flags().GetBool("fix")
			subcommands.Doctor(args, fix)
		},
	}
	doctorSubCommand.Flags().BoolP("fix", "f", false, "Attempt to fix any issues found")
	return doctorSubCommand
}

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "dotman",
		Short: "Manage dot files with git behind the scenes",
		Long: `Simple CLI tool to manage your dot files using git as the backend.
It allows you to easily add, remove, push, and sync your dot files with a
remote repository to serve as the single source of truth.`,
	}

	rootCmd.AddCommand(
		newInitSubCommand(),
		newAddSubCommand(),
		newPushSubCommand(),
		newRemoveSubCommand(),
		newSyncSubCommand(),
		newListSubCommand(),
		newDoctorSubCommand(),
	)
	cobra.CheckErr(rootCmd.Execute())
}
