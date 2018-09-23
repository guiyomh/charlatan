package generator

import (
	"fmt"
	"strings"

	"github.com/guiyomh/go-faker-fixtures/internal/app/model"
)

// Record use ObjectSet and Template to build Row
type Record struct {
	templates  map[string]*model.Template
	objectSets []*model.ObjectSet
	faker      *Value
}

// NewRecord factory to create an instance of NewRecord
func NewRecord(templates map[string]*model.Template, records []*model.ObjectSet) *Record {
	return &Record{
		templates:  templates,
		objectSets: records,
		faker:      NewValue(),
	}
}

// BuildRecord Generate a Row
func (g *Record) BuildRecord() []*model.Row {
	Rows := make([]*model.Row, 0)
	for _, objectSet := range g.objectSets {

		if objectSet.HasExtend {
			objectSet = g.completeField(objectSet)
		}
		rows := g.createRows(objectSet)
		Rows = append(Rows, rows...)
	}
	return Rows
}

func (g *Record) completeField(set *model.ObjectSet) *model.ObjectSet {
	for fieldName, value := range g.templates[set.ParentName].Fields {
		set.Fields[fieldName] = value
	}
	return set
}

func (g *Record) createRows(objectSet *model.ObjectSet) []*model.Row {
	rows := make([]*model.Row, 0)
	for _, rowReference := range objectSet.RangeRowReference {
		row := g.createRow(rowReference, objectSet)
		rows = append(rows, row)
	}
	return rows
}

func (g *Record) createRow(rowReference string, objectSet *model.ObjectSet) *model.Row {

	current := strings.Replace(rowReference, objectSet.Name, "", 1)
	row := model.NewRow(rowReference, objectSet.TableName)
	for field, value := range objectSet.Fields {
		v := g.generateValue(current, value)
		row.Fields[field] = v
		g.getDependency(field, v, row)
	}
	return row
}

func (g *Record) getDependency(field string, value interface{}, row *model.Row) {
	if _, ok := value.(string); !ok {
		return
	}
	relation, err := model.NewRelation(value.(string))
	if err != nil {
		return
	}
	row.DependencyReference[field] = relation
}

func (g *Record) generateValue(current string, value interface{}) interface{} {
	typeof := fmt.Sprintf("%T", value)
	if typeof == "string" && strings.Contains(value.(string), "<Current()>") {
		value = strings.Replace(value.(string), "<Current()>", current, 1)
	}
	if typeof == "string" {
		value = g.faker.Generate(value.(string))
	}
	return value
}
