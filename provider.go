package shipr

// Provider is an interface that can be implemented for deploying a Deployment to
// some platform.
type Provider interface {
	Deploy(Deployment) error
}
