package model

// ObjectSet contains metadatas necessary to build row
type ObjectSet struct {
	TableName           string
	HasExtend           bool
	Name                string
	RangeReference      string
	ParentName          string
	Fields              map[string]interface{}
	RangeRowReference   []string
	DependancyReference []string
}

// NewObjectSet is a factory to create an ObjectSet
func NewObjectSet(tableName string, name string, fields map[string]interface{}, hasExtend bool, rangeReference string, parentName string) *ObjectSet {
	return &ObjectSet{
		TableName:           tableName,
		HasExtend:           hasExtend,
		Fields:              fields,
		Name:                name,
		RangeReference:      rangeReference,
		ParentName:          parentName,
		RangeRowReference:   make([]string, 0),
		DependancyReference: make([]string, 0),
	}
}
