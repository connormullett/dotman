package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dotman",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

var initSubCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize dotman repository",
	Long:  `Initialize dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		Init(args)
	},
}

var addSubCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new dot file to dotman repository",
	Long:  `Add a new dot file to dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation of the add command
		Add(args)
	},
}

var removeSubCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a dot file from dotman repository",
	Long:  `Remove a dot file from dotman repository`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation of the remove command
		Remove(args)
	},
}

var pushSubCommand = &cobra.Command{
	Use:   "push",
	Short: "Push changes to remote repository",
	Long:  `Push changes to remote repository`,
	Run: func(cmd *cobra.Command, args []string) {
		// Implementation of the push command
		Push(args)
	},
}

func Execute() {
	rootCmd.AddCommand(initSubCommand, addSubCommand, pushSubCommand, removeSubCommand)
	cobra.CheckErr(rootCmd.Execute())
}
