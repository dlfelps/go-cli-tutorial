package cmd

import (
        "fmt"
        "gocli-teacher/exercises"
        "gocli-teacher/progress"
        "os"

        "github.com/spf13/cobra"
)

// exerciseCmd represents the exercise command
var exerciseCmd = &cobra.Command{
        Use:   "exercise [name]",
        Short: "Complete an exercise to practice CLI development",
        Long: `The exercise command provides hands-on practice exercises 
for building CLI applications in Go.

Each exercise includes starter code that you can modify and run
to complete the task. The exercises build on the concepts from
the tutorials.

Available exercises:
  simple-cli       - Create a basic CLI tool
  flag-exercise    - Practice using command-line flags
  command-exercise - Create a CLI tool with subcommands
  interactive      - Build an interactive CLI application
`,
        Run: func(cmd *cobra.Command, args []string) {
                if len(args) == 0 {
                        fmt.Println("Please specify an exercise. For example:")
                        fmt.Println("  gocli-teacher exercise simple-cli")
                        fmt.Println("\nAvailable exercises:")
                        fmt.Println("  simple-cli, flag-exercise, command-exercise, interactive")
                        return
                }

                exercise := args[0]
                
                // Map exercise names to their internal names
                exerciseMap := map[string]string{
                        "simple-cli":       "simple_cli",
                        "simple_cli":       "simple_cli",
                        "flag-exercise":    "flag_exercise",
                        "flag_exercise":    "flag_exercise",
                        "command-exercise": "command_exercise",
                        "command_exercise": "command_exercise",
                        "interactive":      "interactive_exercise",
                }
                
                // Normalize exercise name
                normalizedExercise, exists := exerciseMap[exercise]
                if !exists {
                        fmt.Printf("Unknown exercise: %s\n", exercise)
                        fmt.Println("Available exercises: simple-cli, flag-exercise, command-exercise, interactive")
                        return
                }
                
                // Initialize progress tracker
                tracker, err := progress.New()
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Warning: Could not load progress data: %s\n", err)
                        // Continue without progress tracking
                }
                
                // Show progress info if exercise was completed before
                if tracker != nil && tracker.IsExerciseCompleted(normalizedExercise) {
                        fmt.Printf("\nNote: You've already completed this exercise. Running it again for practice.\n\n")
                }
                
                // Run the requested exercise
                var completed bool
                var score int
                
                switch normalizedExercise {
                case "simple_cli":
                        completed, score = exercises.RunSimpleCliExercise()
                case "flag_exercise":
                        completed, score = exercises.RunFlagExercise()
                case "command_exercise":
                        completed, score = exercises.RunCommandExercise()
                case "interactive_exercise":
                        completed, score = exercises.RunInteractiveExercise()
                }
                
                // Mark exercise as completed if successful
                if completed && tracker != nil {
                        if err := tracker.MarkExerciseComplete(normalizedExercise, score); err != nil {
                                fmt.Fprintf(os.Stderr, "Warning: Could not save progress: %s\n", err)
                        } else {
                                fmt.Printf("\nCongratulations! Exercise completed with score: %d/100\n", score)
                                
                                // Show recent progress
                                tutorials := getAllTutorials()
                                exercises := getAllExercises()
                                fmt.Print(progress.FormatRecentProgress(tracker, tutorials, exercises))
                        }
                }
        },
}

func init() {
        rootCmd.AddCommand(exerciseCmd)
}
