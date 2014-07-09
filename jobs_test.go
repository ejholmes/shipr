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
