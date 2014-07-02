package shipr

import (
	"database/sql"
	"log"

	"bitbucket.org/liamstask/goose/lib/goose"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

// Shipr maintains a shipr context.
var (
	// The shipr environment
	Env string

	// Database
	db    *sql.DB
	dbmap *gorp.DbMap

	// Repositories
	Repos *RepoRepository
	Jobs  *JobRepository

	// Deployer
	deployer Deployer
)

func init() {
	if Env == "" {
		Env = "development"
	}

	initDB()
}

func initDB() {
	var err error

	dbconf, err := goose.NewDBConf("db", Env)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open(dbconf.Driver.Name, dbconf.Driver.OpenStr)
	if err != nil {
		log.Fatal(err)
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Repo{}, "repos").SetKeys(true, "ID")
	dbmap.AddTableWithName(Job{}, "jobs").SetKeys(true, "ID")
	dbmap.AddTableWithName(LogLine{}, "log_lines").SetKeys(true, "ID")

	Repos = &RepoRepository{dbmap}
	Jobs = &JobRepository{dbmap}
}

// Deploy takes a Deployable, creates a Job for it and runs the deployment.
func Deploy(d Deployable) error {
	j, err := Jobs.CreateByDeployable(d)
	if err != nil {
		return err
	}
	return j.Run()
}
