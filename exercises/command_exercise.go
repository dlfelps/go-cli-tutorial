package exercises

import (
        "fmt"
        "gocli-teacher/utils"
        "os"
        "path/filepath"
        "time"
)

const commandExerciseTemplate = `package main

import (
        "fmt"
        "os"
        
        "github.com/spf13/cobra"
)

func main() {
        // TODO: Create the root command
        // The root command should print a welcome message and usage information
        
        // TODO: Create a "greet" command
        // The greet command should accept a --name flag and print a greeting
        
        // TODO: Create a "calc" command
        // The calc command should serve as a parent for calculation subcommands
        
        // TODO: Create "calc add" and "calc multiply" subcommands
        // Each should accept two number arguments and perform the respective operation
        
        // TODO: Execute the root command
}
`

const commandExerciseSolution = `package main

import (
        "fmt"
        "os"
        "strconv"
        
        "github.com/spf13/cobra"
)

func main() {
        // Create the root command
        rootCmd := &cobra.Command{
                Use:   "multicmd",
                Short: "A CLI tool with multiple commands",
                Long:  "A CLI tool demonstrating command hierarchy with Cobra",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Welcome to the multi-command tool!")
                        fmt.Println("Use --help to see available commands")
                },
        }
        
        // Create a "greet" command
        var name string
        greetCmd := &cobra.Command{
                Use:   "greet",
                Short: "Greet a person",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Printf("Hello, %s!\n", name)
                },
        }
        
        // Add flags to greet command
        greetCmd.Flags().StringVarP(&name, "name", "n", "World", "name of the person to greet")
        
        // Add greet command to root
        rootCmd.AddCommand(greetCmd)
        
        // Create a "calc" command
        calcCmd := &cobra.Command{
                Use:   "calc",
                Short: "Perform calculations",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Calculator commands:")
                        fmt.Println("  add       - Add two numbers")
                        fmt.Println("  multiply  - Multiply two numbers")
                        fmt.Println("\nUse 'multicmd calc [command] --help' for more information")
                },
        }
        
        // Add calc command to root
        rootCmd.AddCommand(calcCmd)
        
        // Create "calc add" subcommand
        addCmd := &cobra.Command{
                Use:   "add [number1] [number2]",
                Short: "Add two numbers",
                Args:  cobra.ExactArgs(2),
                Run: func(cmd *cobra.Command, args []string) {
                        // Convert arguments to numbers
                        num1, err := strconv.ParseFloat(args[0], 64)
                        if err != nil {
                                fmt.Printf("Error: %s is not a valid number\n", args[0])
                                os.Exit(1)
                        }
                        
                        num2, err := strconv.ParseFloat(args[1], 64)
                        if err != nil {
                                fmt.Printf("Error: %s is not a valid number\n", args[1])
                                os.Exit(1)
                        }
                        
                        // Print the sum
                        fmt.Printf("%g + %g = %g\n", num1, num2, num1+num2)
                },
        }
        
        // Create "calc multiply" subcommand
        multiplyCmd := &cobra.Command{
                Use:   "multiply [number1] [number2]",
                Short: "Multiply two numbers",
                Args:  cobra.ExactArgs(2),
                Run: func(cmd *cobra.Command, args []string) {
                        // Convert arguments to numbers
                        num1, err := strconv.ParseFloat(args[0], 64)
                        if err != nil {
                                fmt.Printf("Error: %s is not a valid number\n", args[0])
                                os.Exit(1)
                        }
                        
                        num2, err := strconv.ParseFloat(args[1], 64)
                        if err != nil {
                                fmt.Printf("Error: %s is not a valid number\n", args[1])
                                os.Exit(1)
                        }
                        
                        // Print the product
                        fmt.Printf("%g × %g = %g\n", num1, num2, num1*num2)
                },
        }
        
        // Add subcommands to calc command
        calcCmd.AddCommand(addCmd)
        calcCmd.AddCommand(multiplyCmd)
        
        // Execute the root command
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// RunCommandExercise runs the command exercise
func RunCommandExercise() (bool, int) {
        // Variables to track progress
        var completed int
        totalTasks := 4 // Total number of tasks in this exercise
        requiredTasks := 3 // Minimum tasks required to consider exercise completed
        utils.ClearScreen()
        title := "Exercise: Command Hierarchy with Cobra"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the Command Hierarchy exercise!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this exercise, you'll build a CLI tool with a command hierarchy using Cobra.")
        fmt.Println("You'll learn how to:")
        fmt.Println("1. Create a root command")
        fmt.Println("2. Add subcommands to create a command hierarchy")
        fmt.Println("3. Add command-specific flags")
        fmt.Println("4. Handle command arguments")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Exercise Instructions:")
        fmt.Println("")
        fmt.Println("You'll create a CLI tool with the following structure:")
        fmt.Println("")
        fmt.Println("multicmd                - The root command (shows welcome message)")
        fmt.Println("  |- greet              - Greets a person (has --name flag)")
        fmt.Println("  |- calc               - Parent for calculation commands")
        fmt.Println("      |- add            - Adds two numbers")
        fmt.Println("      |- multiply       - Multiplies two numbers")
        fmt.Println("")
        fmt.Println("Usage examples:")
        fmt.Println("  multicmd")
        fmt.Println("  multicmd greet --name Alice")
        fmt.Println("  multicmd calc add 5 7")
        fmt.Println("  multicmd calc multiply 3 4")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Here's a template to get you started:")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(commandExerciseTemplate)
        
        fmt.Println("\nNote: This exercise requires the Cobra package.")
        fmt.Println("Make sure to run 'go get github.com/spf13/cobra' before starting.")
        
        // Create a directory for the exercise
        exerciseDir := "command_exercise"
        err := os.MkdirAll(exerciseDir, 0755)
        if err != nil {
                fmt.Printf("Error creating directory: %v\n", err)
                return false, 0
        }
        
        // Create the exercise file
        exerciseFile := filepath.Join(exerciseDir, "main.go")
        err = os.WriteFile(exerciseFile, []byte(commandExerciseTemplate), 0644)
        if err != nil {
                fmt.Printf("Error creating file: %v\n", err)
                return false, 0
        }
        
        fmt.Printf("\nI've created a template file at %s\n", exerciseFile)
        fmt.Println("Edit this file to complete the exercise.")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Hints:")
        fmt.Println("")
        fmt.Println("1. Creating the root command:")
        utils.PrintCodeBlock(`rootCmd := &cobra.Command{
    Use:   "multicmd",
    Short: "A CLI tool with multiple commands",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to the multi-command tool!")
        fmt.Println("Use --help to see available commands")
    },
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Creating a subcommand with a flag:")
        utils.PrintCodeBlock(`var name string
greetCmd := &cobra.Command{
    Use:   "greet",
    Short: "Greet a person",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Hello, %s!\n", name)
    },
}

// Add flag to the command
greetCmd.Flags().StringVarP(&name, "name", "n", "World", "name of the person to greet")

// Add command to the root
rootCmd.AddCommand(greetCmd)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Creating a command with subcommands:")
        utils.PrintCodeBlock(`// Parent command
calcCmd := &cobra.Command{
    Use:   "calc",
    Short: "Perform calculations",
    Run: func(cmd *cobra.Command, args []string) {
        // This runs when 'calc' is called without subcommands
        fmt.Println("Use a calc subcommand: add or multiply")
    },
}

// Add to root
rootCmd.AddCommand(calcCmd)

// Add subcommands to calcCmd
calcCmd.AddCommand(addCmd)
calcCmd.AddCommand(multiplyCmd)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n4. Command with required arguments:")
        utils.PrintCodeBlock(`addCmd := &cobra.Command{
    Use:   "add [number1] [number2]",
    Short: "Add two numbers",
    Args:  cobra.ExactArgs(2),  // Requires exactly 2 arguments
    Run: func(cmd *cobra.Command, args []string) {
        // Convert and use args[0] and args[1]
    },
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Testing Your Solution:")
        fmt.Println("")
        fmt.Println("Once you've completed the exercise, you can test it with these commands:")
        fmt.Println("")
        fmt.Println("1. Root command:")
        fmt.Println("   go run main.go")
        fmt.Println("   Expected: Welcome message")
        fmt.Println("")
        fmt.Println("2. Greet command:")
        fmt.Println("   go run main.go greet")
        fmt.Println("   Expected: Hello, World!")
        fmt.Println("")
        fmt.Println("   go run main.go greet --name Alice")
        fmt.Println("   Expected: Hello, Alice!")
        fmt.Println("")
        fmt.Println("3. Calc commands:")
        fmt.Println("   go run main.go calc")
        fmt.Println("   Expected: List of available calc subcommands")
        fmt.Println("")
        fmt.Println("   go run main.go calc add 5 7")
        fmt.Println("   Expected: 5 + 7 = 12")
        fmt.Println("")
        fmt.Println("   go run main.go calc multiply 3 4")
        fmt.Println("   Expected: 3 × 4 = 12")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Need the solution?")
        showSolution := utils.AskYesNo("Would you like to see the solution?")
        
        if showSolution {
                utils.ClearScreen()
                utils.PrintTitle(title + " - Solution")
                
                fmt.Println("Here's one way to solve the exercise:")
                fmt.Println("")
                utils.PrintCodeWithLineNumbers(commandExerciseSolution)
                
                // Create the solution file
                solutionFile := filepath.Join(exerciseDir, "solution.go")
                err = os.WriteFile(solutionFile, []byte(commandExerciseSolution), 0644)
                if err != nil {
                        fmt.Printf("Error creating solution file: %v\n", err)
                } else {
                        fmt.Printf("\nI've saved the solution to %s\n", solutionFile)
                }
                
                utils.PressEnterToContinue()
        }
        
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Congratulations on completing the Command Hierarchy exercise!")
        fmt.Println("\nWhat you've learned:")
        fmt.Println("1. How to create a command hierarchy with Cobra")
        fmt.Println("2. How to add command-specific flags")
        fmt.Println("3. How to validate command arguments")
        fmt.Println("4. How to organize related functionality in subcommands")
        
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try adding more subcommands to the hierarchy")
        fmt.Println("2. Add global flags that apply to all commands")
        fmt.Println("3. Try the 'interactive-exercise' to learn about interactive CLI features")
        
        utils.PressEnterToContinue()
        
        // Calculate score based on completed tasks (max 100)
        score := completed * (100 / totalTasks)
        if score > 100 {
            score = 100
        }
        
        return completed >= requiredTasks, score
}
