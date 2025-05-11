package utils

import (
        "fmt"
        "strings"
)

// PrintCodeBlock prints a code block with simple formatting
func PrintCodeBlock(code string) {
        fmt.Println("```")
        fmt.Println(code)
        fmt.Println("```")
}

// PrintCodeWithLineNumbers prints a code block with line numbers
func PrintCodeWithLineNumbers(code string) {
        lines := strings.Split(code, "\n")
        
        // Calculate the width needed for line numbers
        lineWidth := len(fmt.Sprintf("%d", len(lines)))
        format := fmt.Sprintf("%%%dd | %%s\n", lineWidth)
        
        fmt.Println("```")
        for i, line := range lines {
                fmt.Printf(format, i+1, line)
        }
        fmt.Println("```")
}

// FormatAsTable formats data as a simple ASCII table
func FormatAsTable(headers []string, rows [][]string) string {
        // Find the maximum width for each column
        colWidths := make([]int, len(headers))
        for i, header := range headers {
                colWidths[i] = len(header)
        }
        
        for _, row := range rows {
                for i, cell := range row {
                        if i < len(colWidths) && len(cell) > colWidths[i] {
                                colWidths[i] = len(cell)
                        }
                }
        }
        
        // Build the table
        var result strings.Builder
        
        // Headers
        result.WriteString("+")
        for _, width := range colWidths {
                result.WriteString(strings.Repeat("-", width+2))
                result.WriteString("+")
        }
        result.WriteString("\n")
        
        result.WriteString("|")
        for i, header := range headers {
                format := fmt.Sprintf(" %%-%ds |", colWidths[i])
                result.WriteString(fmt.Sprintf(format, header))
        }
        result.WriteString("\n")
        
        result.WriteString("+")
        for _, width := range colWidths {
                result.WriteString(strings.Repeat("-", width+2))
                result.WriteString("+")
        }
        result.WriteString("\n")
        
        // Rows
        for _, row := range rows {
                result.WriteString("|")
                for i, cell := range row {
                        if i < len(colWidths) {
                                format := fmt.Sprintf(" %%-%ds |", colWidths[i])
                                result.WriteString(fmt.Sprintf(format, cell))
                        }
                }
                result.WriteString("\n")
        }
        
        result.WriteString("+")
        for _, width := range colWidths {
                result.WriteString(strings.Repeat("-", width+2))
                result.WriteString("+")
        }
        result.WriteString("\n")
        
        return result.String()
}

// FormatAsList formats a list of items with bullets
func FormatAsList(items []string, indent int) string {
        var result strings.Builder
        indentStr := strings.Repeat(" ", indent)
        
        for _, item := range items {
                result.WriteString(indentStr + "• " + item + "\n")
        }
        
        return result.String()
}

// FormatAsSteps formats a list of items as numbered steps
func FormatAsSteps(steps []string) string {
        var result strings.Builder
        
        for i, step := range steps {
                result.WriteString(fmt.Sprintf("%d. %s\n", i+1, step))
        }
        
        return result.String()
}

// FormatNote formats a piece of text as a highlighted note
func FormatNote(text string) string {
        var result strings.Builder
        
        result.WriteString("\n")
        result.WriteString("┌" + strings.Repeat("─", len(text)+6) + "┐\n")
        result.WriteString("│   " + text + "   │\n")
        result.WriteString("└" + strings.Repeat("─", len(text)+6) + "┘\n")
        
        return result.String()
}

// FormatWarning formats a piece of text as a warning
func FormatWarning(text string) string {
        var result strings.Builder
        
        result.WriteString("\n")
        result.WriteString("⚠️  WARNING ⚠️\n")
        result.WriteString(text + "\n")
        result.WriteString("\n")
        
        return result.String()
}
