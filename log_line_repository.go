package main

import (
	"time"

	"github.com/coopernurse/gorp"
)

// LogLineRepository has methods for adding and removing log lines.
type LogLineRepository struct {
	job   *Job
	dbmap *gorp.DbMap
}

// Add adds a LogLine.
func (r *LogLineRepository) Add(output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{JobID: r.job.ID, Output: output, Timestamp: timestamp}
	err := r.Insert(l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Insert inserts a LogLine.
func (r *LogLineRepository) Insert(logLine *LogLine) error {
	return r.dbmap.Insert(logLine)
}

// AllForJob returns a slice of LogLine for the Job.
func (r *LogLineRepository) All() ([]LogLine, error) {
	var lines []LogLine

	_, err := r.dbmap.Select(&lines, `SELECT * FROM log_lines WHERE job_id = $1 ORDER BY timestamp`, r.job.ID)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
