package cmd

import (
	"github.com/dotman/subcommands"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotman",
	Short: "Manage dot files with git behind the scenes",
	Long: `Simple CLI tool to manage your dot files using git as the backend.
It allows you to easily add, remove, push, and sync your dot files with a
remote repository to serve as the single source of truth.`,
}

var initSubCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize dotman repository",
	Long:  `Initialize dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.Init(args)
	},
}

var addSubCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new dot file to dotman repository",
	Long:  `Add a new dot file to dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.Add(args)
	},
}

var removeSubCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a dot file from dotman repository",
	Long:  `Remove a dot file from dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.Remove(args)
	},
}

var pushSubCommand = &cobra.Command{
	Use:   "push",
	Short: "Push changes to remote repository",
	Long:  `Push changes to remote repository`,
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := cmd.Flags().GetBool("force")
		subcommands.Push(args, force)
	},
}

var syncSubCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync dotman repository with remote. It's a good idea to execute \nthis in your shell's startup file to ensure your dotfiles are always up to date.",
	Long:  `Sync dotman repository with remote. It's a good idea to execute \nthis in your shell's startup file to ensure your dotfiles are always up to date.`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.Sync(args)
	},
}

var listSubCommand = &cobra.Command{
	Use:   "list",
	Short: "List all managed dot files",
	Long:  `List all managed dot files`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.List(args)
	},
}

var doctorSubCommand = &cobra.Command{
	Use:   "doctor",
	Short: "Check the health of your dotman setup",
	Long:  `Check the health of your dotman setup`,
	Run: func(cmd *cobra.Command, args []string) {
		subcommands.Doctor(args)
	},
}

func Execute() {
	rootCmd.AddCommand(
		initSubCommand,
		addSubCommand,
		pushSubCommand,
		removeSubCommand,
		syncSubCommand,
		listSubCommand,
		doctorSubCommand,
	)
	pushSubCommand.Flags().BoolP("force", "f", false, "Force push changes to remote repository")
	cobra.CheckErr(rootCmd.Execute())
}
