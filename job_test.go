package main

import (
	"testing"
	"time"
)

func testJob(t *testing.T) *Job {
	job := &Job{Guid: "1234", Sha: "4321", Environment: "production"}
	err := dbmap.Insert(job)
	if err != nil {
		t.Error(err)
	}
	return job
}

func Test_Job_AddLine(t *testing.T) {
	job := testJob(t)

	l, err := job.AddLine("Foo\n", time.Now())
	if err != nil {
		t.Error(err)
	}

	output, err := job.Output()
	if err != nil {
		t.Error(err)
	}

	if output != l.Output {
		t.Errorf("Got %v; want %v", output, l.Output)
	}
}

func Test_Job_Status(t *testing.T) {
	exitOk := 0
	exitFailed := 1

	tests := []struct {
		Status   *int
		Expected JobStatus
	}{
		{nil, StatusPending},
		{&exitOk, StatusSucceeded},
		{&exitFailed, StatusFailed},
	}

	for _, test := range tests {
		job := &Job{ExitStatus: test.Status}

		if job.Status() != test.Expected {
			t.Fatalf("Expected job status to be %v", test.Expected)
		}
	}
}

func Test_Job_IsDone(t *testing.T) {
	status := 1

	tests := []struct {
		Status   *int
		Expected bool
	}{
		{nil, false},
		{&status, true},
	}

	for _, test := range tests {
		job := &Job{ExitStatus: test.Status}

		if job.IsDone() != test.Expected {
			t.Fatalf("Expected IsDone() to return %v", test.Expected)
		}
	}
}
