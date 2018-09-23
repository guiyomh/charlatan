package model

import (
	"fmt"
)

// Row represente data that will be persist in database
type Row struct {
	Name                string
	TableName           string
	Fields              map[string]interface{}
	DependencyReference map[string]*Relation
	Left                *Row
	Right               *Row
	Pk                  interface{}
}

// NewRow Factory to build a Row Structure
func NewRow(name, tableName string) *Row {
	return &Row{
		Name:                name,
		TableName:           tableName,
		Fields:              make(map[string]interface{}),
		DependencyReference: make(map[string]*Relation, 0),
		Pk:                  nil,
	}
}

// PrintDebug is a temporary method to dump a Row
func (r Row) PrintDebug() {
	fmt.Printf("\n------\n- %s\n------\n", r.Name)
	fmt.Printf("    Table: %s\n", r.TableName)
	fmt.Println("    Fields:")
	for field, value := range r.Fields {
		fmt.Printf("%15s => %v\n", field, value)
	}
	if r.HasDependencies() {
		fmt.Println("    Deps:")
		for field, value := range r.DependencyReference {
			fmt.Printf("%15s => %v\n", field, value)
		}
	}
}

// HasDependencyOf is True if the Row have a dependancy with the other row pass
func (r *Row) HasDependencyOf(name string) bool {
	if !r.HasDependencies() {
		return false
	}
	for _, v := range r.DependencyReference {
		if v.RecordName == name {
			return true
		}
	}
	return false
}

// HasDependencies return true if this row has dependancies
func (r *Row) HasDependencies() bool {
	return len(r.DependencyReference) >= 0
}

// SetDependance Set the value of a depandance
// func (r *Row) SetDependance(name string, row *Row) {
// 	r.Fields[name] = value
// }
