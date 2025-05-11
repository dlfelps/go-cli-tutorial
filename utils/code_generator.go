package utils

import (
        "fmt"
        "os"
        "path/filepath"
        "strings"
)

// CreateExerciseFile creates an exercise file with given content in the specified directory
func CreateExerciseFile(directory, filename, content string) (string, error) {
        // Ensure directory exists
        err := os.MkdirAll(directory, 0755)
        if err != nil {
                return "", fmt.Errorf("failed to create directory: %w", err)
        }
        
        // Create full path
        fullPath := filepath.Join(directory, filename)
        
        // Write content to file
        err = os.WriteFile(fullPath, []byte(content), 0644)
        if err != nil {
                return "", fmt.Errorf("failed to write file: %w", err)
        }
        
        return fullPath, nil
}

// ReadExerciseFile reads content from an exercise file
func ReadExerciseFile(filepath string) (string, error) {
        content, err := os.ReadFile(filepath)
        if err != nil {
                return "", fmt.Errorf("failed to read file: %w", err)
        }
        
        return string(content), nil
}

// ValidateGoCode performs a simple validation of Go code
// This is a basic validator that checks for common syntax issues
func ValidateGoCode(code string) (bool, []string) {
        var errors []string
        
        // Check for unbalanced curly braces
        openBraces := strings.Count(code, "{")
        closeBraces := strings.Count(code, "}")
        if openBraces != closeBraces {
                errors = append(errors, fmt.Sprintf("Unbalanced curly braces: %d open, %d close", openBraces, closeBraces))
        }
        
        // Check for unbalanced parentheses
        openParens := strings.Count(code, "(")
        closeParens := strings.Count(code, ")")
        if openParens != closeParens {
                errors = append(errors, fmt.Sprintf("Unbalanced parentheses: %d open, %d close", openParens, closeParens))
        }
        
        // Check for package declaration
        if !strings.Contains(code, "package ") {
                errors = append(errors, "Missing package declaration")
        }
        
        // More validations could be added here
        
        return len(errors) == 0, errors
}

// GenerateMainPackage generates a simple main package Go file
func GenerateMainPackage(appName, description string) string {
        template := `package main

import (
        "fmt"
        "os"
)

// %s - %s
func main() {
        // Check if arguments were provided
        if len(os.Args) < 2 {
                fmt.Println("Usage: %s [command]")
                fmt.Println("Run '%s help' for more information")
                os.Exit(1)
        }

        // Extract the command from arguments
        command := os.Args[1]

        // Process the command
        switch command {
        case "help":
                printHelp()
        default:
                fmt.Printf("Unknown command: %%s\n", command)
                printHelp()
                os.Exit(1)
        }
}

// printHelp displays usage information
func printHelp() {
        fmt.Println("%s - %s")
        fmt.Println()
        fmt.Println("Usage:")
        fmt.Println("  %s [command]")
        fmt.Println()
        fmt.Println("Available commands:")
        fmt.Println("  help    Display this help message")
        fmt.Println()
}
`
        return fmt.Sprintf(template, appName, description, appName, appName, appName, description, appName)
}

// GenerateCobraCommandFile generates a template for a Cobra command file
func GenerateCobraCommandFile(commandName, description string) string {
        template := `package cmd

import (
        "fmt"

        "github.com/spf13/cobra"
)

// %sCmd represents the %s command
var %sCmd = &cobra.Command{
        Use:   "%s",
        Short: "%s",
        Long: "%s

This command allows you to...",
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Running the %s command")
                // Your command logic here
        },
}

func init() {
        rootCmd.AddCommand(%sCmd)

        // Add command-specific flags here
        // %sCmd.Flags().StringVarP(&variable, "flag", "f", "default", "flag description")
}
`
        cmdVarName := strings.ToLower(commandName)
        if !strings.HasSuffix(cmdVarName, "Cmd") {
                cmdVarName += "Cmd"
        }
        
        return fmt.Sprintf(template, 
                cmdVarName, commandName, 
                cmdVarName, 
                commandName, description, description,
                commandName,
                cmdVarName,
                cmdVarName)
}

// GenerateReadme generates a simple README file for a CLI project
func GenerateReadme(appName, description string) string {
        template := `# %s

%s

## Installation

To install the application, run:

~~~bash
go get github.com/yourusername/%s
~~~`
        return fmt.Sprintf(template, appName, description, appName)
}
