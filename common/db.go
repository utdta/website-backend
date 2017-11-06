package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

const (
	DB_TYPE = `mysql`
)

var (
	DBUser    = os.Getenv("API_DB_USER")
	DBName    = os.Getenv("API_DB_NAME")
	DBPwd     = os.Getenv("API_DB_PWD")
	DBHost    = os.Getenv("API_DB_HOST")
	DBPort    = os.Getenv("API_DB_PORT")
	DBSession *sqlx.DB
)

type ErrDBConnInvalid struct {
	DBUser,
	DBPwd,
	DBHost,
	DBPort,
	DBName string
}

func (e ErrDBConnInvalid) Error() string {
	return fmt.Sprintf(`One or more database connection variables is invalid: user=%s; password=%s; host=%s; port=%s; name=%s`,
		e.DBUser, e.DBPwd, e.DBHost, e.DBPort, e.DBName)
}

// Registers a database and confirms a connection
func GetDBConnection() (*sqlx.DB, error) {
	connStr, envErr := getDBConnStr()
	if envErr != nil {
		return nil, envErr
	}

	// Connect to the database
	db, dbErr := sqlx.Connect(DB_TYPE, connStr)
	if dbErr != nil {
		return nil, dbErr
	}
	return db, nil
}

// Creates the database connection string
func getDBConnStr() (string, error) {
	if DBUser == `` || DBPwd == `` || DBHost == `` || DBPort == `` || DBName == `` {
		return ``, ErrDBConnInvalid{DBUser, DBPwd, DBHost, DBPort, DBName}
	}
	return fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8`,
		DBUser,
		DBPwd,
		DBHost,
		DBPort,
		DBName), nil
}
