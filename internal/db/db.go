// Package db  provides method to persist fixtures in databases
package db

import (
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/guiyomh/charlatan/internal/tree"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	connMaxLifeTime = 3 * time.Minute
	maxOpenConn     = 10
	maxIdleConn     = 10
)

type Persistor interface {
	Persist(*tree.Tree) error
	Close() error
}

type mySQLPersistor struct {
	db           *sqlx.DB
	logger       *zap.Logger
	persistedRow map[string]int64
}

func (m mySQLPersistor) Persist(t *tree.Tree) error {
	err := t.InOrderTraverse(m.nodeToSQL)
	if err != nil {
		return err
	}

	return nil
}

func (m *mySQLPersistor) nodeToSQL(node *tree.Node, err error) error {
	if err != nil {
		return err
	}

	sql, values := convertRowToSQL(node.Value())
	values = m.applyRelation(values)

	m.logger.
		With(zap.Any("record_id", node.Value().Meta.RecordID)).
		With(zap.Any("fields", values)).
		Sugar().Debugf("Exec SQL: %s", sql)
	result, err := m.db.NamedExec(sql, values)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	m.persistedRow[string(node.Value().Meta.RecordID)] = id

	return nil
}

func (m *mySQLPersistor) applyRelation(values Values) Values {
	for placeholder, value := range values {
		value := fmt.Sprint(value)
		if !strings.HasPrefix(value, "@") {
			continue
		}
		key := value[1:]
		for identifier, id := range m.persistedRow {
			if identifier == key {
				values[placeholder] = id
			}
		}
	}

	return values
}

func (m *mySQLPersistor) Close() error {
	return m.db.Close()
}

func NewMySQL(logger *zap.Logger, username, password, dbname, host string, port int) (Persistor, error) {

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbname))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxOpenConns(maxOpenConn)
	db.SetMaxIdleConns(maxIdleConn)
	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(err, "Couldn't connect to the database %s:%d", host, port)
	}

	return &mySQLPersistor{
		db:           db,
		logger:       logger,
		persistedRow: make(map[string]int64),
	}, nil
}
