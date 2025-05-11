package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gocli-teacher",
	Short: "An interactive tool for learning how to build CLI applications in Go",
	Long: `GoCliTeacher - Learn to build CLI tools in Go

This application provides an interactive tutorial system to help you
learn how to build command-line applications using Go.

You'll learn about:
- Basic CLI structure
- Command-line flags and arguments
- Subcommands
- Interactive CLI features
- Error handling
- Best practices

Start by running one of the tutorial commands:
  gocli-teacher tutorial basics
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to GoCliTeacher!")
		fmt.Println("To get started, try running 'gocli-teacher tutorial basics'")
		fmt.Println("For help, run 'gocli-teacher --help'")
	},
}

var verbose bool

func init() {
	// Add global flags
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

// ExecuteWithArgs executes the root command with the given arguments for testing
func ExecuteWithArgs(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
