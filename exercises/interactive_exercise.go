package exercises

import (
	"fmt"
	"gocli-teacher/utils"
	"os"
	"path/filepath"
	"time"
)

const interactiveExerciseTemplate = `package main

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// TODO: Create the root command
	
	// TODO: Create an "interactive" command with subcommands
	
	// TODO: Create "interactive form" subcommand
	// This should collect user information (name, age, favorite color) using survey
	
	// TODO: Create "interactive choose" subcommand
	// This should present a multiple choice selection and act on the choice
	
	// TODO: Create "progress" command
	// This should simulate a long-running task with a progress bar
	
	// TODO: Execute the root command
}
`

const interactiveExerciseSolution = `package main

import (
	"fmt"
	"os"
	"time"
	
	"github.com/spf13/cobra"
	"github.com/AlecAivazis/survey/v2"
	"github.com/schollz/progressbar/v3"
)

func main() {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "interactive-cli",
		Short: "A demo of interactive CLI features",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Interactive CLI Demo!")
			fmt.Println("Run 'interactive-cli --help' to see available commands.")
		},
	}
	
	// Create an "interactive" command
	interactiveCmd := &cobra.Command{
		Use:   "interactive",
		Short: "Interactive command examples",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Interactive command subcommands:")
			fmt.Println("  form    - Collect information via a form")
			fmt.Println("  choose  - Make a selection from options")
			fmt.Println("\nUse 'interactive-cli interactive [command]' to run a subcommand")
		},
	}
	rootCmd.AddCommand(interactiveCmd)
	
	// Create "interactive form" subcommand
	formCmd := &cobra.Command{
		Use:   "form",
		Short: "Collect information via interactive prompts",
		Run: func(cmd *cobra.Command, args []string) {
			// Define the questions
			questions := []*survey.Question{
				{
					Name: "name",
					Prompt: &survey.Input{
						Message: "What is your name?",
						Default: "User",
					},
					Validate: survey.Required,
				},
				{
					Name: "age",
					Prompt: &survey.Input{
						Message: "How old are you?",
					},
				},
				{
					Name: "color",
					Prompt: &survey.Select{
						Message: "Choose your favorite color:",
						Options: []string{"Red", "Green", "Blue", "Yellow", "Purple"},
						Default: "Blue",
					},
				},
				{
					Name: "hobbies",
					Prompt: &survey.MultiSelect{
						Message: "Select your hobbies:",
						Options: []string{
							"Reading",
							"Programming",
							"Sports",
							"Music",
							"Gaming",
							"Cooking",
						},
					},
				},
			}
			
			// Answers struct
			answers := struct {
				Name    string
				Age     string
				Color   string
				Hobbies []string
			}{}
			
			// Ask the questions
			err := survey.Ask(questions, &answers)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			
			// Display the answers
			fmt.Println("\nYour information:")
			fmt.Println("------------------")
			fmt.Printf("Name: %s\n", answers.Name)
			fmt.Printf("Age: %s\n", answers.Age)
			fmt.Printf("Favorite color: %s\n", answers.Color)
			
			fmt.Println("Hobbies:")
			if len(answers.Hobbies) == 0 {
				fmt.Println("  No hobbies selected")
			} else {
				for _, hobby := range answers.Hobbies {
					fmt.Printf("  - %s\n", hobby)
				}
			}
		},
	}
	interactiveCmd.AddCommand(formCmd)
	
	// Create "interactive choose" subcommand
	chooseCmd := &cobra.Command{
		Use:   "choose",
		Short: "Make a selection from options",
		Run: func(cmd *cobra.Command, args []string) {
			// Options for the user to choose from
			choice := ""
			prompt := &survey.Select{
				Message: "What would you like to do?",
				Options: []string{
					"Show the current time",
					"Show a greeting",
					"Flip a coin",
					"Exit",
				},
				Default: "Show a greeting",
			}
			
			// Ask for the selection
			survey.AskOne(prompt, &choice)
			
			// Process the choice
			switch choice {
			case "Show the current time":
				fmt.Printf("The current time is: %s\n", time.Now().Format("15:04:05"))
				
			case "Show a greeting":
				name := ""
				namePrompt := &survey.Input{
					Message: "What is your name?",
					Default: "friend",
				}
				survey.AskOne(namePrompt, &name)
				fmt.Printf("Hello, %s! It's nice to meet you.\n", name)
				
			case "Flip a coin":
				options := []string{"Heads", "Tails"}
				result := options[time.Now().UnixNano()%2]
				fmt.Printf("The coin shows: %s\n", result)
				
			case "Exit":
				fmt.Println("Goodbye!")
			}
		},
	}
	interactiveCmd.AddCommand(chooseCmd)
	
	// Create "progress" command
	progressCmd := &cobra.Command{
		Use:   "progress",
		Short: "Demonstrate a progress bar",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting a simulated task...")
			
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
				time.Sleep(30 * time.Millisecond)
			}
			
			fmt.Println("\nTask completed successfully!")
		},
	}
	rootCmd.AddCommand(progressCmd)
	
	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`

// RunInteractiveExercise runs the interactive CLI exercise
func RunInteractiveExercise() {
	utils.ClearScreen()
	title := "Exercise: Interactive CLI Features"
	utils.PrintTitle(title)

	fmt.Println("Welcome to the Interactive CLI Features exercise!")
	time.Sleep(1 * time.Second)
	
	fmt.Println("\nIn this exercise, you'll build a CLI tool with interactive features.")
	fmt.Println("You'll learn how to:")
	fmt.Println("1. Create interactive prompts to collect user input")
	fmt.Println("2. Present selection menus for user choices")
	fmt.Println("3. Display progress bars for long-running tasks")
	fmt.Println("4. Combine interactive elements with a command structure")
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Exercise Instructions:")
	fmt.Println("")
	fmt.Println("You'll create a CLI tool with the following structure:")
	fmt.Println("")
	fmt.Println("interactive-cli                  - The root command")
	fmt.Println("  |- interactive                 - Parent for interactive commands")
	fmt.Println("      |- form                    - Collect user info via prompts")
	fmt.Println("      |- choose                  - Present options and act on selection")
	fmt.Println("  |- progress                    - Show a progress bar demonstration")
	fmt.Println("")
	fmt.Println("The 'form' command should collect:")
	fmt.Println("- Name (text input)")
	fmt.Println("- Age (text input)")
	fmt.Println("- Favorite color (selection)")
	fmt.Println("- Hobbies (multi-selection)")
	fmt.Println("")
	fmt.Println("The 'choose' command should present options and perform different")
	fmt.Println("actions based on the selected option.")
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Here's a template to get you started:")
	fmt.Println("")
	utils.PrintCodeWithLineNumbers(interactiveExerciseTemplate)
	
	fmt.Println("\nNote: This exercise requires additional packages:")
	fmt.Println("- github.com/spf13/cobra")
	fmt.Println("- github.com/AlecAivazis/survey/v2")
	fmt.Println("- github.com/schollz/progressbar/v3")
	
	// Create a directory for the exercise
	exerciseDir := "interactive_exercise"
	err := os.MkdirAll(exerciseDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	
	// Create the exercise file
	exerciseFile := filepath.Join(exerciseDir, "main.go")
	err = os.WriteFile(exerciseFile, []byte(interactiveExerciseTemplate), 0644)
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
	fmt.Println("1. Creating an interactive form:")
	utils.PrintCodeBlock(`// Define the questions
questions := []*survey.Question{
    {
        Name: "name",
        Prompt: &survey.Input{
            Message: "What is your name?",
            Default: "User",
        },
        Validate: survey.Required,
    },
    {
        Name: "color",
        Prompt: &survey.Select{
            Message: "Choose your favorite color:",
            Options: []string{"Red", "Green", "Blue"},
            Default: "Blue",
        },
    },
}

// Answers struct
answers := struct {
    Name  string
    Color string
}{}

// Ask the questions
err := survey.Ask(questions, &answers)
if err != nil {
    fmt.Println("Error:", err)
    return
}

// Use the answers
fmt.Printf("Hello, %s! Your favorite color is %s.\n", answers.Name, answers.Color)`)
	
	utils.PressEnterToContinue()
	
	fmt.Println("\n2. Creating a selection menu and handling the choice:")
	utils.PrintCodeBlock(`// Options for the user to choose from
choice := ""
prompt := &survey.Select{
    Message: "What would you like to do?",
    Options: []string{
        "Show the current time",
        "Show a greeting",
        "Exit",
    },
}

// Ask for the selection
survey.AskOne(prompt, &choice)

// Process the choice
switch choice {
case "Show the current time":
    fmt.Printf("The current time is: %s\n", time.Now().Format("15:04:05"))
    
case "Show a greeting":
    fmt.Println("Hello, world!")
    
case "Exit":
    fmt.Println("Goodbye!")
}`)
	
	utils.PressEnterToContinue()
	
	fmt.Println("\n3. Creating a progress bar for a long-running task:")
	utils.PrintCodeBlock(`// Create a new progress bar
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

fmt.Println("\nTask completed successfully!")`)
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Testing Your Solution:")
	fmt.Println("")
	fmt.Println("Once you've completed the exercise, you can test it with these commands:")
	fmt.Println("")
	fmt.Println("1. Interactive form:")
	fmt.Println("   go run main.go interactive form")
	fmt.Println("   Expected: A series of prompts collecting information")
	fmt.Println("")
	fmt.Println("2. Interactive choice:")
	fmt.Println("   go run main.go interactive choose")
	fmt.Println("   Expected: A menu of options to select from")
	fmt.Println("")
	fmt.Println("3. Progress bar:")
	fmt.Println("   go run main.go progress")
	fmt.Println("   Expected: A progress bar for a simulated task")
	
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
		utils.PrintCodeWithLineNumbers(interactiveExerciseSolution)
		
		// Create the solution file
		solutionFile := filepath.Join(exerciseDir, "solution.go")
		err = os.WriteFile(solutionFile, []byte(interactiveExerciseSolution), 0644)
		if err != nil {
			fmt.Printf("Error creating solution file: %v\n", err)
		} else {
			fmt.Printf("\nI've saved the solution to %s\n", solutionFile)
		}
		
		utils.PressEnterToContinue()
	}
	
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Congratulations on completing the Interactive CLI Features exercise!")
	fmt.Println("\nWhat you've learned:")
	fmt.Println("1. How to create interactive prompts with the survey package")
	fmt.Println("2. How to collect different types of input (text, selection, multi-selection)")
	fmt.Println("3. How to display progress for long-running tasks")
	fmt.Println("4. How to combine interactive elements with a command structure")
	
	fmt.Println("\nNext steps:")
	fmt.Println("1. Add validation to the form inputs")
	fmt.Println("2. Add more interactive elements (password input, confirmation, etc.)")
	fmt.Println("3. Create a more complex application combining all the techniques you've learned")
	
	utils.PressEnterToContinue()
}
