package database

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

const (
	getNodeSessionFromKey = `SELECT * FROM cherry_picker_session WHERE public_key = $1 
		AND session_key = $2`
	getNodeSessionFromHeight = `SELECT * FROM cherry_picker_session WHERE public_key = $1 
		AND session_height = $2`
)

var (
	// ErrMissingPublicKey when public key is empty
	ErrMissingPublicKey = errors.New("missing publicKey field")
	// ErrMissingSessionKeyAndHeight when session key and height fields are empty
	ErrMissingSessionKeyAndHeight = errors.New("missing sessionKeight and sessionheight fields")
)

type PostgresDriver struct {
	*sqlx.DB
}

// NewPostgresDriverFromConnectionString returns PostgresDriver instance from connection string
func NewPostgresDriverFromConnectionString(connectionString string) (*PostgresDriver, error) {
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &PostgresDriver{
		DB: db,
	}, nil
}

// NewPostgresDriverFromSQLDBInstance returns PostgresDriver instance from sdl.DB instance
// mostly used for mocking tests
func NewPostgresDriverFromSQLDBInstance(db *sql.DB) *PostgresDriver {
	return &PostgresDriver{
		DB: sqlx.NewDb(db, "postgres"),
	}
}

func (pg *PostgresDriver) GetSessionHeight(publicKey string, sessionHeight int) (*Session, error) {
	var session Session
	return &session, pg.Get(&session, getNodeSessionFromHeight, publicKey, sessionHeight)
}

func (pg *PostgresDriver) GetSessionFromKey(publicKey, sessionKey string) (*Session, error) {
	var session Session
	return &session, pg.Get(&session, getNodeSessionFromHeight, publicKey, sessionKey)
}
