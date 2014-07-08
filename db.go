package shipr

import "bitbucket.org/liamstask/goose/lib/goose"

type DB struct {
	*goose.DBConf
}

func NewDB(path, env string) (*DB, error) {
	dbconf, err := goose.NewDBConf(path, env)
	if err != nil {
		return nil, err
	}
	return &DB{DBConf: dbconf}, nil
}

func (db *DB) Close() {
}
