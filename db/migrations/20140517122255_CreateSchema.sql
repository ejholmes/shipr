
-- +goose Up
CREATE TABLE jobs (
  id int NOT NULL,
  guid text NOT NULL,
  sha text NOT NULL,
  environment text NOT NULL,
  PRIMARY KEY(id)
);

CREATE INDEX index_jobs_on_guid ON jobs USING btree (guid);
CREATE INDEX index_jobs_on_sha ON jobs USING btree (sha);

CREATE TABLE log_lines (
  id int NOT NULL,
  job_id int NOT NULL,
  output text NOT NULL,
  timestamp timestamp NOT NULL,
  PRIMARY KEY(id)
);

CREATE INDEX index_log_lines_on_timestamp ON log_lines USING btree (timestamp);

-- +goose Down
DROP TABLE jobs;
DROP TABLE log_lines;
