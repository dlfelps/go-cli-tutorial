package tutorials

import (
        "fmt"
        "gocli-teacher/utils"
        "time"
)

// Standard library flag example
const stdFlagExample = `package main

import (
        "flag"
        "fmt"
)

func main() {
        // Define flags
        namePtr := flag.String("name", "World", "a name to say hello to")
        agePtr := flag.Int("age", 0, "age of the person")
        verbosePtr := flag.Bool("verbose", false, "enable verbose output")
        
        // Parse the flags
        flag.Parse()
        
        // Use the flags
        if *verbosePtr {
                fmt.Printf("Name is %s\n", *namePtr)
                fmt.Printf("Age is %d\n", *agePtr)
        }
        
        fmt.Printf("Hello, %s!\n", *namePtr)
        
        // Remaining arguments (after flags)
        if flag.NArg() > 0 {
                fmt.Println("Additional arguments:")
                for i, arg := range flag.Args() {
                        fmt.Printf("  Arg %d: %s\n", i+1, arg)
                }
        }
}
`

// Cobra flag example
const cobraFlagExample = `package main

import (
        "fmt"
        "os"
        
        "github.com/spf13/cobra"
)

func main() {
        var (
                name    string
                age     int
                verbose bool
        )
        
        // Create the root command
        rootCmd := &cobra.Command{
                Use:   "greet [args]",
                Short: "A friendly greeting CLI",
                Run: func(cmd *cobra.Command, args []string) {
                        if verbose {
                                fmt.Printf("Name is %s\n", name)
                                fmt.Printf("Age is %d\n", age)
                        }
                        
                        fmt.Printf("Hello, %s!\n", name)
                        
                        if len(args) > 0 {
                                fmt.Println("Additional arguments:")
                                for i, arg := range args {
                                        fmt.Printf("  Arg %d: %s\n", i+1, arg)
                                }
                        }
                },
        }
        
        // Add flags to the command
        rootCmd.Flags().StringVarP(&name, "name", "n", "World", "a name to say hello to")
        rootCmd.Flags().IntVarP(&age, "age", "a", 0, "age of the person")
        rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
        
        // Execute the command
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// RunFlagsTutorial runs the flags tutorial
func RunFlagsTutorial() bool {
        utils.ClearScreen()
        title := "Command Line Flags in Go"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the Command Line Flags tutorial!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this tutorial, you'll learn:")
        fmt.Println("1. What command-line flags are and why they're useful")
        fmt.Println("2. How to implement flags using Go's standard library")
        fmt.Println("3. How to implement flags using the Cobra library")
        fmt.Println("4. Best practices for working with flags")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("What are command-line flags?")
        fmt.Println("")
        fmt.Println("Flags are named parameters that modify the behavior of a command.")
        fmt.Println("They typically start with one or two dashes and can be provided in any order.")
        fmt.Println("")
        fmt.Println("Examples:")
        fmt.Println("  --name=John     (long form with equals sign)")
        fmt.Println("  --name John     (long form with space)")
        fmt.Println("  -n John         (short form)")
        fmt.Println("  --verbose       (boolean flag)")
        fmt.Println("  -v              (short boolean flag)")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Implementing Flags with Go's Standard Library")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(stdFlagExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key components of using the standard flag package:")
        fmt.Println("")
        fmt.Println("1. Defining flags:")
        utils.PrintCodeBlock(`// Define different types of flags
namePtr := flag.String("name", "World", "a name to say hello to")
agePtr := flag.Int("age", 0, "age of the person")
verbosePtr := flag.Bool("verbose", false, "enable verbose output")

// Alternative syntax for existing variables
var name string
flag.StringVar(&name, "name", "World", "a name to say hello to")`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Parsing flags:")
        utils.PrintCodeBlock(`// Must be called after flags are defined and before they're used
flag.Parse()

// Always use after Parse() since flags return pointers
fmt.Printf("Hello, %s!\n", *namePtr)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Accessing remaining arguments:")
        utils.PrintCodeBlock(`// Check if there are arguments after the flags
if flag.NArg() > 0 {
    // flag.Args() returns non-flag arguments
    for i, arg := range flag.Args() {
        fmt.Printf("  Arg %d: %s\n", i+1, arg)
    }
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("How to run the standard flag example:")
        fmt.Println("")
        fmt.Println("Basic usage:")
        fmt.Println("  ./myprog")
        fmt.Println("  Output: Hello, World!")
        fmt.Println("")
        fmt.Println("With flags:")
        fmt.Println("  ./myprog --name=John --age=30 --verbose")
        fmt.Println("  Output: Name is John")
        fmt.Println("          Age is 30")
        fmt.Println("          Hello, John!")
        fmt.Println("")
        fmt.Println("With flags and additional arguments:")
        fmt.Println("  ./myprog --name=John extra args here")
        fmt.Println("  Output: Hello, John!")
        fmt.Println("          Additional arguments:")
        fmt.Println("            Arg 1: extra")
        fmt.Println("            Arg 2: args")
        fmt.Println("            Arg 3: here")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Implementing Flags with Cobra")
        fmt.Println("")
        fmt.Println("Cobra is a popular library for creating powerful CLI applications in Go.")
        fmt.Println("It offers more features than the standard library, including:")
        fmt.Println("- Nested subcommands")
        fmt.Println("- Automatic help generation")
        fmt.Println("- Shell autocompletion")
        fmt.Println("- Better flag handling")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(cobraFlagExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key components of using Cobra for flags:")
        fmt.Println("")
        fmt.Println("1. Creating a command:")
        utils.PrintCodeBlock(`rootCmd := &cobra.Command{
    Use:   "greet [args]",
    Short: "A friendly greeting CLI",
    Run: func(cmd *cobra.Command, args []string) {
        // Command logic goes here
        fmt.Printf("Hello, %s!\n", name)
    },
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Adding flags to a command:")
        utils.PrintCodeBlock(`// PersistentFlags are inherited by subcommands
rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "World", "a name to say hello to")

// Local flags only apply to this command
rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

// The P in StringVarP, BoolVarP, etc. allows for short flag forms (-n vs --name)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Executing the command:")
        utils.PrintCodeBlock(`// Execute parses flags and runs the command
if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Flag Best Practices:")
        fmt.Println("")
        fmt.Println("1. Provide sensible defaults for optional flags")
        fmt.Println("2. Use short flags (-v) for common options, long flags (--verbose) for all")
        fmt.Println("3. Be consistent with flag naming conventions")
        fmt.Println("4. Provide clear, concise help text for each flag")
        fmt.Println("5. Use appropriate flag types (string, int, bool, etc.)")
        fmt.Println("6. Handle invalid flag values gracefully")
        fmt.Println("7. Consider required vs. optional flags")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's test your understanding:")
        
        correct := 0
        
        // Question 1
        answer := utils.AskQuestion("\n1. Which function must be called before using flags from the standard library?",
                []string{
                        "flag.Define()",
                        "flag.Parse()",
                        "flag.Init()",
                        "flag.Execute()",
                })
        
        if answer == "flag.Parse()" {
                fmt.Println("Correct! flag.Parse() must be called after defining flags and before using them.")
                correct++
        } else {
                fmt.Println("Not quite. flag.Parse() must be called after defining flags and before using them.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 2
        answer = utils.AskQuestion("\n2. In Cobra, what's the difference between PersistentFlags and Flags?",
                []string{
                        "PersistentFlags are required, Flags are optional",
                        "PersistentFlags are inherited by subcommands, Flags only apply to the current command",
                        "PersistentFlags accept all value types, Flags only accept strings",
                        "There is no difference",
                })
        
        if answer == "PersistentFlags are inherited by subcommands, Flags only apply to the current command" {
                fmt.Println("Correct! PersistentFlags are inherited by all subcommands in the hierarchy.")
                correct++
        } else {
                fmt.Println("Not quite. PersistentFlags are inherited by all subcommands, while Flags only apply to the current command.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 3
        answer = utils.AskQuestion("\n3. What does the P in StringVarP stand for in Cobra?",
                []string{
                        "Pointer",
                        "Parameter",
                        "Persistent",
                        "It allows for short flag forms",
                })
        
        if answer == "It allows for short flag forms" {
                fmt.Println("Correct! The P variants let you specify both long (--name) and short (-n) forms of a flag.")
                correct++
        } else {
                fmt.Println("Not quite. The P variants let you specify both long (--name) and short (-n) forms of a flag.")
        }
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        // Final score
        fmt.Printf("You got %d out of 3 questions correct!\n\n", correct)
        
        fmt.Println("Congratulations on completing the Command Line Flags tutorial!")
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try the 'flag-exercise': gocli-teacher exercise flag-exercise")
        fmt.Println("2. Learn about commands: gocli-teacher tutorial commands")
        
        utils.PressEnterToContinue()
        
        // Consider the tutorial completed if the user got at least 2 out of 3 questions correct
        return correct >= 2
}
