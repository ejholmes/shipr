package shipr

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// The DB connection.
	*DB
}

// Returns a new Shipr context.
func New(path, env string) (*Shipr, error) {
	db, err := NewDB(path, env)
	if err != nil {
		return nil, err
	}
	return &Shipr{Env: env, DB: db}, nil
}

func (c *Shipr) Close() {
	c.DB.Close()
}
