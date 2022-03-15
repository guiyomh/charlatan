package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/guiyomh/charlatan/internal/dto"
)

type Values map[string]interface{}

const (
	bitSize = 64
)

func convertRowToSQL(row dto.Row) (string, Values) {
	sqlColumns := make([]string, 0, len(row.Fields))
	placeholders := make([]string, 0, len(row.Fields))
	values := make(Values)

	for field, value := range row.Fields {
		sqlColumns = append(sqlColumns, string(field))
		placeholder := fmt.Sprintf(":%s", field)
		placeholders = append(placeholders, placeholder)
		values[placeholder[1:]] = convertField(value)
	}
	sql := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES (%s)",
		string(row.Meta.Table),
		strings.Join(sqlColumns, ", "),
		strings.Join(placeholders, ", "),
	)

	return sql, values
}

func convertField(value string) interface{} {

	if number, err := strconv.Atoi(value); err == nil {
		return number
	}

	if number, err := strconv.ParseFloat(value, bitSize); err == nil {
		return number
	}

	if flag, err := strconv.ParseBool(value); err == nil {
		return flag
	}

	if time, err := tryStrToDate(value); err == nil {
		return time
	}

	return value
}

func tryStrToDate(s string) (time.Time, error) {
	timeFormats := []string{
		"2006-01-02",
		"2006-01-02 15:04",
		"2006-01-02 15:04:05",
		"20060102",
		"20060102 15:04",
		"20060102 15:04:05",
		"02/01/2006",
		"02/01/2006 15:04",
		"02/01/2006 15:04:05",
		"2006-01-02T15:04-07:00",
		"2006-01-02T15:04:05-07:00",
		"2006-01-02T15:04:05Z",
	}
	for _, f := range timeFormats {
		t, err := time.ParseInLocation(f, s, time.Local)
		if err != nil {
			continue
		}

		return t, nil
	}

	return time.Time{}, ErrCouldNotConvertToTime
}
