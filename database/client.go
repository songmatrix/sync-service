package database

type driver interface {
	// establish connection to the database
	connect() error

	// ensure database schema matches,
	//
	// TODO - if not, apply migrations?
	assertSchema() error

	// TODO - how long do we keep events?
	// DeleteExpired(ttl int64) error

	// terminate database connection
	Close()
}

func Init() (driver, error) { return &mockDriver{}, nil }
