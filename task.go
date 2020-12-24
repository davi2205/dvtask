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
	fixed    bool
	previous *Task
	next     *Task
}

// Creates a new fixed task, that is, a task that cannot be moved
// by other tasks during scheduling.
func NewFixedTask(start, end time.Time) (*Task, error) {
	if start.After(end) {
		return nil, errors.New("start must happen before the end")
	}

	return &Task{start: start, end: end, fixed: true}
}
