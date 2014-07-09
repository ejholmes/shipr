package shipr

type Deployer interface {
	Deploy(Deployable) error
}

type HerokuDeployer struct {
}

func (h *HerokuDeployer) Deploy(d Deployable) error {
	return nil
}
