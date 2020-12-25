// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"time"
)

// Scheduler schedules tasks
type Scheduler struct {
	firstTask *Task
}

// NewScheduler creates a new Scheduler (may remove this in the future).
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// ScheduledTaskAt gets the schedule task at a specific time. To be tested.
func (s *Scheduler) ScheduledTaskAt(time time.Time) *Task {
	for currentTask := s.firstTask; currentTask != nil; currentTask = currentTask.next {
		if !(time.Before(currentTask.start) || time.After(currentTask.end)) {
			return currentTask
		}
	}
	return nil
}
