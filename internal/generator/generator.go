package generator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/guiyomh/go-faker-fixtures/internal/app/model"
	"github.com/guiyomh/go-faker-fixtures/pkg/faker"
	"github.com/guiyomh/go-faker-fixtures/pkg/ranger"
)

var (
	objectSetRegex, _ = regexp.Compile(`(?i)^(?P<record>[a-z0-9-_]+)(?P<quantifier>\{[a-z0-9\.,]+\})?( \(((?P<isTemplate>template)|(extends (?P<template>[a-z0-9-_]+))\)))?`)
	myRanger          = ranger.NewRanger()
)

type Generator struct {
	faker *faker.Value
}

// NewGenerator factory to create a Generator
func NewGenerator() *Generator {
	return &Generator{
		faker: faker.NewValue(),
	}
}

// GenerateRecords build records from fixture
func (g Generator) GenerateRecords(data model.FixtureTables) ([]*model.Row, error) {
	tpls, recordSets := g.classifyData(data)
	rows, err := g.buildRecord(tpls, recordSets)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (g Generator) classifyData(tbls model.FixtureTables) (map[string]*model.Template, []*model.ObjectSet) {
	tpls := make(map[string]*model.Template)
	objs := make([]*model.ObjectSet, 0)
	for tableName, records := range tbls {
		for recordName, fields := range records {

			groups := objectSetRegex.FindAllStringSubmatch(recordName, -1)[0]
			name := groups[1]
			rangeRef := groups[2]
			isTemplate := groups[4] == "template"
			hasExtend := groups[7] != ""
			parent := groups[7]
			if isTemplate {
				tpls[name] = model.NewTemplate(name, fields)
			} else {
				objectSet := model.NewObjectSet(tableName, name, fields, hasExtend, rangeRef, parent)
				if rangeRef != "" {
					objectSet.RangeRowReference = myRanger.BuildRecordName(name, rangeRef)
				} else {
					objectSet.RangeRowReference = []string{objectSet.Name}
				}
				objs = append(objs, objectSet)
			}
		}
	}
	return tpls, objs
}

func (g Generator) buildRecord(templates map[string]*model.Template, recordSets []*model.ObjectSet) ([]*model.Row, error) {
	rows := make([]*model.Row, 0)
	for _, objectSet := range recordSets {

		if objectSet.HasExtend {
			objectSet = g.completeField(templates, objectSet)
		}
		rowSets := g.createRows(objectSet)
		rows = append(rows, rowSets...)
	}
	return rows, nil
}

func (g Generator) completeField(templates map[string]*model.Template, set *model.ObjectSet) *model.ObjectSet {
	for fieldName, value := range templates[set.ParentName].Fields {
		set.Fields[fieldName] = value
	}
	return set
}

func (g Generator) createRows(objectSet *model.ObjectSet) []*model.Row {
	rows := make([]*model.Row, 0)
	for _, rowReference := range objectSet.RangeRowReference {
		row := g.createRow(rowReference, objectSet)
		rows = append(rows, row)
	}
	return rows
}

func (g Generator) createRow(rowReference string, objectSet *model.ObjectSet) *model.Row {

	current := strings.Replace(rowReference, objectSet.Name, "", 1)
	row := model.NewRow(rowReference, objectSet.TableName)
	for field, value := range objectSet.Fields {
		v := g.generateValue(current, value)
		row.Fields[field] = v
		g.getDependency(field, v, row)
	}
	return row
}

func (g Generator) getDependency(field string, value interface{}, row *model.Row) {
	if _, ok := value.(string); !ok {
		return
	}
	relation, err := model.NewRelation(value.(string))
	if err != nil {
		return
	}
	row.DependencyReference[field] = relation
}

func (g Generator) generateValue(current string, value interface{}) interface{} {
	typeof := fmt.Sprintf("%T", value)
	if typeof == "string" && strings.Contains(value.(string), "<Current()>") {
		value = strings.Replace(value.(string), "<Current()>", current, 1)
	}
	if typeof == "string" {
		value = g.faker.Generate(value.(string))
	}
	return value
}
