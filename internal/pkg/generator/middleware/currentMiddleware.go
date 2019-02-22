package middleware

import (
	"fmt"
	"strings"
)

func CurrentMiddleware(current string) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		typeof := fmt.Sprintf("%T", value)
		if typeof == "string" && strings.Contains(value.(string), "<Current()>") {
			value = strings.Replace(value.(string), "<Current()>", current, 1)
		}
		return value
	}
}
