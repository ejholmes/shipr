package shipr

import "time"

// LogLinesService manages the `log_lines` table.
type LogLinesService struct {
	*Datastore
}

// CreateLine creates a new log line for the given job.
func (s *LogLinesService) CreateLine(job *Job, output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{
		JobID:     job.ID,
		Job:       job,
		Output:    output,
		Timestamp: timestamp,
	}

	return l, s.Insert(l)
}

// Output returns the string output of log lines for a job.
func (s *LogLinesService) Output(job *Job) (string, error) {
	var lines []*LogLine

	sql := `SELECT * FROM log_lines WHERE job_id = $1 ORDER BY timestamp ASC`
	err := s.Select(&lines, sql, job.ID)
	if err != nil {
		return "", err
	}

	out := ""
	for _, l := range lines {
		out = out + l.Output
	}

	return out, nil
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
