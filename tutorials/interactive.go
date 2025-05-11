package tutorials

import (
        "fmt"
        "gocli-teacher/utils"
        "time"
)

// Interactive prompt example
const interactivePromptExample = `package main

import (
        "fmt"
        "os"
        
        "github.com/AlecAivazis/survey/v2"
        "github.com/spf13/cobra"
)

func main() {
        var rootCmd = &cobra.Command{
                Use:   "interactive-demo",
                Short: "Demonstrates interactive CLI features",
                Run: func(cmd *cobra.Command, args []string) {
                        // Simple text input
                        name := ""
                        prompt := &survey.Input{
                                Message: "What is your name?",
                                Default: "User",
                        }
                        survey.AskOne(prompt, &name)
                        
                        // Multiple choice selection
                        color := ""
                        colorPrompt := &survey.Select{
                                Message: "Choose a color:",
                                Options: []string{"Red", "Green", "Blue", "Yellow"},
                        }
                        survey.AskOne(colorPrompt, &color)
                        
                        // Confirmation
                        confirm := false
                        confirmPrompt := &survey.Confirm{
                                Message: "Are you sure you want to continue?",
                                Default: true,
                        }
                        survey.AskOne(confirmPrompt, &confirm)
                        
                        // Using the collected data
                        fmt.Printf("Hello, %s! You selected %s.\n", name, color)
                        if confirm {
                                fmt.Println("Proceeding with the operation...")
                        } else {
                                fmt.Println("Operation cancelled.")
                        }
                },
        }
        
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// Progress bar example
const progressBarExample = `package main

import (
        "fmt"
        "os"
        "time"
        
        "github.com/spf13/cobra"
        "github.com/schollz/progressbar/v3"
)

func main() {
        var rootCmd = &cobra.Command{
                Use:   "progress-demo",
                Short: "Demonstrates progress bars in CLI",
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("Starting a long-running task...")
                        
                        // Create a new progress bar
                        bar := progressbar.NewOptions(100,
                                progressbar.OptionEnableColorCodes(true),
                                progressbar.OptionShowBytes(false),
                                progressbar.OptionSetWidth(15),
                                progressbar.OptionSetDescription("[cyan]Processing..."),
                                progressbar.OptionSetTheme(progressbar.Theme{
                                        Saucer:        "[green]=[reset]",
                                        SaucerHead:    "[green]>[reset]",
                                        SaucerPadding: " ",
                                        BarStart:      "[",
                                        BarEnd:        "]",
                                }))
                        
                        // Simulate work
                        for i := 0; i < 100; i++ {
                                bar.Add(1)
                                time.Sleep(50 * time.Millisecond)
                        }
                        
                        fmt.Println("\nTask completed successfully!")
                },
        }
        
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// Table output example
const tableOutputExample = `package main

import (
        "fmt"
        "os"
        
        "github.com/spf13/cobra"
        "github.com/olekukonko/tablewriter"
)

func main() {
        var rootCmd = &cobra.Command{
                Use:   "table-demo",
                Short: "Demonstrates formatted table output",
                Run: func(cmd *cobra.Command, args []string) {
                        // Sample data
                        data := [][]string{
                                {"1", "Alice", "Developer", "2018-01-15"},
                                {"2", "Bob", "Designer", "2019-03-20"},
                                {"3", "Charlie", "Manager", "2017-11-05"},
                                {"4", "Diana", "DevOps", "2020-05-12"},
                        }
                        
                        // Create a new table
                        table := tablewriter.NewWriter(os.Stdout)
                        table.SetHeader([]string{"ID", "Name", "Role", "Start Date"})
                        
                        // Optional styling
                        table.SetBorder(true)
                        table.SetRowLine(true)
                        table.SetCenterSeparator("|")
                        table.SetHeaderColor(
                                tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
                                tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
                                tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
                                tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
                        )
                        
                        // Add data and render
                        table.AppendBulk(data)
                        table.Render()
                },
        }
        
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}
`

// RunInteractiveTutorial runs the interactive CLI tutorial
func RunInteractiveTutorial() bool {
        utils.ClearScreen()
        title := "Interactive CLI Features"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the Interactive CLI Features tutorial!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this tutorial, you'll learn:")
        fmt.Println("1. How to create interactive prompts and collect user input")
        fmt.Println("2. How to display progress bars for long-running operations")
        fmt.Println("3. How to format data in tables and other visually appealing ways")
        fmt.Println("4. Best practices for creating user-friendly CLI interfaces")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Interactive Prompts with survey")
        fmt.Println("")
        fmt.Println("The 'survey' package provides a set of interactive prompts for CLI apps:")
        fmt.Println("- Text input")
        fmt.Println("- Password input")
        fmt.Println("- Multiple choice selection")
        fmt.Println("- Checkbox selection")
        fmt.Println("- Confirmation prompts")
        fmt.Println("- And more!")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(interactivePromptExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key Components of Interactive Prompts:")
        fmt.Println("")
        fmt.Println("1. Text Input:")
        utils.PrintCodeBlock(`name := ""
prompt := &survey.Input{
    Message: "What is your name?",
    Default: "User",
}
survey.AskOne(prompt, &name)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Multiple Choice Selection:")
        utils.PrintCodeBlock(`color := ""
colorPrompt := &survey.Select{
    Message: "Choose a color:",
    Options: []string{"Red", "Green", "Blue", "Yellow"},
}
survey.AskOne(colorPrompt, &color)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Confirmation Prompt:")
        utils.PrintCodeBlock(`confirm := false
confirmPrompt := &survey.Confirm{
    Message: "Are you sure you want to continue?",
    Default: true,
}
survey.AskOne(confirmPrompt, &confirm)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n4. Multiple Questions at Once:")
        utils.PrintCodeBlock(`// Define the questions
questions := []*survey.Question{
    {
        Name:     "name",
        Prompt:   &survey.Input{Message: "What is your name?"},
        Validate: survey.Required,
    },
    {
        Name: "color",
        Prompt: &survey.Select{
            Message: "Choose a color:",
            Options: []string{"Red", "Green", "Blue"},
        },
    },
}

// Ask them all and store the answers in a struct
answers := struct {
    Name  string
    Color string
}{}

survey.Ask(questions, &answers)`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Progress Bars for Long-Running Tasks")
        fmt.Println("")
        fmt.Println("Progress bars help users understand how long a task will take:")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(progressBarExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key Components of Progress Bars:")
        fmt.Println("")
        fmt.Println("1. Creating a Progress Bar:")
        utils.PrintCodeBlock(`// Create a new progress bar with 100 steps
bar := progressbar.NewOptions(100,
    progressbar.OptionEnableColorCodes(true),
    progressbar.OptionShowBytes(false),
    progressbar.OptionSetWidth(15),
    progressbar.OptionSetDescription("[cyan]Processing..."),
    progressbar.OptionSetTheme(progressbar.Theme{
        Saucer:        "[green]=[reset]",
        SaucerHead:    "[green]>[reset]",
        SaucerPadding: " ",
        BarStart:      "[",
        BarEnd:        "]",
    }))`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Updating the Progress Bar:")
        utils.PrintCodeBlock(`// For each step of work
bar.Add(1)

// Or update by a specific amount
bar.Add(10)

// Or set to a specific value
bar.Set(50)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Simpler Alternative with Default Options:")
        utils.PrintCodeBlock(`// Create a simple bar
bar := progressbar.Default(100)

// Update it in a loop
for i := 0; i < 100; i++ {
    bar.Add(1)
    time.Sleep(50 * time.Millisecond)
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Formatted Table Output")
        fmt.Println("")
        fmt.Println("Tables are great for displaying structured data:")
        fmt.Println("")
        utils.PrintCodeWithLineNumbers(tableOutputExample)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Key Components of Table Output:")
        fmt.Println("")
        fmt.Println("1. Creating a Table:")
        utils.PrintCodeBlock(`// Create a new table writer that outputs to stdout
table := tablewriter.NewWriter(os.Stdout)

// Set the table headers
table.SetHeader([]string{"ID", "Name", "Role", "Start Date"})`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Styling the Table:")
        utils.PrintCodeBlock(`// Add borders
table.SetBorder(true)

// Add lines between rows
table.SetRowLine(true)

// Change the separator character
table.SetCenterSeparator("|")

// Add colors to headers
table.SetHeaderColor(
    tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
    tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
    tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
    tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
)`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Adding Data and Rendering:")
        utils.PrintCodeBlock(`// Sample data as a slice of string slices
data := [][]string{
    {"1", "Alice", "Developer", "2018-01-15"},
    {"2", "Bob", "Designer", "2019-03-20"},
}

// Add all rows at once
table.AppendBulk(data)

// Or add one row at a time
table.Append([]string{"3", "Charlie", "Manager", "2017-11-05"})

// Render the table to output
table.Render()`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Best Practices for Interactive CLIs:")
        fmt.Println("")
        fmt.Println("1. Respect User Input:")
        fmt.Println("   - Provide sensible defaults")
        fmt.Println("   - Allow cancelling operations")
        fmt.Println("   - Validate input before proceeding")
        fmt.Println("")
        fmt.Println("2. Keep Users Informed:")
        fmt.Println("   - Show progress for long-running tasks")
        fmt.Println("   - Provide clear success/failure messages")
        fmt.Println("   - Use colors and formatting judiciously")
        fmt.Println("")
        fmt.Println("3. Be Consistent:")
        fmt.Println("   - Use similar patterns throughout your app")
        fmt.Println("   - Follow platform conventions")
        fmt.Println("   - Match the level of interaction to the task")
        fmt.Println("")
        fmt.Println("4. Consider Accessibility:")
        fmt.Println("   - Don't rely solely on colors for information")
        fmt.Println("   - Ensure your CLI works in different terminal types")
        fmt.Println("   - Provide non-interactive alternatives when appropriate")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's test your understanding:")
        
        correct := 0
        
        // Question 1
        answer := utils.AskQuestion("\n1. Which package is used for interactive prompts in the examples?",
                []string{
                        "github.com/spf13/prompt",
                        "github.com/AlecAivazis/survey/v2",
                        "github.com/olekukonko/prompter",
                        "github.com/interactive/cli",
                })
        
        if answer == "github.com/AlecAivazis/survey/v2" {
                fmt.Println("Correct! The survey package provides interactive prompt capabilities.")
                correct++
        } else {
                fmt.Println("Not quite. The survey package (github.com/AlecAivazis/survey/v2) is used for interactive prompts.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 2
        answer = utils.AskQuestion("\n2. When should you use a progress bar in a CLI application?",
                []string{
                        "For any operation that takes longer than 1 second",
                        "Only for file operations",
                        "For long-running operations where the user might wonder if the program is still working",
                        "Progress bars should be avoided in CLI applications",
                })
        
        if answer == "For long-running operations where the user might wonder if the program is still working" {
                fmt.Println("Correct! Progress bars help users understand that the program is still running during long operations.")
                correct++
        } else {
                fmt.Println("Not quite. Progress bars are best used for long-running operations where users might wonder if the program is still working.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 3
        answer = utils.AskQuestion("\n3. What method is used to display a table after setting it up with tablewriter?",
                []string{
                        "table.Show()",
                        "table.Display()",
                        "table.Render()",
                        "table.Print()",
                })
        
        if answer == "table.Render()" {
                fmt.Println("Correct! The Render() method is used to output the table to the specified writer.")
                correct++
        } else {
                fmt.Println("Not quite. The correct method is table.Render().")
        }
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        // Final score
        fmt.Printf("You got %d out of 3 questions correct!\n\n", correct)
        
        fmt.Println("Congratulations on completing the Interactive CLI Features tutorial!")
        fmt.Println("\nNext steps:")
        fmt.Println("1. Try the 'interactive-exercise': gocli-teacher exercise interactive")
        fmt.Println("2. Learn about CLI best practices: gocli-teacher tutorial best-practices")
        
        utils.PressEnterToContinue()
}
