package generator

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"

	fake "github.com/manveru/faker"
)

// Value is a helper to generate fake data
type Value struct {
	faker *fake.Faker
}

var regexFaker, _ = regexp.Compile(`(?i)<(?P<method>[^\(]+)\((?P<args>[^\)]+)?\)>`)

// NewValue is a factory to create a generator of value
func NewValue() *Value {
	fakerInstance, _ := fake.New("en")
	return &Value{
		faker: fakerInstance,
	}
}

// Generate is a method that use faker to generate ramdom data
func (g *Value) Generate(value string) interface{} {

	fakesData := regexFaker.FindStringSubmatch(value)
	if len(fakesData) <= 0 {
		return value
	}
	method := fakesData[1]
	args := make([]string, 0)
	if fakesData[2] != "" {
		args = strings.Split(fakesData[2], ",")
	}

	return g.invoke(method, args)
}

func (g *Value) invoke(method string, args []string) interface{} {
	inputs := make([]reflect.Value, len(args))
	m := reflect.ValueOf(g.faker).MethodByName(method)
	for i := range args {
		inputs[i] = reflect.ValueOf(argConverter(args[i], m.Type().In(i).String()))
	}
	result := m.Call(inputs)[0]
	value, ok := result.Interface().([]string)
	if ok {
		return strings.Join(value, " ")
	}
	return result.Interface()
}

func argConverter(value string, typeof string) interface{} {

	switch typeof {
	case "bool":
		castValue, _ := strconv.ParseBool(value)
		return castValue
	case "int":
		castValue, _ := strconv.Atoi(value)
		return castValue
	default:
		return value
	}
}
