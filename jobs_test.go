package shipr

import "testing"

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
