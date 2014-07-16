package shipr

import "time"

// LogLinesService manages the `log_lines` table.
type LogLinesService struct {
	*Datastore
}

func (c *LogLinesService) CreateLine(job *Job, output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{
		JobID:     job.ID,
		Job:       job,
		Output:    output,
		Timestamp: timestamp,
	}

	return l, c.Insert(l)
}

type LogLine struct {
	ID        int
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time

	// Memoized Job.
	Job *Job `db:"-"`
}

func (l *LogLine) table() string {
	return "log_lines"
}
