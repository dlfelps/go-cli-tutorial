package tutorials

import (
        "fmt"
        "gocli-teacher/utils"
        "time"
)

// Basic CLI structure tutorial code
const basicCliExample = `package main

import (
        "fmt"
        "os"
)

func main() {
        // Check if arguments were provided
        if len(os.Args) < 2 {
                fmt.Println("Usage: myapp [command]")
                fmt.Println("Available commands: hello, version")
                os.Exit(1)
        }

        // Extract the command from arguments
        command := os.Args[1]

        // Process the command
        switch command {
        case "hello":
                fmt.Println("Hello, CLI world!")
        case "version":
                fmt.Println("v1.0.0")
        default:
                fmt.Printf("Unknown command: %s\n", command)
                fmt.Println("Available commands: hello, version")
                os.Exit(1)
        }
}
`

// RunBasicsTutorial runs the basics tutorial
func RunBasicsTutorial() bool {
        utils.ClearScreen()
        title := "CLI Basics in Go"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the basics of CLI development with Go!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this tutorial, you'll learn:")
        fmt.Println("1. The basic structure of a CLI application")
        fmt.Println("2. How to access command-line arguments")
        fmt.Println("3. How to handle basic commands")
        fmt.Println("4. Basic error handling")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's start with a simple CLI application:")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(basicCliExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Breaking down the key components:")
        fmt.Println("")
        fmt.Println("1. Accessing arguments:")
        utils.PrintCodeBlock(`// os.Args contains all command-line arguments
// os.Args[0] is the program name
// os.Args[1:] are the actual arguments
if len(os.Args) < 2 {
    fmt.Println("Usage: myapp [command]")
    os.Exit(1)
}
command := os.Args[1]`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Processing commands:")
        utils.PrintCodeBlock(`switch command {
case "hello":
    fmt.Println("Hello, CLI world!")
case "version":
    fmt.Println("v1.0.0")
default:
    fmt.Printf("Unknown command: %s\n", command)
    os.Exit(1)
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's try to run our example:")
        fmt.Println("\nIf we run: myapp hello")
        fmt.Println("Output: Hello, CLI world!")
        fmt.Println("\nIf we run: myapp version")
        fmt.Println("Output: v1.0.0")
        fmt.Println("\nIf we run: myapp")
        fmt.Println("Output: Usage: myapp [command]")
        fmt.Println("        Available commands: hello, version")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key takeaways:")
        fmt.Println("")
        fmt.Println("1. CLI apps in Go use the os.Args slice to access command line arguments")
        fmt.Println("2. Always check if required arguments are provided")
        fmt.Println("3. Provide helpful usage information when arguments are missing")
        fmt.Println("4. Use exit codes (os.Exit) to indicate success (0) or failure (non-zero)")
        
        fmt.Println("\nWhile this approach works for simple CLIs, more complex tools")
        fmt.Println("benefit from using dedicated CLI libraries like 'cobra' or 'urfave/cli'.")
        fmt.Println("These libraries make it easier to handle flags, nested commands, and more.")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's test your understanding with a few questions:")
        
        correct := 0
        
        // Question 1
        answer := utils.AskQuestion("\n1. How would you access the first user-provided argument in a Go CLI app?",
                []string{
                        "args[0]",
                        "os.Args[0]",
                        "os.Args[1]",
                        "flag.Arg(0)",
                })
        
        if answer == "os.Args[1]" {
                fmt.Println("Correct! os.Args[1] is the first user argument (os.Args[0] is the program name).")
                correct++
        } else {
                fmt.Println("Not quite. os.Args[1] is the first user argument (os.Args[0] is the program name).")
        }
        
        utils.PressEnterToContinue()
        
        // Question 2
        answer = utils.AskQuestion("\n2. What would be a good exit code for when a command succeeds?",
                []string{
                        "0",
                        "1",
                        "-1",
                        "Any non-zero value",
                })
        
        if answer == "0" {
                fmt.Println("Correct! By convention, 0 indicates success, while non-zero values indicate errors.")
                correct++
        } else {
                fmt.Println("Actually, by convention, 0 indicates success, while non-zero values indicate errors.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 3
        answer = utils.AskQuestion("\n3. What's the main limitation of using just os.Args for a complex CLI tool?",
                []string{
                        "It's too slow for processing many arguments",
                        "It doesn't handle flags and nested commands well",
                        "It only works on Unix systems",
                        "It has a limit of 10 arguments",
                })
        
        if answer == "It doesn't handle flags and nested commands well" {
                fmt.Println("Correct! For complex CLI tools with flags and subcommands, a library like 'cobra' is better.")
                correct++
        } else {
                fmt.Println("Not quite. The main limitation is that it doesn't handle flags and nested commands well.")
        }
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        // Final score
        fmt.Printf("You got %d out of 3 questions correct!\n\n", correct)
        
        fmt.Println("Congratulations on completing the CLI Basics tutorial!")
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try the 'simple-cli' exercise: gocli-teacher exercise simple-cli")
        fmt.Println("2. Learn about command-line flags: gocli-teacher tutorial flags")
        
        utils.PressEnterToContinue()
        
        // Consider the tutorial completed if the user got at least 2 out of 3 questions correct
        return correct >= 2
}
