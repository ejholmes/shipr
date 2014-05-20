package main

type Deployer interface {
	// Starts the deployment of the repo.
	Deploy() error
}
