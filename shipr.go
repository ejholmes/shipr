package shipr

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string
}

// Returns a new Shipr context.
func New() *Shipr {
	return &Shipr{Env: "development"}
}
