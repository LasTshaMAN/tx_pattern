package pg

import (
	"database/sql"

	"github.com/pkg/errors"

	// pg driver registers itself as being available to the database/sql package.
	_ "github.com/lib/pq"
)

// Client represents a client to the underlying PostgreSQL data store.
type Client struct {
	db *sql.DB
}

// NewClient returns a new client with open DB connection.
func NewClient(connStr string) (*Client, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open connection to postgres")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "failed to ping postgres through open connection")
	}

	client := Client{
		db: db,
	}

	return &client, nil
}

// Close closes PostgreSQL connection.
func (c *Client) Close() error {
	return c.db.Close()
}
