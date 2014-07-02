package shipr

// Deployers are capable of deploying jobs.
type Deployer interface {
	Deploy(*Job) error
}
