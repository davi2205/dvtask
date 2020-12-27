// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"errors"
	"time"
)

// Scheduler schedules tasks
type Scheduler struct {
	startDate time.Time
	deadline  time.Time
	workStart Hour
	workEnd   Hour
	firstTask *Task
}

// NewScheduler creates a new Scheduler (may remove this in the future). To be tested.
func NewScheduler(startDate, deadline time.Time, workStart, workEnd Hour) *Scheduler {
	return &Scheduler{
		startDate: startDate,
		deadline:  deadline,
		workStart: workStart,
		workEnd:   workEnd,
	}
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

// Not ready yet. To be tested.
func (s *Scheduler) ScheduleTask(task *Task, fromTime time.Time) error {
	if fromTime.Before(time.Now().Add(5 * time.Minute)) {
		return errors.New("tasks must be scheduled at least five minutes from now")
	}

	taskDuration := task.Duration()

	if s.firstTask == nil {
		task.start = fromTime
		task.end = fromTime.Add(taskDuration)

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
