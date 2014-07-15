package shipr

import (
	"database/sql"

	"github.com/coopernurse/gorp"

	"bitbucket.org/liamstask/goose/lib/goose"
)

// DB is an interface that allows us to CRUD records.
type DB interface {
	Insert(v interface{}) error
	Update(v interface{}) error
	//Get(holder interface{}, field string, value interface{}) error
	SelectOne(holder interface{}, query string, args ...interface{}) error
	Close() error
}

// db is an implementation of the DB interface backed by gorp and postgres.
type db struct {
	// The database configuration.
	DBConf *goose.DBConf

	// The database connection.
	DB *sql.DB

	// The gorp dbmap. Also mixin methods.
	Map *gorp.DbMap
}

func NewDB(path, env string) (DB, error) {
	dbconf, err := goose.NewDBConf(path, env)
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open(dbconf.Driver.Name, dbconf.Driver.OpenStr)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: conn, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Repo{}, "repos").SetKeys(true, "ID")
	dbmap.AddTableWithName(Job{}, "jobs").SetKeys(true, "ID")
	dbmap.AddTableWithName(LogLine{}, "log_lines").SetKeys(true, "ID")

	return &db{DBConf: dbconf, DB: conn, Map: dbmap}, nil
}

func (d *db) Insert(v interface{}) error {
	return d.Map.Insert(v)
}

func (d *db) Update(v interface{}) error {
	_, err := d.Map.Update(v)
	return err
}

func (d *db) SelectOne(holder interface{}, query string, args ...interface{}) error {
	return d.Map.SelectOne(holder, query, args...)
}

func (d *db) Close() error {
	return d.DB.Close()
}

// Datastore holds all of our services and DB reference.
type Datastore struct {
	DB
	Repos    *ReposService
	Jobs     *JobsService
	LogLines *LogLinesService
}

// NewDatastore returns a new Datastore with all services configured.
func NewDatastore(db DB) *Datastore {
	s := &Datastore{DB: db}
	s.Repos = &ReposService{s}
	s.Jobs = &JobsService{s}
	s.LogLines = &LogLinesService{s}
	return s
}
