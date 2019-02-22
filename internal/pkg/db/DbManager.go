package db

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	// ErrDbDriverIsNotSupported is returned when the db driver is not supported
	ErrDbDriverIsNotSupported = errors.New("This db driver is not supported")

	// ErrWrongCastNotAMap is returned when a map is not a map[interface{}]interface{}
	ErrWrongCastNotAMap = errors.New("Could not cast record: not a map[interface{}]interface{}")

	// ErrKeyIsNotString is returned when a record is not of type string
	ErrKeyIsNotString = errors.New("Record map key is not string")
)

// InsertError will be returned if any error happens on database while
// inserting the record
type InsertError struct {
	Err    error
	SQL    string
	Params []interface{}
}

func (e *InsertError) Error() string {
	return fmt.Sprintf(
		"charlatan: error inserting record: %v, sql: %s, params: %v",
		e.Err,
		e.SQL,
		e.Params,
	)
}

type DbManager interface {
	BuildInsertSQL(schema string, table string, fields map[string]interface{}) (string, []interface{}, error)
	TruncateTable(schema string, table string) (sql.Result, error)
	Exec(sqlStr string, params []interface{}) (sql.Result, error)
}

type DbManagerFactory struct{}

func (dm DbManagerFactory) NewDbManager(
	driverName string,
	host string,
	port int16,
	username string,
	password string,
) (DbManager, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8", username, password, host, port)
	myDb, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	switch driverName {
	case "mysql":
		return &MySQL{connection: myDb}, nil
	}
	return nil, ErrDbDriverIsNotSupported
}
