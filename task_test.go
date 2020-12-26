// Copyright (c) 2020 Davi Villalva.
// license can be found in the LICENSE file.

package dvtask

import (
	"testing"
	"time"
)

var (
	now             = time.Now()
	oneMinuteAgo    = now.Add(-time.Minute)
	oneMinuteLater  = now.Add(time.Minute)
	twoMinutesAgo   = now.Add(-2 * time.Minute)
	twoMinutesLater = now.Add(2 * time.Minute)
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
		{"valid", args{time.Minute}, false},
		{"invalid", args{-time.Minute}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTask("", tt.args.duration, 0, nil)
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
			_, err := NewFixedTask("", tt.args.start, tt.args.end, 0, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFixedTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTask_IntersectsWithTimeInterval(t *testing.T) {
	type fields struct {
		start time.Time
		end   time.Time
	}
	type args struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"intersection", fields{twoMinutesAgo, now}, args{oneMinuteAgo, oneMinuteLater}, true},
		{"nointersection", fields{twoMinutesAgo, oneMinuteAgo}, args{now, oneMinuteLater}, false},
		{"edgeintersection0", fields{oneMinuteAgo, now}, args{now, oneMinuteLater}, true},
		{"edgeintersection0inv", fields{now, oneMinuteLater}, args{oneMinuteAgo, now}, true},
		{"edgeintersection1", fields{now, oneMinuteLater}, args{oneMinuteLater, twoMinutesLater}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := task.IntersectsWithTimeInterval(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Task.IntersectsWithTimeInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_IntersectsWithTask(t *testing.T) {
	type fields struct {
		start time.Time
		end   time.Time
	}
	tests := []struct {
		name    string
		aFields fields
		bFields fields
		want    bool
	}{
		{"intersection", fields{twoMinutesAgo, now}, fields{oneMinuteAgo, oneMinuteLater}, true},
		{"nointersection", fields{twoMinutesAgo, oneMinuteAgo}, fields{now, oneMinuteLater}, false},
		{"edgeintersection0", fields{oneMinuteAgo, now}, fields{now, oneMinuteLater}, true},
		{"edgeintersection0inv", fields{now, oneMinuteLater}, fields{oneMinuteAgo, now}, true},
		{"edgeintersection1", fields{now, oneMinuteLater}, fields{oneMinuteLater, twoMinutesLater}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskA := &Task{
				start: tt.aFields.start,
				end:   tt.aFields.end,
			}
			taskB := &Task{
				start: tt.bFields.start,
				end:   tt.bFields.end,
			}
			if got := taskA.IntersectsWithTask(taskB); got != tt.want {
				t.Errorf("Task.IntersectsWithTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_ContainsTime(t *testing.T) {
	type fields struct {
		start time.Time
		end   time.Time
	}
	type args struct {
		time time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"contains", fields{now, twoMinutesLater}, args{oneMinuteLater}, true},
		{"edgycontains", fields{now, twoMinutesLater}, args{twoMinutesLater}, true},
		{"doesnotcontain", fields{now, oneMinuteLater}, args{twoMinutesLater}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				start: tt.fields.start,
				end:   tt.fields.end,
			}
			if got := task.ContainsTime(tt.args.time); got != tt.want {
				t.Errorf("Task.ContainsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
