package dto

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

var (
	isRange           = regexp.MustCompile(`{[a-z0-9\.,]+\}`)
	numericRangeRegex = regexp.MustCompile(`(?m)(?P<min>[0-9]+)\.\.(?P<max>[0-9]+)`)
)

type ranger struct{}

func (r ranger) Range(id SetID) (list []string, iterators []string, err error) {

	name := string(id)
	prefix := name[:strings.Index(name, "{")]
	quantifier := name[strings.Index(name, "{")+1 : strings.Index(name, "}")]

	if strings.Contains(quantifier, "..") {
		iterators, err = r.makeNumericList(quantifier)
		if err != nil {
			return []string{}, []string{}, err
		}
	} else if strings.Contains(quantifier, ",") {
		iterators = strings.Split(quantifier, ",")
	}
	list = make([]string, len(iterators))
	for k, value := range iterators {
		list[k] = fmt.Sprintf("%s%s", prefix, value)
	}

	return list, iterators, nil
}

func (r ranger) makeNumericList(quantifier string) ([]string, error) {
	groups := numericRangeRegex.FindAllStringSubmatch(quantifier, -1)
	if len(groups) == 0 || len(groups[0]) != 3 {
		return []string{}, errors.Errorf("Cannot parse the range of record %s", quantifier)
	}
	min, err := strconv.Atoi(groups[0][1])
	if err != nil {
		return []string{}, err
	}
	max, err := strconv.Atoi(groups[0][2])
	if err != nil {
		return []string{}, err
	}
	list := make([]string, 0)
	for i := min; i <= max; i++ {
		list = append(list, strconv.Itoa(i))
	}

	return list, nil
}

func MakeRange(id SetID) (list []string, iterators []string, err error) {
	r := ranger{}

	return r.Range(id)
}
