package tests

import (
	"kita-go-scheduler/scheduler"
	"testing"
	"time"
)

func TestScheduler(t *testing.T) {
	// Initialize scheduler
	s := scheduler.NewScheduler()

	// Create a task
	task := scheduler.NewTask("Test Task", "@every 1s", "echo 'Task executed'")

	// Add the task to the scheduler
	err := s.AddTask(task)
	if err != nil {
		t.Fatalf("Failed to add task: %v", err)
	}

	// Start the scheduler
	go s.Start()
	defer s.Stop()

	// Allow the task to run
	time.Sleep(2 * time.Second) // Wait for at least one execution

	// Check if the task executed (log output should confirm execution)
	// Note: This is a basic test and assumes manual verification of logs
	// For more robust testing, consider mocking the exec.Command call.
	t.Log("Check logs to verify task execution.")
}
