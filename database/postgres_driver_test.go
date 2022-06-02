package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestPostgresDriver_GetNodeSessionFromHeight(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.NoError(err)

	driver := NewPostgresDriverFromSQLDBInstance(db)

	rows := sqlmock.NewRows([]string{"public_key", "chain", "session_key", "address", "session_height", "total_success", "total_failure", "avg_success_time", "failure"}).
		AddRow("1234", "0001", "1234", "", 1, 2, 2, 2, false)
	mock.ExpectQuery("SELECT (.+) FROM cherry_picker_session").WillReturnRows(rows)

	session, err := driver.GetSessionHeight("1234", 1)
	c.NoError(err)
	c.NotEmpty(session)

	_, err = driver.GetSessionHeight("123410", 1)
	c.Error(err)
}

func TestPostgresDriver_GetNodeSessionFromKey(t *testing.T) {
	c := require.New(t)

	db, mock, err := sqlmock.New()
	c.NoError(err)

	driver := NewPostgresDriverFromSQLDBInstance(db)

	rows := sqlmock.NewRows([]string{"public_key", "chain", "session_key", "address", "session_height", "total_success", "total_failure", "avg_success_time", "failure"}).
		AddRow("1234", "0001", "1234", "", 1, 2, 2, 2, false)
	mock.ExpectQuery("SELECT (.+) FROM cherry_picker_session").WillReturnRows(rows)

	session, err := driver.GetSessionFromKey("1234", "1234")
	c.NoError(err)
	c.NotEmpty(session)

	_, err = driver.GetSessionFromKey("123410", "1234")
	c.Error(err)
}
