// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"errors"
	"time"
)

// Task represents a temporal task.
type Task struct {
	start    time.Time
	end      time.Time
	isFixed  bool
	previous *Task
	next     *Task
}

// NewTask creates a new "moving" task, that is, a task that might be moved
// during scheduling.
func NewTask(duration time.Duration) (*Task, error) {
	if duration < 0 {
		return nil, errors.New("duration must be >= 0")
	}

	return &Task{
		start:   time.Time{},
		end:     time.Time{}.Add(duration),
		isFixed: false, // Probably redundant
	}, nil
}

// NewFixedTask creates a new fixed task, that is, a task that cannot be moved
// by other tasks during scheduling.
func NewFixedTask(start, end time.Time) (*Task, error) {
	if start.After(end) {
		return nil, errors.New("start must happen before end")
	}

	return &Task{start: start, end: end, isFixed: true}, nil
}

// Start returns the start time of the task t. Might be zero if not scheduled yet.
func (t *Task) Start() time.Time {
	return t.start
}

// End returns the end time of the task t. Might be zero + duration if not scheduled yet (see Duration).
func (t *Task) End() time.Time {
	return t.end
}

// IsFixed returns whether the task is fixed or not.
func (t *Task) IsFixed() bool {
	return t.isFixed
}

// Duration returns the duration of the task (in nanoseconds).
func (t *Task) Duration() time.Duration {
	return t.end.Sub(t.start)
}

// IntersectsWithTask checks whether t intersects with task.
func (t *Task) IntersectsWithTask(task *Task) bool {
	return !(t.start.After(task.end) || t.end.Before(task.start))
}

// ContainsTime checks whether time happens within t's period.
func (t *Task) ContainsTime(time time.Time) bool {
	return !(t.start.After(time) || t.end.Before(time))
}
