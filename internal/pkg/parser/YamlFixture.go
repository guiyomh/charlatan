package parser

import (
	"regexp"

	"github.com/guiyomh/go-faker-fixtures/internal/app/model"
	"gopkg.in/yaml.v2"
)

type tables map[string]records
type records map[string]fields
type fields map[string]interface{}

type YamlFixture struct {
}

var (
	objectSetRegex, _ = regexp.Compile(`(?i)^(?P<record>[a-z0-9-_]+)(?P<quantifier>\{[a-z0-9\.,]+\})?( \(((?P<isTemplate>template)|(extends (?P<template>[a-z0-9-_]+))\)))?`)
	ranger            = &Ranger{}
)

// Load the content of yamlfile and return a set of template and a set of records
func (p *YamlFixture) Load(content []byte) (map[string]*model.Template, []*model.ObjectSet, error) {
	tbls := tables{}
	err := yaml.Unmarshal(content, tbls)
	tpls, objs := p.parse(tbls)
	return tpls, objs, err
}

func (p *YamlFixture) parse(tbls tables) (map[string]*model.Template, []*model.ObjectSet) {
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
					objectSet.RangeRowReference = ranger.BuildRecordName(name, rangeRef)
				} else {
					objectSet.RangeRowReference = []string{objectSet.Name}
				}
				objs = append(objs, objectSet)
			}
		}
	}
	return tpls, objs
}
