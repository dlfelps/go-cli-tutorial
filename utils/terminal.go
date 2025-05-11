package utils

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "runtime"
        "strings"
)

// TestMode indicates whether we are running in non-interactive test mode
var TestMode bool

// ClearScreen clears the terminal screen
func ClearScreen() {
        var cmd *exec.Cmd
        switch runtime.GOOS {
        case "linux", "darwin":
                cmd = exec.Command("clear")
        case "windows":
                cmd = exec.Command("cmd", "/c", "cls")
        default:
                // If the platform is not recognized, print newlines as a fallback
                fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
                return
        }

        cmd.Stdout = os.Stdout
        err := cmd.Run()
        if err != nil {
                // If command fails, print newlines as a fallback
                fmt.Print("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
        }
}

// PrintTitle prints a title with formatting
func PrintTitle(title string) {
        fmt.Println(strings.Repeat("=", len(title)+4))
        fmt.Println("  " + title)
        fmt.Println(strings.Repeat("=", len(title)+4))
        fmt.Println()
}

// PressEnterToContinue displays a prompt and waits for user to press Enter
// In test mode, it continues without user input
func PressEnterToContinue() {
        fmt.Print("\nPress Enter to continue...")
        if !TestMode {
                bufio.NewReader(os.Stdin).ReadBytes('\n')
        } else {
                fmt.Println(" (auto-continuing in test mode)")
        }
}

// AskYesNo asks a yes/no question and returns true for yes, false for no
func AskYesNo(question string) bool {
        reader := bufio.NewReader(os.Stdin)
        
        for {
                fmt.Printf("%s (y/n): ", question)
                input, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println("Error reading input:", err)
                        continue
                }
                
                input = strings.TrimSpace(strings.ToLower(input))
                
                switch input {
                case "y", "yes":
                        return true
                case "n", "no":
                        return false
                default:
                        fmt.Println("Please answer with 'y' or 'n'")
                }
        }
}

// AskQuestion asks a multiple-choice question and returns the selected answer
func AskQuestion(question string, options []string) string {
        fmt.Println(question)
        for i, option := range options {
                fmt.Printf("%d. %s\n", i+1, option)
        }
        
        reader := bufio.NewReader(os.Stdin)
        
        for {
                fmt.Print("\nEnter your choice (1-" + fmt.Sprint(len(options)) + "): ")
                input, err := reader.ReadString('\n')
                if err != nil {
                        fmt.Println("Error reading input:", err)
                        continue
                }
                
                input = strings.TrimSpace(input)
                choice := 0
                _, err = fmt.Sscanf(input, "%d", &choice)
                if err != nil || choice < 1 || choice > len(options) {
                        fmt.Println("Invalid choice. Please enter a number between 1 and", len(options))
                        continue
                }
                
                return options[choice-1]
        }
}

// AskForInput prompts for text input with an optional default value
func AskForInput(prompt string, defaultValue string) string {
        reader := bufio.NewReader(os.Stdin)
        
        if defaultValue != "" {
                fmt.Printf("%s [%s]: ", prompt, defaultValue)
        } else {
                fmt.Printf("%s: ", prompt)
        }
        
        input, err := reader.ReadString('\n')
        if err != nil {
                fmt.Println("Error reading input:", err)
                return defaultValue
        }
        
        input = strings.TrimSpace(input)
        if input == "" {
                return defaultValue
        }
        
        return input
}
