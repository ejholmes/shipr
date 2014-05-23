package main

type HerokuDeployer struct {
	ApiKey string
}

func (h *HerokuDeployer) Deploy(j *Job) error {
	return nil
}
