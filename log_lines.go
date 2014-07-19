package shipr

import "time"

// LogLinesService manages the `log_lines` table.
type LogLinesService struct {
	*Datastore
}

// CreateLine creates a new log line for the given job.
func (c *LogLinesService) CreateLine(job *Job, output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{
		JobID:     job.ID,
		Job:       job,
		Output:    output,
		Timestamp: timestamp,
	}

	return l, c.Insert(l)
}

// LogLine maps fields from the `log_lines` table.
type LogLine struct {
	ID        int
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time

	// Memoized Job.
	Job *Job `db:"-"`
}
