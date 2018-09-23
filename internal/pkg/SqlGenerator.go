package pkg

import (
	"database/sql"
	"fmt"

	"github.com/azer/logger"
	"github.com/guiyomh/go-faker-fixtures/internal/app/model"

	"gopkg.in/doug-martin/goqu.v5"

	// Register some standard stuff
	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

// This is a helper to build a Sql from a row
type SqlGenerator struct {
	db     *goqu.Database
	logger *logger.Logger
}

// NewSqlGenerator Create a SqlGenerator
func NewSqlGenerator(driverName, dataSource string) *SqlGenerator {
	myDb, err := sql.Open(driverName, dataSource)
	if err != nil {
		panic(err.Error())
	}
	db := goqu.New(driverName, myDb)
	return &SqlGenerator{db: db, logger: logger.New("sql-generator")}
}

// ToSQL
func (s *SqlGenerator) ToSQL(row *model.Row) (int64, error) {
	insert := s.db.From(row.TableName).Insert(row.Fields)
	s.logger.Info(insert.Sql)
	result, err := insert.Exec()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	lastInsertID, err := result.LastInsertId()
	s.logger.Info("last insert Id : %d", lastInsertID)

	return lastInsertID, err
}
