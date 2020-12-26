// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"errors"
	"time"
)

// Task represents a temporal task.
type Task struct {
	name        string
	start       time.Time
	end         time.Time
	isFixed     bool
	priority    int
	Meta        interface{}
	isScheduled bool
	previous    *Task
	next        *Task
}

// DefaultPriority is the default priority number for tasks.
const DefaultPriority int = 0

// NewTask creates a new "moving" task, that is, a task that might be moved
// during scheduling.
func NewTask(name string, duration time.Duration, priority int, meta interface{}) (*Task, error) {
	if duration < 0 {
		return nil, errors.New("duration must be >= 0")
	}

	return &Task{
		name:     name,
		start:    time.Time{},
		end:      time.Time{}.Add(duration),
		priority: priority,
		Meta:     meta,
	}, nil
}

// NewFixedTask creates a new fixed task, that is, a task that cannot be moved
// by other tasks during scheduling.
func NewFixedTask(name string, start, end time.Time, priority int, meta interface{}) (*Task, error) {
	if start.After(end) {
		return nil, errors.New("start must happen before end")
	}

	return &Task{
		name:     name,
		start:    start,
		end:      end,
		isFixed:  true,
		priority: priority,
		Meta:     meta,
	}, nil
}

// Name returns the name of the task t.
func (t *Task) Name() string {
	return t.name
}

// Start returns the start time of the task t. Might be zero if not scheduled yet.
func (t *Task) Start() time.Time {
	return t.start
}

// End returns the end time of the task t. Might be zero + duration if not scheduled yet (see Duration).
func (t *Task) End() time.Time {
	return t.end
}

// IsFixed returns whether the task t is fixed or not.
func (t *Task) IsFixed() bool {
	return t.isFixed
}

// Priority returns the priority of the task t.
func (t *Task) Priority() int {
	return t.priority
}

// IsScheduled returns whether the task t is already scheduled.
func (t *Task) IsScheduled() bool {
	return t.isScheduled
}

// Duration returns the duration of the task (in nanoseconds).
func (t *Task) Duration() time.Duration {
	return t.end.Sub(t.start)
}

// IntersectsWithTimeInterval checks whether t's period intersects with the given time interval.
func (t *Task) IntersectsWithTimeInterval(start, end time.Time) bool {
	return !(t.start.After(end) || t.end.Before(start))
}

// IntersectsWithTask checks whether t's period intersects with task's period.
func (t *Task) IntersectsWithTask(task *Task) bool {
	return t.IntersectsWithTimeInterval(task.start, task.end)
}

// ContainsTime checks whether time happens within t's period.
func (t *Task) ContainsTime(time time.Time) bool {
	return !(t.start.After(time) || t.end.Before(time))
}

// FreeTimeBetweenTasks returns the free time interval between tasks a and b. To be tested.
func FreeTimeBetweenTasks(a, b *Task) (start, end time.Time, err error) {
	if a.next != b {
		return time.Time{}, time.Time{}, errors.New("a and b must be consecutives tasks")
	} else if b.start.Sub(a.end) <= time.Minute {
		return a.end, a.end, nil
	}
	return a.end.Add(time.Minute), b.start.Add(-time.Minute), nil
}
