package faker

import (
	"reflect"
	"testing"
)

func Test_argConverter(t *testing.T) {
	type args struct {
		value  string
		typeof string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"Testing a integer", args{"13", "int"}, 13},
		{"Testing a bool True", args{"True", "bool"}, true},
		{"Testing a bool true", args{"true", "bool"}, true},
		{"Testing a bool 1", args{"1", "bool"}, true},
		{"Testing a bool False", args{"False", "bool"}, false},
		{"Testing a bool false", args{"false", "bool"}, false},
		{"Testing a bool 0", args{"0", "bool"}, false},
		{"Testing a float", args{"3.14", "float"}, "3.14"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := argConverter(tt.args.value, tt.args.typeof); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("argConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
