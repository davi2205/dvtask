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
	now             = time.Now()
	oneMinuteAgo    = now.Add(-oneMinute)
	oneMinuteLater  = now.Add(oneMinute)
	twoMinutesAgo   = now.Add(-twoMinutes)
	twoMinutesLater = now.Add(twoMinutes)
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

func TestTask_IntersectsWithInterval(t *testing.T) {
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
		{"edgeintersection", fields{oneMinuteAgo, now}, args{now, oneMinuteLater}, true},
		{"edgeintersection_inv", fields{now, oneMinuteLater}, args{oneMinuteAgo, now}, true},
		{"edgeintersection_1", fields{now, oneMinuteLater}, args{oneMinuteLater, twoMinutesLater}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task, _ := dvtask.NewFixedTask(tt.fields.start, tt.fields.end)
			if got := task.IntersectsWithInterval(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Task.IntersectsWithInterval() = %v, want %v", got, tt.want)
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
		{"edgeintersection", fields{oneMinuteAgo, now}, fields{now, oneMinuteLater}, true},
		{"edgeintersection_inv", fields{now, oneMinuteLater}, fields{oneMinuteAgo, now}, true},
		{"edgeintersection_1", fields{now, oneMinuteLater}, fields{oneMinuteLater, twoMinutesLater}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			taskA, _ := dvtask.NewFixedTask(tt.aFields.start, tt.aFields.end)
			taskB, _ := dvtask.NewFixedTask(tt.bFields.start, tt.bFields.end)
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
			task, _ := dvtask.NewFixedTask(tt.fields.start, tt.fields.end)
			if got := task.ContainsTime(tt.args.time); got != tt.want {
				t.Errorf("Task.ContainsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
