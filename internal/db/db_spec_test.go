package db

import (
	"testing"
	"time"

	"github.com/guiyomh/charlatan/internal/dto"
	"github.com/stretchr/testify/assert"
)

func TestMysqlPersistor(t *testing.T) {

	row := dto.Row{
		Fields: dto.Fields{
			dto.Field("name"):     "bob",
			dto.Field("number"):   "111-222-333",
			dto.Field("user_id"):  "11",
			dto.Field("birthday"): "1989-10-13T02:49:35Z",
			dto.Field("enable"):   "true",
		},
		Meta: dto.Meta{
			RecordID: dto.RecordID("bob"),
			Table:    dto.TableName("user"),
		},
	}

	expectedValue := Values{
		"name":     "bob",
		"number":   "111-222-333",
		"user_id":  11,
		"enable":   true,
		"birthday": time.Date(1989, time.October, 13, 02, 49, 35, 0, time.Local),
	}
	request, value := convertRowToSQL(row)

	assert.Contains(t, request, "INSERT INTO `user` (")
	assert.Equal(t, expectedValue, value)

}
