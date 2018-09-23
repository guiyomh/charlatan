package model

import (
	"errors"
	"regexp"
)

var relationRegex, _ = regexp.Compile(`(?i)^\@(?P<rowname>[a-z0-9\._-]+)\.?(?P<fieldname>[a-z0-9\._-]+)?$`)

var (
	errNotFoundRelation = errors.New("No relation found")
)

// Relation represents a relation to an another record
type Relation struct {
	RecordName string
	FieldName  string
}

//NewRelation create a relation structure
func NewRelation(value string) (*Relation, error) {
	deps := relationRegex.FindStringSubmatch(value)
	if len(deps) <= 0 {
		return nil, errNotFoundRelation
	}
	relation := &Relation{RecordName: deps[1]}
	if deps[2] != "" {
		relation.FieldName = deps[2]
	}
	return relation, nil
}
