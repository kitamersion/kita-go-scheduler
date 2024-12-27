package scheduler

import (
	"log"
	"os/exec"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron *cron.Cron
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		cron: cron.New(),
	}
}

func (s *Scheduler) AddTask(task Task) error {
	if err := task.Validate(); err != nil {
		return err
	}

	_, err := s.cron.AddFunc(task.Schedule, func() {
		log.Printf("Executing task: %s", task.Name)
		if output, err := exec.Command("sh", "-c", task.Command).CombinedOutput(); err != nil {
			log.Printf("Task %s failed: %v\nOutput: %s", task.Name, err, string(output))
		} else {
			log.Printf("Task %s succeeded. Output: %s", task.Name, string(output))
		}
	})
	return err
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
