package progress

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// TutorialInfo contains info about a tutorial
type TutorialInfo struct {
	Name        string
	Description string
}

// ExerciseInfo contains info about an exercise
type ExerciseInfo struct {
	Name        string
	Description string
	Difficulty  string
}

// FormatAllProgress formats the user's overall progress
func FormatAllProgress(tracker *Tracker, tutorials []TutorialInfo, exercises []ExerciseInfo) string {
	// Calculate overall progress
	completed, total, percentage := tracker.GetProgress(len(tutorials), len(exercises))
	
	// Format the progress bar
	progressBar := createProgressBar(percentage, 40)
	
	// Build the output
	var sb strings.Builder
	sb.WriteString("\n==========================================\n")
	sb.WriteString("             Learning Progress\n")
	sb.WriteString("==========================================\n\n")
	sb.WriteString(fmt.Sprintf("Overall Progress: %d/%d (%.1f%%)\n", completed, total, percentage))
	sb.WriteString(fmt.Sprintf("%s\n\n", progressBar))
	
	// Add tutorial progress
	completedTutorials := tracker.GetCompletedTutorials()
	sb.WriteString(fmt.Sprintf("Tutorials: %d/%d\n", len(completedTutorials), len(tutorials)))
	sb.WriteString("------------------------------------------\n")
	
	for _, tut := range tutorials {
		status := "[ ]"
		if tracker.IsTutorialCompleted(tut.Name) {
			status = "[✓]"
		}
		sb.WriteString(fmt.Sprintf("%s %s: %s\n", status, tut.Name, tut.Description))
	}
	
	sb.WriteString("\n")
	
	// Add exercise progress
	completedExercises := tracker.GetCompletedExercises()
	sb.WriteString(fmt.Sprintf("Exercises: %d/%d\n", len(completedExercises), len(exercises)))
	sb.WriteString("------------------------------------------\n")
	
	for _, ex := range exercises {
		status := "[ ]"
		if tracker.IsExerciseCompleted(ex.Name) {
			status = "[✓]"
		}
		sb.WriteString(fmt.Sprintf("%s %s (%s): %s\n", status, ex.Name, ex.Difficulty, ex.Description))
	}
	
	return sb.String()
}

// FormatRecentProgress formats information about recent progress
func FormatRecentProgress(tracker *Tracker, tutorials []TutorialInfo, exercises []ExerciseInfo) string {
	// Get progress data
	var sb strings.Builder
	
	completed, total, percentage := tracker.GetProgress(len(tutorials), len(exercises))
	
	// Format the progress bar
	progressBar := createProgressBar(percentage, 30)
	
	sb.WriteString("\n=================================\n")
	sb.WriteString("       Recent Activity\n")
	sb.WriteString("=================================\n\n")
	sb.WriteString(fmt.Sprintf("Overall Progress: %d/%d (%.1f%%)\n", completed, total, percentage))
	sb.WriteString(fmt.Sprintf("%s\n\n", progressBar))
	
	// Get recently completed items (for a real implementation, this would sort by completion date)
	// For now, we'll just show completed items
	recentCount := 3 // Show up to 3 recent completions
	
	// Create a list of recently completed items
	type CompletedItem struct {
		Type        string // "Tutorial" or "Exercise"
		Name        string
		Description string
		CompletedAt time.Time
	}
	
	var completedItems []CompletedItem
	
	// Add tutorials
	for _, tut := range tutorials {
		if status, exists := tracker.data.Tutorials[tut.Name]; exists && status.Completed {
			completedItems = append(completedItems, CompletedItem{
				Type:        "Tutorial",
				Name:        tut.Name,
				Description: tut.Description,
				CompletedAt: status.CompletedAt,
			})
		}
	}
	
	// Add exercises
	for _, ex := range exercises {
		if status, exists := tracker.data.Exercises[ex.Name]; exists && status.Completed {
			completedItems = append(completedItems, CompletedItem{
				Type:        "Exercise",
				Name:        ex.Name,
				Description: ex.Description,
				CompletedAt: status.CompletedAt,
			})
		}
	}
	
	// Sort by completion date (most recent first)
	sort.Slice(completedItems, func(i, j int) bool {
		return completedItems[i].CompletedAt.After(completedItems[j].CompletedAt)
	})
	
	// Limit to recent items
	if len(completedItems) > recentCount {
		completedItems = completedItems[:recentCount]
	}
	
	if len(completedItems) > 0 {
		sb.WriteString("Recently Completed:\n")
		sb.WriteString("---------------------------------\n")
		
		for _, item := range completedItems {
			timeStr := item.CompletedAt.Format("Jan 02, 2006 15:04:05")
			sb.WriteString(fmt.Sprintf("[%s] %s: %s\n", item.Type, item.Name, item.Description))
			sb.WriteString(fmt.Sprintf("    Completed: %s\n", timeStr))
		}
	} else {
		sb.WriteString("No completed tutorials or exercises yet.\n")
		sb.WriteString("Start by running 'gocli-teacher tutorial basics'\n")
	}
	
	sb.WriteString("\n=================================\n")
	
	return sb.String()
}

// createProgressBar returns a string representation of a progress bar
func createProgressBar(percentage float64, width int) string {
	// Calculate the number of filled positions
	filled := int(percentage / 100 * float64(width))
	if filled > width {
		filled = width
	}
	
	// Build the progress bar
	bar := "["
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "="
		} else {
			bar += " "
		}
	}
	bar += "]"
	
	return bar
}