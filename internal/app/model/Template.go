package model

// Template Container to describe the common fields shared
// between ObjectSet
type Template struct {
	Name   string
	Fields map[string]interface{}
}

// NewTemplate is a factory to create a Template structure
func NewTemplate(Name string, Fields map[string]interface{}) *Template {
	return &Template{
		Name:   Name,
		Fields: Fields,
	}
}
