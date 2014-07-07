package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"bitbucket.org/liamstask/goose/lib/goose"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

// Deployment is an interface that's describes information about a deployment.
// This can be used to create jobs, and is also implemented by Job.
type Deployment interface {
	// Guid should return a unique identifier for this deployment.
	Guid() int

	// RepoName should return the string name of the repo to deploy.
	RepoName() string

	// Sha should return the git sha that we want to deploy.
	Sha() string

	// Ref should return the git ref that is being requested.
	Ref() string

	// Environment should return the name of the environment that the repo is being deploy to.
	Environment() string

	// Description should return a string description for the deploy.
	Description() string
}

// Deployable is an interface the Job implements and provides methods that
// deployers can use to get information about the Job, and update it.
type Deployable interface {
	Deployment

	// AddLine adds a log line to the output.
	AddLine(string, time.Time) error
}

// Deployers are capable of deploying jobs.
type Deployer interface {
	Deploy(Deployable) error
}

var (
	Env string

	// Database
	db    *sql.DB
	dbmap *gorp.DbMap

	// Repositories
	Repos *RepoRepository
	Jobs  *JobRepository

	// Deployers
	deployer Deployer
)

func init() {
	if Env == "" {
		Env = "development"
	}

	initDb()

	// Setup deployers.
	deployer = NewHerokuDeployer("", "")
}

func initDb() {
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

func main() {
	defer db.Close()

	server := NewServer()
	http.ListenAndServe(":3001", server)
}

// Deploy takes a Jobable, creates a Job for it and runs the deployment.
func Deploy(d Deployment) error {
	j, err := Jobs.CreateFromDeployment(d)
	if err != nil {
		return err
	}
	return j.Run()
}
