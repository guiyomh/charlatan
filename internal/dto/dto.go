// Copyright (C) 2022 Guillaume Camus <guillaume.camus@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package dto provides data structures and methodes to manipulate it
package dto

import (
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type TableName string
type RecordID string
type Field string

type Tables map[TableName]Record
type Record map[RecordID]*Row
type Row struct {
	Fields Fields
	Meta   Meta
}
type Fields map[Field]string
type Meta struct {
	RecordID RecordID
	Table    TableName
}

type FixtureName string
type SetID string

type FixtureSet map[FixtureName]Fixture
type Fixture map[SetID]Set
type Set map[Field]EntryValue
type EntryValue interface{}

var (
	faker = gofakeit.New(time.Now().UnixNano())
)

func FakeData(tables Tables) Tables {
	t := Tables{}
	for tableName, record := range tables {
		t[tableName] = Record{}
		for recordID, row := range record {
			t[tableName][recordID] = &Row{
				Fields: make(Fields),
				Meta:   row.Meta, // TODO: make a unit test to cover this case (keep meta)
			}
			for field, value := range row.Fields {
				t[tableName][recordID].Fields[field] = fake(value)
			}
		}
	}

	return t
}

func fake(input string) string {
	return faker.Generate(fmt.Sprint(input))
}

func (r Row) HasDependencies() bool {
	for _, value := range r.Fields {
		if strings.HasPrefix(value, "@") {
			return true
		}
	}

	return false
}

func (r Row) HasDependencyOf(other string) bool {
	for _, value := range r.Fields {
		if value == fmt.Sprintf("@%s", other) {
			return true
		}
	}

	return false
}
