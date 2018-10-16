package db

import "github.com/guiyomh/go-faker-fixtures/internal/app/model"

type Generator interface {
	ToSQL(row *model.Row) (int64, error)
}
