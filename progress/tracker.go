package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CompletionStatus represents the status of a tutorial or exercise
type CompletionStatus struct {
	Completed  bool      `json:"completed"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
	Score      int       `json:"score,omitempty"` // For exercises with scoring
}

// ProgressData stores all user progress
type ProgressData struct {
	Tutorials map[string]CompletionStatus `json:"tutorials"` // Maps tutorial name to status
	Exercises map[string]CompletionStatus `json:"exercises"` // Maps exercise name to status
}

// Tracker manages progress tracking
type Tracker struct {
	data      ProgressData
	configDir string
	filePath  string
}

// New creates a new progress tracker
func New() (*Tracker, error) {
	// Get user config directory
	configDir, err := getUserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get config directory: %w", err)
	}

	// Initialize tracker
	tracker := &Tracker{
		data: ProgressData{
			Tutorials: make(map[string]CompletionStatus),
			Exercises: make(map[string]CompletionStatus),
		},
		configDir: configDir,
		filePath:  filepath.Join(configDir, "progress.json"),
	}

	// Load existing data if available
	if err := tracker.load(); err != nil {
		// If file doesn't exist, we'll create it when saving
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to load progress data: %w", err)
		}
	}

	return tracker, nil
}

// MarkTutorialComplete marks a tutorial as completed
func (t *Tracker) MarkTutorialComplete(name string) error {
	t.data.Tutorials[name] = CompletionStatus{
		Completed:  true,
		CompletedAt: time.Now(),
	}
	return t.save()
}

// MarkExerciseComplete marks an exercise as completed with an optional score
func (t *Tracker) MarkExerciseComplete(name string, score int) error {
	t.data.Exercises[name] = CompletionStatus{
		Completed:  true,
		CompletedAt: time.Now(),
		Score:      score,
	}
	return t.save()
}

// IsTutorialCompleted checks if a tutorial has been completed
func (t *Tracker) IsTutorialCompleted(name string) bool {
	status, exists := t.data.Tutorials[name]
	return exists && status.Completed
}

// IsExerciseCompleted checks if an exercise has been completed
func (t *Tracker) IsExerciseCompleted(name string) bool {
	status, exists := t.data.Exercises[name]
	return exists && status.Completed
}

// GetCompletedTutorials returns a list of completed tutorials
func (t *Tracker) GetCompletedTutorials() []string {
	var completed []string
	for name, status := range t.data.Tutorials {
		if status.Completed {
			completed = append(completed, name)
		}
	}
	return completed
}

// GetCompletedExercises returns a list of completed exercises
func (t *Tracker) GetCompletedExercises() []string {
	var completed []string
	for name, status := range t.data.Exercises {
		if status.Completed {
			completed = append(completed, name)
		}
	}
	return completed
}

// GetProgress returns a summary of overall progress
func (t *Tracker) GetProgress(totalTutorials, totalExercises int) (int, int, float64) {
	completedTutorials := len(t.GetCompletedTutorials())
	completedExercises := len(t.GetCompletedExercises())
	
	total := totalTutorials + totalExercises
	completed := completedTutorials + completedExercises
	
	var percentage float64 = 0
	if total > 0 {
		percentage = float64(completed) / float64(total) * 100
	}
	
	return completed, total, percentage
}

// Reset clears all progress
func (t *Tracker) Reset() error {
	t.data = ProgressData{
		Tutorials: make(map[string]CompletionStatus),
		Exercises: make(map[string]CompletionStatus),
	}
	return t.save()
}

// save stores the progress data to disk
func (t *Tracker) save() error {
	// Ensure directory exists
	if err := os.MkdirAll(t.configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	
	// Marshal data to JSON
	data, err := json.MarshalIndent(t.data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal progress data: %w", err)
	}
	
	// Write to file
	if err := os.WriteFile(t.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write progress file: %w", err)
	}
	
	return nil
}

// load reads the progress data from disk
func (t *Tracker) load() error {
	// Read file
	data, err := os.ReadFile(t.filePath)
	if err != nil {
		return err
	}
	
	// Unmarshal JSON
	if err := json.Unmarshal(data, &t.data); err != nil {
		return fmt.Errorf("failed to parse progress data: %w", err)
	}
	
	return nil
}

// getUserConfigDir returns the directory for storing config
func getUserConfigDir() (string, error) {
	// Get user config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		// Fallback to home directory
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configDir = filepath.Join(homeDir, ".config")
	}
	
	// Create specific app directory
	appConfigDir := filepath.Join(configDir, "gocli-teacher")
	return appConfigDir, nil
}