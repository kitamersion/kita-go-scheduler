package scheduler

import "fmt"

type Task struct {
	Name     string // Name of the task
	Schedule string // Cron-style schedule string
	Command  string // Command to execute
}

// NewTask creates a new Task instance.
func NewTask(name, schedule, command string) Task {
	return Task{
		Name:     name,
		Schedule: schedule,
		Command:  command,
	}
}

// Validate checks if the Task's fields are valid.
func (t *Task) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("task name cannot be empty")
	}
	if t.Schedule == "" {
		return fmt.Errorf("task schedule cannot be empty")
	}
	if t.Command == "" {
		return fmt.Errorf("task command cannot be empty")
	}
	return nil
}
