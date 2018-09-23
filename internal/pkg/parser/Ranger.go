package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Ranger parse a object set labe to find a range
type Ranger struct {
}

var numericRangeRegex = regexp.MustCompile(`(?m)\{(?P<min>[0-9]+)\.\.(?P<max>[0-9]+)\}`)

//BuildRecordName parse the name of an ObjectSet an generate
// a list of RecordName
func (r *Ranger) BuildRecordName(objectSetName, quantifier string) []string {
	var list []string
	if strings.Contains(quantifier, "..") {
		min, max := r.parseRange(quantifier)
		list = r.makeRange(min, max)
	} else if strings.Contains(quantifier, ",") {
		list = r.parseList(quantifier)
	}
	for k, value := range list {
		list[k] = fmt.Sprintf("%s%s", objectSetName, value)
	}
	return list
}

func (r *Ranger) parseRange(quantifier string) (min, max int) {

	groups := numericRangeRegex.FindAllStringSubmatch(quantifier, -1)[0]

	min, _ = strconv.Atoi(groups[1])
	max, _ = strconv.Atoi(groups[2])

	return
}

func (r *Ranger) makeRange(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = strconv.Itoa(min + i)
	}
	return a
}

func (r *Ranger) parseList(quantifier string) []string {
	return strings.Split(quantifier[1:len(quantifier)-1], ",")
}
