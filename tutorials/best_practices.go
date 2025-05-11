package tutorials

import (
        "fmt"
        "gocli-teacher/utils"
        "time"
)

// RunBestPracticesTutorial runs the best practices tutorial
func RunBestPracticesTutorial() bool {
        utils.ClearScreen()
        title := "CLI Design Best Practices"
        utils.PrintTitle(title)

        fmt.Println("Welcome to the CLI Design Best Practices tutorial!")
        time.Sleep(1 * time.Second)
        
        fmt.Println("\nIn this tutorial, you'll learn:")
        fmt.Println("1. Principles of good CLI design")
        fmt.Println("2. Error handling best practices")
        fmt.Println("3. Documentation and help text guidelines")
        fmt.Println("4. Performance considerations")
        fmt.Println("5. Testing strategies for CLI applications")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Principles of Good CLI Design")
        fmt.Println("")
        fmt.Println("1. Follow the Unix Philosophy:")
        fmt.Println("   - Do one thing and do it well")
        fmt.Println("   - Compose small, focused tools to solve complex problems")
        fmt.Println("   - Process text streams as a universal interface")
        fmt.Println("")
        fmt.Println("2. Be Consistent:")
        fmt.Println("   - Use standard flag formats (--long-flag, -s)")
        fmt.Println("   - Follow established command naming patterns")
        fmt.Println("   - Be consistent with exit codes (0 for success, non-zero for errors)")
        fmt.Println("")
        fmt.Println("3. Respect the Platform:")
        fmt.Println("   - Use appropriate file paths for each OS")
        fmt.Println("   - Follow platform-specific conventions")
        fmt.Println("   - Handle terminal capabilities appropriately")
        fmt.Println("")
        fmt.Println("4. Provide Sensible Defaults:")
        fmt.Println("   - Tools should work reasonably without extensive configuration")
        fmt.Println("   - Default behavior should be safe and predictable")
        fmt.Println("   - Allow customization for advanced users")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Error Handling Best Practices")
        fmt.Println("")
        fmt.Println("1. Be Clear and Specific:")
        utils.PrintCodeBlock(`// Bad error handling
if err != nil {
    fmt.Println("Error occurred")
    os.Exit(1)
}

// Good error handling
if err != nil {
    fmt.Fprintf(os.Stderr, "Failed to read config file: %v\n", err)
    os.Exit(1)
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Use Appropriate Exit Codes:")
        utils.PrintCodeBlock(`const (
    ExitSuccess       = 0  // Success
    ExitError         = 1  // General error
    ExitUsageError    = 2  // Command line usage error
    ExitDataError     = 3  // Input data error
    ExitNoPermission  = 4  // Permission denied
    ExitNotFound      = 5  // Resource not found
)

// In your code
if _, err := os.Stat(configFile); os.IsNotExist(err) {
    fmt.Fprintf(os.Stderr, "Config file %s not found\n", configFile)
    os.Exit(ExitNotFound)
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Consider Verbosity Levels:")
        utils.PrintCodeBlock(`// Define a global verbosity flag
var verbose bool

// Use it to control output detail
if err != nil {
    if verbose {
        // Detailed error for debugging
        fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
        fmt.Fprintf(os.Stderr, "Stack trace: %s\n", debug.Stack())
    } else {
        // User-friendly error
        fmt.Fprintf(os.Stderr, "Unable to read file. Use --verbose for details.\n")
    }
    os.Exit(1)
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Documentation and Help Text")
        fmt.Println("")
        fmt.Println("1. Command Usage Information:")
        utils.PrintCodeBlock(`var rootCmd = &cobra.Command{
    Use:   "app [command]",        // Format: command name + arguments
    Short: "A brief description",  // One-line summary
    Long: "A longer description that spans multiple lines.\nExplain the purpose and basic usage of the application.\nProvide context for when this command should be used.",
    Example: "  app command arg1   // Short example comment\n  app command arg2   // Another example"
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Flag Descriptions:")
        utils.PrintCodeBlock(`// Good flag descriptions
rootCmd.Flags().StringVarP(&output, "output", "o", "", "output file path (defaults to stdout)")
rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose logging")
rootCmd.Flags().IntVarP(&retries, "retries", "r", 3, "number of retry attempts (0-10)")

// Bad flag descriptions
rootCmd.Flags().StringVarP(&output, "output", "o", "", "the output")
rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "verbosity")
rootCmd.Flags().IntVarP(&retries, "retries", "r", 3, "retries")`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Include Examples:")
        utils.PrintCodeBlock(`var convertCmd = &cobra.Command{
    Use:   "convert [input_file] [output_file]",
    Short: "Convert files between formats",
    Example: "  // Convert a JSON file to YAML\n  app convert data.json data.yaml\n  \n  // Convert and compress the output\n  app convert --compress data.json data.yaml",
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Performance Considerations")
        fmt.Println("")
        fmt.Println("1. Startup Time:")
        fmt.Println("   - CLI tools should start quickly, especially frequently used ones")
        fmt.Println("   - Defer expensive operations until needed")
        fmt.Println("   - Consider lazy loading for rarely used features")
        fmt.Println("")
        fmt.Println("2. Resource Usage:")
        fmt.Println("   - Be mindful of memory usage for large data processing")
        fmt.Println("   - Use streaming approaches when possible")
        fmt.Println("   - Consider the impact on the system")
        fmt.Println("")
        fmt.Println("3. Progress Feedback:")
        fmt.Println("   - Show progress for operations > 2 seconds")
        fmt.Println("   - Consider providing estimates for long-running tasks")
        fmt.Println("   - Allow safe interruption of long operations")
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Testing Strategies for CLI Applications")
        fmt.Println("")
        fmt.Println("1. Unit Testing:")
        utils.PrintCodeBlock(`func TestProcessData(t *testing.T) {
    input := []byte("{\\"name\\": \\"test\\"}")
    expected := []byte("name: test")
    
    output, err := processData(input, "json", "yaml")
    
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    
    if !bytes.Equal(output, expected) {
        t.Errorf("Expected %s, got %s", expected, output)
    }
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n2. Integration Testing:")
        utils.PrintCodeBlock(`func TestCLICommand(t *testing.T) {
    // Create a temporary directory for test files
    tempDir, err := ioutil.TempDir("", "cli-test")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tempDir)
    
    // Create test input file
    inputFile := filepath.Join(tempDir, "input.json")
    if err := ioutil.WriteFile(inputFile, []byte("{\\"name\\": \\"test\\"}"), 0644); err != nil {
        t.Fatal(err)
    }
    
    // Run the command
    output, err := exec.Command("./app", "convert", inputFile, "--format", "yaml").Output()
    if err != nil {
        t.Fatalf("Command failed: %v", err)
    }
    
    // Verify the output
    expected := "name: test\n"
    if string(output) != expected {
        t.Errorf("Expected output %q, got %q", expected, string(output))
    }
}`)
        
        utils.PressEnterToContinue()
        
        fmt.Println("\n3. Testing the Command Structure:")
        utils.PrintCodeBlock(`func TestRootCommand(t *testing.T) {
    // Save os.Args
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()
    
    // Test cases
    tests := []struct {
        args        []string
        expectedErr bool
    }{
        {[]string{"app"}, false},                           // Root command should work
        {[]string{"app", "--help"}, false},                 // Help should work
        {[]string{"app", "convert"}, true},                 // Missing required args
        {[]string{"app", "convert", "input.txt"}, false},   // Valid command
        {[]string{"app", "nonexistent"}, true},             // Invalid command
    }
    
    for _, tc := range tests {
        t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
            os.Args = tc.args
            
            err := cmd.Execute()
            
            if tc.expectedErr && err == nil {
                t.Error("Expected error but got none")
            }
            if !tc.expectedErr && err != nil {
                t.Errorf("Expected no error but got: %v", err)
            }
        })
    }
}`)
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        fmt.Println("Let's test your understanding:")
        
        correct := 0
        
        // Question 1
        answer := utils.AskQuestion("\n1. Which exit code should a CLI application use to indicate success?",
                []string{
                        "0",
                        "1",
                        "-1",
                        "Any non-zero value",
                })
        
        if answer == "0" {
                fmt.Println("Correct! By convention, exit code 0 indicates success.")
                correct++
        } else {
                fmt.Println("Not quite. By convention, exit code 0 indicates success, while non-zero values indicate errors.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 2
        answer = utils.AskQuestion("\n2. What's a good principle to follow when designing a CLI tool?",
                []string{
                        "Add as many features as possible",
                        "Make all flags required for clarity",
                        "Do one thing and do it well",
                        "Always use interactive prompts",
                })
        
        if answer == "Do one thing and do it well" {
                fmt.Println("Correct! This principle from the Unix philosophy encourages focused tools that can be composed together.")
                correct++
        } else {
                fmt.Println("Not quite. The Unix philosophy recommends 'Do one thing and do it well' for focused, composable tools.")
        }
        
        utils.PressEnterToContinue()
        
        // Question 3
        answer = utils.AskQuestion("\n3. Which of these is a good practice for error handling in CLI applications?",
                []string{
                        "Always exit with code 1 for any error",
                        "Print detailed stack traces for all errors",
                        "Provide specific, actionable error messages",
                        "Silently fail to avoid confusing the user",
                })
        
        if answer == "Provide specific, actionable error messages" {
                fmt.Println("Correct! Specific, actionable error messages help users understand and resolve issues.")
                correct++
        } else {
                fmt.Println("Not quite. The best practice is to provide specific, actionable error messages to help users resolve issues.")
        }
        
        utils.PressEnterToContinue()
        utils.ClearScreen()
        utils.PrintTitle(title)
        
        // Final score
        fmt.Printf("You got %d out of 3 questions correct!\n\n", correct)
        
        fmt.Println("Congratulations on completing the CLI Design Best Practices tutorial!")
        fmt.Println("\nYou've now completed all the tutorials in this CLI teaching tool!")
        fmt.Println("To continue your learning journey:")
        fmt.Println("1. Review tutorials you found challenging")
        fmt.Println("2. Try all the exercises")
        fmt.Println("3. Build your own CLI tool using what you've learned")
        
        utils.PressEnterToContinue()
        
        // Consider the tutorial completed if the user got at least 2 out of 3 questions correct
        return correct >= 2
}
