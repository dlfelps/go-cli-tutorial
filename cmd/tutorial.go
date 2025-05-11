package cmd

import (
        "fmt"
        "gocli-teacher/progress"
        "gocli-teacher/tutorials"
        "os"

        "github.com/spf13/cobra"
)

// tutorialCmd represents the tutorial command
var tutorialCmd = &cobra.Command{
        Use:   "tutorial [topic]",
        Short: "Start an interactive tutorial on a specific topic",
        Long: `The tutorial command starts an interactive learning session
on a specific topic of CLI development with Go.

Available topics:
  basics         - Learn the basic structure of a CLI app in Go
  flags          - Understand how to use command-line flags
  commands       - Learn about subcommands and command hierarchy
  interactive    - Build interactive CLI features
  best-practices - Understand CLI design best practices
`,
        Run: func(cmd *cobra.Command, args []string) {
                if len(args) == 0 {
                        fmt.Println("Please specify a tutorial topic. For example:")
                        fmt.Println("  gocli-teacher tutorial basics")
                        fmt.Println("\nAvailable topics:")
                        fmt.Println("  basics, flags, commands, interactive, best-practices")
                        return
                }

                topic := args[0]
                
                // Map topics to their correct internal names
                topicMap := map[string]string{
                        "basics":         "basics",
                        "flags":          "flags",
                        "commands":       "commands",
                        "interactive":    "interactive",
                        "best-practices": "best_practices",
                        "best_practices": "best_practices",
                }
                
                // Normalize topic name
                normalizedTopic, exists := topicMap[topic]
                if !exists {
                        fmt.Printf("Unknown tutorial topic: %s\n", topic)
                        fmt.Println("Available topics: basics, flags, commands, interactive, best-practices")
                        return
                }
                
                // Check if tutorial was completed before
                tracker, err := progress.New()
                if err != nil {
                        fmt.Fprintf(os.Stderr, "Warning: Could not load progress data: %s\n", err)
                        // Continue without progress tracking
                }
                
                // Show progress info if tutorial was completed before
                if tracker != nil && tracker.IsTutorialCompleted(normalizedTopic) {
                        fmt.Printf("\nNote: You've already completed this tutorial. Running it again for review.\n\n")
                }
                
                // Run the requested tutorial
                var completed bool
                switch normalizedTopic {
                case "basics":
                        completed = tutorials.RunBasicsTutorial()
                case "flags":
                        completed = tutorials.RunFlagsTutorial()
                case "commands":
                        completed = tutorials.RunCommandsTutorial()
                case "interactive":
                        completed = tutorials.RunInteractiveTutorial()
                case "best_practices":
                        completed = tutorials.RunBestPracticesTutorial()
                }
                
                // Mark tutorial as completed if successful
                if completed && tracker != nil {
                        if err := tracker.MarkTutorialComplete(normalizedTopic); err != nil {
                                fmt.Fprintf(os.Stderr, "Warning: Could not save progress: %s\n", err)
                        } else {
                                fmt.Println("\nCongratulations! Tutorial completed and progress saved.")
                                
                                // Show recent progress
                                tutorials := getAllTutorials()
                                exercises := getAllExercises()
                                fmt.Print(progress.FormatRecentProgress(tracker, tutorials, exercises))
                        }
                }
        },
}

func init() {
        rootCmd.AddCommand(tutorialCmd)
}
