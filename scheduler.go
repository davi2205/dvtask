// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"errors"
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

// ScheduledTaskAt returns the scheduled task at a specific time. To be tested.
func (s *Scheduler) ScheduledTaskAt(time time.Time) *Task {
	for currentTask := s.firstTask; currentTask != nil; currentTask = currentTask.next {
		if currentTask.ContainsTime(time) {
			return currentTask
		}
	}
	return nil
}

// ScheduledTasksInInterval returns the first and last scheduled tasks within the given interval.
// Returns nil, nil if no tasks are found within the interval. To be tested.
func (s *Scheduler) ScheduledTasksInInterval(start, end time.Time) (first, last *Task) {
	for currentTask := s.firstTask; currentTask != nil; currentTask = currentTask.next {
		if currentTask.IntersectsWithInterval(start, end) {
			if first != nil {
				return
			} else {
				continue
			}
		}

		if first == nil {
			first = currentTask
		}

		last = currentTask
	}
	return
}

func (s *Scheduler) scheduleNonFixedTask(task *Task) error {
	return errors.New("not implemented yet")
}

func (s *Scheduler) Schedule(task *Task) error {
	if task.isFixed {
		return errors.New("not implemented yet")
	} else {
		return s.scheduleNonFixedTask(task)
	}
}
