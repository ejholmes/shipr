package shipr

type Deployer interface {
	Deploy(Deployable) error
}
