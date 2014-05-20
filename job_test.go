package main

import (
	"testing"
	"time"
)

func testRepo(t *testing.T) *Repo {
	repo := &Repo{Name: "remind101/test"}
	err := dbmap.Insert(repo)
	if err != nil {
		t.Error(err)
	}
	return repo
}

func testJob(t *testing.T) *Job {
	repo := testRepo(t)
	job := &Job{RepoID: repo.ID, Guid: "1234", Sha: "4321", Environment: "production"}
	err := dbmap.Insert(job)
	if err != nil {
		t.Error(err)
	}
	return job
}

func Test_Job_AddLine(t *testing.T) {
	type line struct {
		Output    string
		Timestamp time.Time
	}

	tests := []struct {
		Lines    []line
		Expected string
	}{
		{
			[]line{
				line{"Foo\n", time.Now()},
			},
			"Foo\n",
		},
		{
			[]line{
				line{"First\n", time.Now()},
				line{"Second\n", time.Now()},
			},
			"First\nSecond\n",
		},
		{
			[]line{
				line{"Second\n", time.Date(2009, time.November, 10, 24, 0, 0, 0, time.UTC)},
				line{"First\n", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)},
			},
			"First\nSecond\n",
		},
	}

	for i, test := range tests {
		job := testJob(t)

		for _, l := range test.Lines {
			job.AddLine(l.Output, l.Timestamp)
		}

		output, err := job.Output()
		if err != nil {
			t.Error(err)
		}

		if output != test.Expected {
			t.Errorf("%d: Got %v; want %v", i, output, test.Expected)
		}
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
