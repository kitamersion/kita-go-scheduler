package main

import (
	"io"
	"kita-go-scheduler/config"
	"kita-go-scheduler/constants"
	"kita-go-scheduler/scheduler"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set up logging
	var logFile *os.File
	if cfg.Logs.Enabled {
		// Get the user's home directory to save log in ~/.config/{constants.PROJECT_NAME}/
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Failed to get home directory: %v", err)
		}

		// Set up log file path
		logDir := filepath.Join(homeDir, ".config", constants.PROJECT_NAME)
		if _, err := os.Stat(logDir); os.IsNotExist(err) {
			err := os.MkdirAll(logDir, 0755)
			if err != nil {
				log.Fatalf("Failed to create log directory: %v", err)
			}
		}

		// Create or open the log file
		logFile, err = os.OpenFile(filepath.Join(logDir, constants.PROJECT_NAME+".log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		defer logFile.Close()

		// Set up logging to both console and file
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
	} else {
		// If logging is disabled, only log to stdout
		log.SetOutput(os.Stdout)
	}

	log.Println("Starting", constants.PROJECT_NAME)

	// Initialize scheduler
	s := scheduler.NewScheduler()

	// Add tasks from configuration
	for _, taskConfig := range cfg.Tasks {
		task := scheduler.NewTask(taskConfig.Name, taskConfig.Schedule, taskConfig.Command)
		if err := s.AddTask(task); err != nil {
			log.Fatalf("Failed to add task %s: %v", task.Name, err)
		}
	}

	// Start scheduler
	s.Start()
	defer s.Stop()

	// Keep the application running
	select {}
}
