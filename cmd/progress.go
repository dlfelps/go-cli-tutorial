package cmd

import (
	"fmt"
	"gocli-teacher/progress"
	"os"

	"github.com/spf13/cobra"
)

// progressCmd represents the progress command
var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "View your learning progress",
	Long: `View your progress through tutorials and exercises.
This command shows which tutorials and exercises you've completed,
and provides statistics on your overall progress.`,
	Run: func(cmd *cobra.Command, args []string) {
		showProgress(cmd, args)
	},
}

// reset flag for progress command
var resetProgress bool
var showRecent bool

func init() {
	rootCmd.AddCommand(progressCmd)

	// Add flags
	progressCmd.Flags().BoolVarP(&resetProgress, "reset", "r", false, "Reset all progress tracking data")
	progressCmd.Flags().BoolVarP(&showRecent, "recent", "", false, "Show only recent activity")
}

func showProgress(cmd *cobra.Command, args []string) {
	// Initialize progress tracker
	tracker, err := progress.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to load progress data: %s\n", err)
		os.Exit(1)
	}

	// Check if reset flag is set
	if resetProgress {
		if err := tracker.Reset(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to reset progress: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Progress data has been reset.")
		return
	}

	// Get all available tutorials and exercises
	tutorials := getAllTutorials()
	exercises := getAllExercises()

	// Display progress
	if showRecent {
		fmt.Print(progress.FormatRecentProgress(tracker, tutorials, exercises))
	} else {
		fmt.Print(progress.FormatAllProgress(tracker, tutorials, exercises))
	}
}

// getAllTutorials returns info about all available tutorials
func getAllTutorials() []progress.TutorialInfo {
	return []progress.TutorialInfo{
		{Name: "basics", Description: "Basic CLI structure and command line arguments"},
		{Name: "flags", Description: "Working with command line flags"},
		{Name: "commands", Description: "Creating and organizing subcommands"},
		{Name: "interactive", Description: "Building interactive CLI applications"},
		{Name: "best_practices", Description: "Best practices for CLI development"},
	}
}

// getAllExercises returns info about all available exercises
func getAllExercises() []progress.ExerciseInfo {
	return []progress.ExerciseInfo{
		{Name: "simple_cli", Description: "Build a simple CLI tool", Difficulty: "Easy"},
		{Name: "flag_exercise", Description: "Create a CLI with multiple flags", Difficulty: "Medium"},
		{Name: "command_exercise", Description: "Implement a CLI with subcommands", Difficulty: "Medium"},
		{Name: "interactive_exercise", Description: "Build an interactive CLI", Difficulty: "Hard"},
	}
}