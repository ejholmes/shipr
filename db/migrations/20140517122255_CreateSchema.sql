
-- +goose Up
CREATE TABLE jobs (
  id int NOT NULL,
  guid text NOT NULL,
  sha text NOT NULL,
  environment text NOT NULL,
  force boolean NOT NULL DEFAULT false,
  description text,
  exit_status integer,
  PRIMARY KEY(id)
);

CREATE SEQUENCE jobs_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER SEQUENCE jobs_id_seq OWNED BY jobs.id;
ALTER TABLE ONLY jobs ALTER COLUMN id SET DEFAULT nextval('jobs_id_seq'::regclass);

CREATE INDEX index_jobs_on_guid ON jobs USING btree (guid);
CREATE INDEX index_jobs_on_sha ON jobs USING btree (sha);

CREATE TABLE log_lines (
  id int NOT NULL,
  job_id int NOT NULL,
  output text NOT NULL,
  timestamp timestamp NOT NULL,
  PRIMARY KEY(id)
);

CREATE SEQUENCE log_lines_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER SEQUENCE log_lines_id_seq OWNED BY log_lines.id;
ALTER TABLE ONLY log_lines ALTER COLUMN id SET DEFAULT nextval('log_lines_id_seq'::regclass);

CREATE INDEX index_log_lines_on_job_id ON log_lines USING btree (job_id);
CREATE INDEX index_log_lines_on_timestamp ON log_lines USING btree (timestamp);

-- +goose Down
DROP TABLE jobs;
DROP TABLE log_lines;
