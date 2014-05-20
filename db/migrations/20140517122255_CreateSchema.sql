
-- +goose Up
CREATE TABLE repos (
  id SERIAL,
  name character varying(255),
  PRIMARY KEY(id)
);

CREATE INDEX index_repos_on_name ON repos USING btree (name);

CREATE TABLE jobs (
  id SERIAL,
  repo_id int NOT NULL references repos(id),
  guid character varying(255) NOT NULL,
  sha character varying(255) NOT NULL,
  environment character varying(255) NOT NULL,
  force boolean NOT NULL DEFAULT false,
  description text,
  exit_status integer,
  PRIMARY KEY(id)
);

CREATE INDEX index_jobs_on_guid ON jobs USING btree (guid);
CREATE INDEX index_jobs_on_sha ON jobs USING btree (sha);

CREATE TABLE log_lines (
  id SERIAL,
  job_id int NOT NULL references jobs(id),
  output text NOT NULL,
  timestamp timestamp NOT NULL,
  PRIMARY KEY(id)
);

CREATE INDEX index_log_lines_on_job_id ON log_lines USING btree (job_id);
CREATE INDEX index_log_lines_on_timestamp ON log_lines USING btree (timestamp);

-- +goose Down
DROP TABLE repos;
DROP TABLE jobs;
DROP TABLE log_lines;
