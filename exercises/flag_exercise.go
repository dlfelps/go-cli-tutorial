package exercises

import (
	"fmt"
	"gocli-teacher/utils"
	"os"
	"path/filepath"
	"time"
)

const flagExerciseTemplate = `package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO: Define flags
	// - name: string flag for user's name (default: "World")
	// - uppercase: boolean flag to convert output to uppercase
	// - repeat: integer flag for number of times to repeat (default: 1)
	
	// TODO: Parse the flags
	
	// TODO: Generate greeting message
	// Format: "Hello, {name}!"
	
	// TODO: Apply uppercase conversion if the flag is set
	
	// TODO: Repeat the message based on the repeat flag
}
`

const flagExerciseSolution = `package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define flags
	namePtr := flag.String("name", "World", "your name")
	uppercasePtr := flag.Bool("uppercase", false, "convert output to uppercase")
	repeatPtr := flag.Int("repeat", 1, "number of times to repeat the message")
	
	// Parse the flags
	flag.Parse()
	
	// Generate greeting message
	message := fmt.Sprintf("Hello, %s!", *namePtr)
	
	// Apply uppercase conversion if the flag is set
	if *uppercasePtr {
		message = strings.ToUpper(message)
	}
	
	// Validate repeat count
	if *repeatPtr < 1 {
		fmt.Fprintln(os.Stderr, "Error: repeat count must be at least 1")
		os.Exit(1)
	}
	
	// Repeat the message
	for i := 0; i < *repeatPtr; i++ {
		fmt.Println(message)
	}
	
	// If any non-flag arguments were provided, print them
	if flag.NArg() > 0 {
		fmt.Println("\nAdditional arguments:")
		for i, arg := range flag.Args() {
			fmt.Printf("  %d: %s\n", i+1, arg)
		}
	}
}
`

// RunFlagExercise runs the flag exercise
func RunFlagExercise() {
	utils.ClearScreen()
	title := "Exercise: Working with Command-Line Flags"
	utils.PrintTitle(title)

	fmt.Println("Welcome to the Command-Line Flags exercise!")
	time.Sleep(1 * time.Second)
	
	fmt.Println("\nIn this exercise, you'll build a CLI tool that uses flags to customize output.")
	fmt.Println("You'll learn how to:")
	fmt.Println("1. Define different types of flags (string, boolean, integer)")
	fmt.Println("2. Parse and use flag values")
	fmt.Println("3. Handle default values and validation")
	fmt.Println("4. Process non-flag arguments")
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Exercise Instructions:")
	fmt.Println("")
	fmt.Println("You'll create a greeting CLI tool with the following flags:")
	fmt.Println("")
	fmt.Println("1. --name string")
	fmt.Println("   The name to greet (default: \"World\")")
	fmt.Println("")
	fmt.Println("2. --uppercase")
	fmt.Println("   Convert the output to uppercase")
	fmt.Println("")
	fmt.Println("3. --repeat int")
	fmt.Println("   Number of times to repeat the greeting (default: 1)")
	fmt.Println("")
	fmt.Println("The tool should generate a greeting message, apply any transformations,")
	fmt.Println("and repeat it the specified number of times.")
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Here's a template to get you started:")
	fmt.Println("")
	utils.PrintCodeWithLineNumbers(flagExerciseTemplate)
	
	// Create a directory for the exercise
	exerciseDir := "flag_exercise"
	err := os.MkdirAll(exerciseDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	
	// Create the exercise file
	exerciseFile := filepath.Join(exerciseDir, "main.go")
	err = os.WriteFile(exerciseFile, []byte(flagExerciseTemplate), 0644)
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
	fmt.Println("1. Define flags using the flag package:")
	utils.PrintCodeBlock(`namePtr := flag.String("name", "World", "your name")
uppercasePtr := flag.Bool("uppercase", false, "convert output to uppercase")
repeatPtr := flag.Int("repeat", 1, "number of times to repeat the message")`)
	
	fmt.Println("\n2. Parse flags before using them:")
	utils.PrintCodeBlock(`flag.Parse()`)
	
	fmt.Println("\n3. Access flag values using pointers:")
	utils.PrintCodeBlock(`message := fmt.Sprintf("Hello, %s!", *namePtr)`)
	
	fmt.Println("\n4. Convert a string to uppercase:")
	utils.PrintCodeBlock(`if *uppercasePtr {
    message = strings.ToUpper(message)
}`)
	
	fmt.Println("\n5. Access non-flag arguments:")
	utils.PrintCodeBlock(`if flag.NArg() > 0 {
    for i, arg := range flag.Args() {
        fmt.Printf("Arg %d: %s\n", i+1, arg)
    }
}`)
	
	utils.PressEnterToContinue()
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Testing Your Solution:")
	fmt.Println("")
	fmt.Println("Once you've completed the exercise, you can test it with these commands:")
	fmt.Println("")
	fmt.Println("1. Basic usage:")
	fmt.Println("   go run main.go")
	fmt.Println("   Expected: Hello, World!")
	fmt.Println("")
	fmt.Println("2. With name flag:")
	fmt.Println("   go run main.go --name Alice")
	fmt.Println("   Expected: Hello, Alice!")
	fmt.Println("")
	fmt.Println("3. With uppercase flag:")
	fmt.Println("   go run main.go --name Bob --uppercase")
	fmt.Println("   Expected: HELLO, BOB!")
	fmt.Println("")
	fmt.Println("4. With repeat flag:")
	fmt.Println("   go run main.go --name Charlie --repeat 3")
	fmt.Println("   Expected: Hello, Charlie! (repeated 3 times)")
	fmt.Println("")
	fmt.Println("5. With all flags and additional arguments:")
	fmt.Println("   go run main.go --name Dave --uppercase --repeat 2 extra args")
	fmt.Println("   Expected: HELLO, DAVE! (repeated 2 times)")
	fmt.Println("             Additional arguments: extra args")
	
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
		utils.PrintCodeWithLineNumbers(flagExerciseSolution)
		
		// Create the solution file
		solutionFile := filepath.Join(exerciseDir, "solution.go")
		err = os.WriteFile(solutionFile, []byte(flagExerciseSolution), 0644)
		if err != nil {
			fmt.Printf("Error creating solution file: %v\n", err)
		} else {
			fmt.Printf("\nI've saved the solution to %s\n", solutionFile)
		}
		
		utils.PressEnterToContinue()
	}
	
	utils.ClearScreen()
	utils.PrintTitle(title)
	
	fmt.Println("Congratulations on completing the Command-Line Flags exercise!")
	fmt.Println("\nWhat you've learned:")
	fmt.Println("1. How to define different types of flags")
	fmt.Println("2. How to parse and use flag values")
	fmt.Println("3. How to set default values and validate input")
	fmt.Println("4. How to access non-flag arguments")
	
	fmt.Println("\nNext steps:")
	fmt.Println("1. Try adding more flags with different types")
	fmt.Println("2. Experiment with flag.StringVar, flag.BoolVar, etc. for existing variables")
	fmt.Println("3. Try the 'command-exercise' to learn about command hierarchies")
	
	utils.PressEnterToContinue()
}
