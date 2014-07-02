package main

import "time"

// LogLine represents a line of log output from the deploy job.
type LogLine struct {
	ID        int
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time
}
