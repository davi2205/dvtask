// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask_test

import (
	"testing"
	"time"

	"github.com/davi2205/dvtask"
)

const (
	oneMinute  time.Duration = 60_000_000_000
	twoMinutes time.Duration = 120_000_000_000
)

var (
	now            = time.Now()
	oneMinuteAgo   = now.Add(-oneMinute)
	oneMinuteLater = now.Add(oneMinute)
)

func TestNewTask(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{oneMinute}, false},
		{"invalid", args{-oneMinute}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dvtask.NewTask(tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewFixedTask(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{now, oneMinuteLater}, false},
		{"invalid", args{now, oneMinuteAgo}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := dvtask.NewFixedTask(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFixedTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
