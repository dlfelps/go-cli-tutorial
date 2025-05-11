package exercises

import (
        "fmt"
        "gocli-teacher/utils"
        "os"
        "path/filepath"
        "time"
)

const simpleCliTemplate = `package main

import (
        "fmt"
        "os"
)

func main() {
        // TODO: Check if arguments were provided
        // If no arguments are provided, print usage and exit
        
        // TODO: Extract the command from arguments
        
        // TODO: Process different commands (hello, echo, add)
        // "hello" - print "Hello, CLI world!"
        // "echo" - echo back all arguments after the command
        // "add" - convert the next two arguments to numbers and add them
}
`

const simpleCliSolution = `package main

import (
        "fmt"
        "os"
        "strconv"
)

func main() {
        // Check if arguments were provided
        if len(os.Args) < 2 {
                fmt.Println("Usage: simplecli [command] [args...]")
                fmt.Println("Available commands: hello, echo, add")
                os.Exit(1)
        }

        // Extract the command from arguments
        command := os.Args[1]

        // Process different commands
        switch command {
        case "hello":
                fmt.Println("Hello, CLI world!")
        
        case "echo":
                if len(os.Args) < 3 {
                        fmt.Println("Usage: simplecli echo [text to echo]")
                        os.Exit(1)
                }
                // Join all arguments after "echo" with spaces
                fmt.Println(strings.Join(os.Args[2:], " "))
        
        case "add":
                if len(os.Args) < 4 {
                        fmt.Println("Usage: simplecli add [number1] [number2]")
                        os.Exit(1)
                }
                
                // Convert arguments to numbers
                num1, err := strconv.Atoi(os.Args[2])
                if err != nil {
                        fmt.Printf("Error: %s is not a valid number\n", os.Args[2])
                        os.Exit(1)
                }
                
                num2, err := strconv.Atoi(os.Args[3])
                if err != nil {
                        fmt.Printf("Error: %s is not a valid number\n", os.Args[3])
                        os.Exit(1)
                }
                
                // Print the sum
                fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
        
        default:
                fmt.Printf("Unknown command: %s\n", command)
                fmt.Println("Available commands: hello, echo, add")
                os.Exit(1)
        }
}
`

// RunSimpleCliExercise runs the simple CLI exercise
func RunSimpleCliExercise() {
        utils.ClearScreen()
        title := "Exercise: Building a Simple CLI"
        utils.PrintTitle(title)

        fmt.Println("Welcome to your first CLI exercise!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this exercise, you'll build a simple CLI tool that can:")
        fmt.Println("1. Process different commands (hello, echo, add)")
        fmt.Println("2. Handle command-line arguments")
        fmt.Println("3. Provide helpful usage information")
        fmt.Println("4. Handle errors gracefully")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Exercise Instructions:")
        fmt.Println("")
        fmt.Println("You'll create a simple CLI tool with the following commands:")
        fmt.Println("")
        fmt.Println("1. hello - Prints 'Hello, CLI world!'")
        fmt.Println("   Usage: simplecli hello")
        fmt.Println("")
        fmt.Println("2. echo - Echoes back the provided arguments")
        fmt.Println("   Usage: simplecli echo [text to echo]")
        fmt.Println("")
        fmt.Println("3. add - Adds two numbers together")
        fmt.Println("   Usage: simplecli add [number1] [number2]")
        fmt.Println("")
        fmt.Println("The tool should handle missing arguments and show usage information.")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Here's a template to get you started:")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(simpleCliTemplate)
        
        // Create a directory for the exercise
        exerciseDir := "simple_cli_exercise"
        err := os.MkdirAll(exerciseDir, 0755)
        if err != nil {
                fmt.Printf("Error creating directory: %v\n", err)
                return
        }
        
        // Create the exercise file
        exerciseFile := filepath.Join(exerciseDir, "main.go")
        err = os.WriteFile(exerciseFile, []byte(simpleCliTemplate), 0644)
        if err != nil {
                fmt.Printf("Error creating file: %v\n", err)
                return
        }
        
        fmt.Printf("\nI've created a template file at %s\n", exerciseFile)
        fmt.Println("Edit this file to complete the exercise.")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Hints:")
        fmt.Println("")
        fmt.Println("1. Use os.Args to access command-line arguments")
        fmt.Println("   - os.Args[0] is the program name")
        fmt.Println("   - os.Args[1] should be the command (hello, echo, add)")
        fmt.Println("   - os.Args[2:] are additional arguments for the command")
        fmt.Println("")
        fmt.Println("2. Use a switch statement to handle different commands")
        fmt.Println("")
        fmt.Println("3. For the 'add' command, you'll need to:")
        fmt.Println("   - Convert string arguments to integers (strconv.Atoi)")
        fmt.Println("   - Handle conversion errors")
        fmt.Println("")
        fmt.Println("4. Use os.Exit(1) for error conditions")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Testing Your Solution:")
        fmt.Println("")
        fmt.Println("Once you've completed the exercise, you can test it with these commands:")
        fmt.Println("")
        fmt.Println("1. Basic usage:")
        fmt.Println("   go run main.go")
        fmt.Println("   Expected: Usage information")
        fmt.Println("")
        fmt.Println("2. Hello command:")
        fmt.Println("   go run main.go hello")
        fmt.Println("   Expected: Hello, CLI world!")
        fmt.Println("")
        fmt.Println("3. Echo command:")
        fmt.Println("   go run main.go echo Hello there!")
        fmt.Println("   Expected: Hello there!")
        fmt.Println("")
        fmt.Println("4. Add command:")
        fmt.Println("   go run main.go add 5 7")
        fmt.Println("   Expected: 5 + 7 = 12")
        
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
                utils.PrintCodeWithLineNumbers(simpleCliSolution)
                
                // Create the solution file
                solutionFile := filepath.Join(exerciseDir, "solution.go")
                err = os.WriteFile(solutionFile, []byte(simpleCliSolution), 0644)
                if err != nil {
                        fmt.Printf("Error creating solution file: %v\n", err)
                } else {
                        fmt.Printf("\nI've saved the solution to %s\n", solutionFile)
                }
                
                utils.PressEnterToContinue()
        }
        
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Congratulations on completing the Simple CLI exercise!")
        fmt.Println("\nWhat you've learned:")
        fmt.Println("1. How to access and process command-line arguments")
        fmt.Println("2. How to implement different commands in a CLI tool")
        fmt.Println("3. How to provide usage information and handle errors")
        fmt.Println("4. How to convert string arguments to other types")
        
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try adding more commands to your CLI tool")
        fmt.Println("2. Add validation for input arguments")
        fmt.Println("3. Try the 'flag-exercise' to learn about command-line flags")
        
        utils.PressEnterToContinue()
}
