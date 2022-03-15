package dto

import (
	"fmt"
	"strings"
)

func ConvertFixtureSetToTable(fixtureSet FixtureSet) (Tables, error) {
	tables := Tables{}
	for name, fixture := range fixtureSet {
		record, err := ConvertFixtureToRecord(fixture, name)
		if err != nil {
			return Tables{}, err
		}
		tables[TableName(name)] = record

	}

	return tables, nil
}

func ConvertSetToRow(s Set, meta Meta) *Row {
	r := Row{
		Fields: make(Fields),
		Meta:   meta,
	}
	for field, value := range s {
		r.Fields[field] = fmt.Sprint(value)
	}

	return &r
}

func ConvertFixtureToRecord(f Fixture, name FixtureName) (Record, error) {
	record := Record{}
	tableName := TableName(name)
	f = applyTemplate(f)
	for id, set := range f {
		if strings.Contains(string(id), "(template)") {
			continue
		}
		if !isRange.Match([]byte(id)) {
			meta := Meta{
				RecordID: RecordID(id),
				Table:    tableName,
			}
			record[meta.RecordID] = ConvertSetToRow(set, meta)

			continue
		}
		recordIds, iterators, err := MakeRange(id)
		if err != nil {
			return Record{}, err
		}
		row := ConvertSetToRow(set, Meta{
			Table: tableName,
		})
		for i, rID := range recordIds {
			currentRow := replaceCurrent(iterators[i], row)
			currentRow.Meta.RecordID = RecordID(rID)
			record[currentRow.Meta.RecordID] = currentRow
		}

	}

	return record, nil
}

func replaceCurrent(current string, row *Row) *Row {
	newRow := &Row{Fields: make(Fields)}
	for field, value := range row.Fields {
		newRow.Fields[field] = strings.ReplaceAll(fmt.Sprint(value), "{current}", current)
	}
	newRow.Meta = row.Meta

	return newRow
}

func applyTemplate(fixture Fixture) Fixture {
	tpls := make(map[string]Set)
	for id, set := range fixture {
		if strings.Contains(string(id), "(template)") {
			tplName := strings.TrimSpace(strings.ReplaceAll(string(id), "(template)", ""))
			tpls[tplName] = set
		}
	}
	newFixture := Fixture{}
	for id, set := range fixture {
		newID := string(id)
		if strings.Contains(string(id), "(template)") {
			continue
		}
		if strings.Contains(string(id), "(extends") {
			tplName := string(id)
			tplName = strings.TrimSpace(tplName[strings.Index(tplName, "(extends")+9 : strings.Index(tplName, ")")])
			if tpl, ok := tpls[tplName]; ok {
				set = mergeSet(tpl, set)
			}

			newID = strings.TrimSpace(newID[:strings.Index(newID, "(extends")])
		}
		newFixture[SetID(newID)] = set
	}

	return newFixture
}

func mergeSet(sets ...Set) Set {
	res := Set{}
	for _, set := range sets {
		for k, v := range set {
			res[k] = v
		}
	}

	return res
}
