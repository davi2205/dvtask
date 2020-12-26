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

// ScheduledTasksInTimeInterval returns the first and last scheduled tasks within the given time interval.
// Returns nil, nil if no tasks are found within the time interval. To be tested.
func (s *Scheduler) ScheduledTasksInTimeInterval(start, end time.Time) (first, last *Task) {
	for currentTask := s.firstTask; currentTask != nil; currentTask = currentTask.next {
		if currentTask.IntersectsWithTimeInterval(start, end) {
			if first != nil {
				return
			}
			continue
		}
		if first == nil {
			first = currentTask
		}
		last = currentTask
	}
	return
}

// Not ready yet
func (s *Scheduler) scheduleNonFixedTask(task *Task) error {
	if s.firstTask == nil {
		s.firstTask = task
		return nil
	}

	for currentTask := s.firstTask; ; currentTask = currentTask.next {
		if currentTask.priority < task.priority {

			break
		} else if currentTask.next == nil {

			break
		}
	}

	task.isScheduled = true

	return errors.New("not implemented yet")
}

// Not ready yet
func (s *Scheduler) Schedule(task *Task) error {
	if task.isScheduled {
		return errors.New("task already scheduled")
	}

	if task.isFixed {
		return errors.New("not implemented yet")
	}
	return s.scheduleNonFixedTask(task)
}
