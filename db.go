package shipr

import (
	"database/sql"

	"github.com/coopernurse/gorp"

	"bitbucket.org/liamstask/goose/lib/goose"
)

type DB struct {
	// The database configuration.
	DBConf *goose.DBConf

	// The database connection.
	DB *sql.DB

	// The gorp dbmap. Also mixin methods.
	*gorp.DbMap

	// Services
	Repos *ReposService
	Jobs  *JobsService
}

func NewDB(path, env string) (*DB, error) {
	dbconf, err := goose.NewDBConf(path, env)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbconf.Driver.Name, dbconf.Driver.OpenStr)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Repo{}, "repos").SetKeys(true, "ID")

	d := &DB{DBConf: dbconf, DB: db, DbMap: dbmap}
	d.Repos = &ReposService{d}
	d.Jobs = &JobsService{d}
	return d, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
