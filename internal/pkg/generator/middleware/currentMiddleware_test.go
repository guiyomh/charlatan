package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentMiddleware(t *testing.T) {
	type args struct {
		current string
		value   interface{}
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test 1", args{current: "foo", value: "<Current()>"}, "foo"},
		{"test 2", args{current: "foo", value: "Bar_<Current()>"}, "Bar_foo"},
		{"test 3", args{current: "foo", value: 123}, 123},
		{"test 4", args{current: "foo", value: "test"}, "test"},
		{"test 5", args{current: "foo", value: "<current()>"}, "<current()>"},
	}
	for _, tt := range tests {
		mid := CurrentMiddleware(tt.args.current)
		got := mid(tt.args.value)
		assert.Equal(t, tt.want, got)
	}
}
