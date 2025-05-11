package tutorials

import (
        "fmt"
        "gocli-teacher/utils"
        "time"
)

// Command structure example
const commandStructureExample = `package main

import (
        "fmt"
        "os"

        "github.com/spf13/cobra"
)

func main() {
        // Create the root command
        var rootCmd = &cobra.Command{
                Use:   "myapp",
                Short: "MyApp is a CLI application example",
                Long:  "MyApp demonstrates how to structure a CLI app with commands and subcommands",
                Run: func(cmd *cobra.Command, args []string) {
                        // This code runs when no subcommand is specified
                        fmt.Println("Welcome to MyApp! Use --help to see available commands.")
                },
        }

        // Add a 'version' command
        var versionCmd = &cobra.Command{
                Use:   "version",
                Short: "Print the version number",
                Long:  "Print the version number of MyApp",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("MyApp v1.0.0")
                },
        }
        rootCmd.AddCommand(versionCmd)

        // Add a 'user' command with subcommands
        var userCmd = &cobra.Command{
                Use:   "user",
                Short: "User management commands",
                Long:  "Commands for managing users in the system",
                Run: func(cmd *cobra.Command, args []string) {
                        cmd.Help()
                },
        }
        rootCmd.AddCommand(userCmd)

        // Add 'user list' subcommand
        var userListCmd = &cobra.Command{
                Use:   "list",
                Short: "List all users",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Listing all users...")
                        // User listing logic would go here
                },
        }
        userCmd.AddCommand(userListCmd)

        // Add 'user add' subcommand with flags
        var (
                userName  string
                userEmail string
        )
        var userAddCmd = &cobra.Command{
                Use:   "add",
                Short: "Add a new user",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Printf("Adding user: %s (%s)\n", userName, userEmail)
                        // User creation logic would go here
                },
        }
        userAddCmd.Flags().StringVarP(&userName, "name", "n", "", "User's full name")
        userAddCmd.Flags().StringVarP(&userEmail, "email", "e", "", "User's email address")
        userAddCmd.MarkFlagRequired("name")
        userAddCmd.MarkFlagRequired("email")
        userCmd.AddCommand(userAddCmd)

        // Execute the root command
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// RunCommandsTutorial runs the commands tutorial
func RunCommandsTutorial() bool {
        utils.ClearScreen()
        title := "Commands and Subcommands in CLI Applications"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the Commands and Subcommands tutorial!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this tutorial, you'll learn:")
        fmt.Println("1. How to structure a CLI application with commands")
        fmt.Println("2. How to create subcommands for a hierarchical command structure")
        fmt.Println("3. How to handle command-specific flags")
        fmt.Println("4. Best practices for command organization")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Command Structure in CLI Applications")
        fmt.Println("")
        fmt.Println("Modern CLI tools often use a command-subcommand structure:")
        fmt.Println("")
        fmt.Println("  rootCommand subcommand [flags] [args]")
        fmt.Println("")
        fmt.Println("Examples:")
        fmt.Println("  git commit -m \"message\"")
        fmt.Println("  docker container list")
        fmt.Println("  kubectl get pods --namespace default")
        fmt.Println("")
        fmt.Println("This structure allows for:")
        fmt.Println("- Logical grouping of related functionality")
        fmt.Println("- Better discoverability through help text")
        fmt.Println("- Command-specific flags and arguments")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Command Structure Example with Cobra")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(commandStructureExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key Components of Command Structure:")
        fmt.Println("")
        fmt.Println("1. Creating the Root Command:")
        utils.PrintCodeBlock(`var rootCmd = &cobra.Command{
    Use:   "myapp",
    Short: "MyApp is a CLI application example",
    Long:  "MyApp demonstrates how to structure a CLI app with commands and subcommands",
    Run: func(cmd *cobra.Command, args []string) {
        // This code runs when no subcommand is specified
        fmt.Println("Welcome to MyApp! Use --help to see available commands.")
    },
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Adding Subcommands:")
        utils.PrintCodeBlock(`// Create a subcommand
var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("MyApp v1.0.0")
    },
}

// Add it to the parent command
rootCmd.AddCommand(versionCmd)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Creating Command Hierarchies:")
        utils.PrintCodeBlock(`// Parent command
var userCmd = &cobra.Command{
    Use:   "user",
    Short: "User management commands",
}
rootCmd.AddCommand(userCmd)

// Child command (becomes 'myapp user list')
var userListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all users",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing all users...")
    },
}
userCmd.AddCommand(userListCmd)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n4. Command-Specific Flags:")
        utils.PrintCodeBlock(`var userName string
var userAddCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new user",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Adding user: %s\n", userName)
    },
}

// Add flags to the specific command
userAddCmd.Flags().StringVarP(&userName, "name", "n", "", "User's full name")

// Mark a flag as required
userAddCmd.MarkFlagRequired("name")`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Using Our Example CLI")
        fmt.Println("")
        fmt.Println("Our example creates a CLI with this structure:")
        fmt.Println("")
        fmt.Println("  myapp             - The root command")
        fmt.Println("    |- version      - Prints version info")
        fmt.Println("    |- user         - User management")
        fmt.Println("        |- list     - Lists users")
        fmt.Println("        |- add      - Adds a user (requires flags)")
        fmt.Println("")
        fmt.Println("Example usage:")
        fmt.Println("")
        fmt.Println("  $ myapp")
        fmt.Println("  Welcome to MyApp! Use --help to see available commands.")
        fmt.Println("")
        fmt.Println("  $ myapp version")
        fmt.Println("  MyApp v1.0.0")
        fmt.Println("")
        fmt.Println("  $ myapp user list")
        fmt.Println("  Listing all users...")
        fmt.Println("")
        fmt.Println("  $ myapp user add --name \"John Doe\" --email \"john@example.com\"")
        fmt.Println("  Adding user: John Doe (john@example.com)")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Command Design Best Practices:")
        fmt.Println("")
        fmt.Println("1. Command Naming:")
        fmt.Println("   - Use clear, descriptive verb-noun pairs (e.g., 'add user', not 'user-add')")
        fmt.Println("   - Be consistent with naming patterns")
        fmt.Println("   - Use common conventions (list, create, delete, etc.)")
        fmt.Println("")
        fmt.Println("2. Command Organization:")
        fmt.Println("   - Group related commands under parent commands")
        fmt.Println("   - Limit nesting to 2-3 levels for usability")
        fmt.Println("   - Most used commands should be easier to access")
        fmt.Println("")
        fmt.Println("3. Help and Documentation:")
        fmt.Println("   - Provide concise 'Short' descriptions (shown in command lists)")
        fmt.Println("   - Provide detailed 'Long' descriptions (shown in command help)")
        fmt.Println("   - Include examples in help text")
        fmt.Println("")
        fmt.Println("4. Error Handling:")
        fmt.Println("   - Use appropriate exit codes")
        fmt.Println("   - Provide helpful error messages")
        fmt.Println("   - Consider implementing a --debug flag for verbose errors")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's test your understanding:")
        
        correct := 0
        
        // Question 1
        answer := utils.AskQuestion("\n1. What function is used to add a subcommand to a parent command in Cobra?",
                []string{
                        "parentCmd.AttachCommand(childCmd)",
                        "parentCmd.AddCommand(childCmd)",
                        "childCmd.SetParent(parentCmd)",
                        "cobra.Connect(parentCmd, childCmd)",
                })
        
        if answer == "parentCmd.AddCommand(childCmd)" {
                fmt.Println("Correct! AddCommand() is used to add a subcommand to a parent command.")
                correct++
        } else {
                fmt.Println("Not quite. The correct function is parentCmd.AddCommand(childCmd).")
        }
        
        utils.PressEnterToContinue()
        
        // Question 2
        answer = utils.AskQuestion("\n2. What happens if you run a command with no Run function defined?",
                []string{
                        "It displays an error",
                        "It does nothing",
                        "It automatically displays the help text",
                        "It runs the parent command's Run function",
                })
        
        if answer == "It automatically displays the help text" {
                fmt.Println("Correct! Cobra automatically displays help for commands with no Run function.")
                correct++
        } else {
                fmt.Println("Actually, Cobra automatically displays help for commands with no Run function.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 3
        answer = utils.AskQuestion("\n3. How do you mark a flag as required in Cobra?",
                []string{
                        "flag.Required = true",
                        "cmd.RequireFlag(\"flagname\")",
                        "cmd.MarkFlagRequired(\"flagname\")",
                        "flag.SetRequired(true)",
                })
        
        if answer == "cmd.MarkFlagRequired(\"flagname\")" {
                fmt.Println("Correct! MarkFlagRequired is used to mark a flag as required in Cobra.")
                correct++
        } else {
                fmt.Println("Not quite. The correct method is cmd.MarkFlagRequired(\"flagname\").")
        }
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        // Final score
        fmt.Printf("You got %d out of 3 questions correct!\n\n", correct)
        
        fmt.Println("Congratulations on completing the Commands and Subcommands tutorial!")
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try the 'command-exercise': gocli-teacher exercise command-exercise")
        fmt.Println("2. Learn about interactive CLI features: gocli-teacher tutorial interactive")
        
        utils.PressEnterToContinue()
        
        // Consider the tutorial completed if the user got at least 2 out of 3 questions correct
        return correct >= 2
}
