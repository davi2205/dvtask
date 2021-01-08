// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import "time"

// Hour represents an exact hour
type Hour struct {
	Hour, Min, Sec int8
}

// Time converts the Hour h to the type time.Time. To be tested?
func (h Hour) Time() time.Time {
	return time.Date(0, 0, 0, int(h.Hour), int(h.Min), int(h.Sec), 0, time.UTC)
}
