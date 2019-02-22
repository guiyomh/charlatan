package middleware

import (
	"fmt"

	"github.com/guiyomh/charlatan/pkg/faker"
)

func FakerMiddleware(faker *faker.Value) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		typeof := fmt.Sprintf("%T", value)
		if typeof == "string" {
			value = faker.Generate(value.(string))
		}
		return value
	}
}
