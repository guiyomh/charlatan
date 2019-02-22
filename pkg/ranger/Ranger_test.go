package ranger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRanger_BuildRecordName(t *testing.T) {
	type args struct {
		objectSetName string
		quantifier    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test 1",
			args{objectSetName: "foo_", quantifier: "{1..3}"},
			[]string{
				"foo_1",
				"foo_2",
				"foo_3",
			},
		},
		{
			"test 2",
			args{objectSetName: "bar_", quantifier: "{bob,harry,foo}"},
			[]string{
				"bar_bob",
				"bar_harry",
				"bar_foo",
			},
		},
	}
	r := &Ranger{}
	for _, tt := range tests {
		got := r.BuildRecordName(tt.args.objectSetName, tt.args.quantifier)
		assert.Equal(t, tt.want, got)
	}
}

func TestRanger_parseRange(t *testing.T) {
	type args struct {
		quantifier string
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
	}{
		{"test 1", args{quantifier: "{1..10}"}, 1, 10},
		{"test 2", args{quantifier: "{99..111}"}, 99, 111},
	}
	r := &Ranger{}
	for _, tt := range tests {
		gotMin, gotMax := r.parseRange(tt.args.quantifier)
		assert.Equal(t, tt.wantMin, gotMin)
		assert.Equal(t, tt.wantMax, gotMax)
	}
}

func TestRanger_makeRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test 1", args{min: 0, max: 8}, []string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}},
		{"test 2", args{min: 3, max: 8}, []string{"3", "4", "5", "6", "7", "8"}},
	}
	r := &Ranger{}
	for _, tt := range tests {
		got := r.makeRange(tt.args.min, tt.args.max)
		assert.Equal(t, tt.want, got)
	}
}

func TestRanger_parseList(t *testing.T) {
	type args struct {
		quantifier string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test 1", args{quantifier: "{foo,bar,bob,harry}"}, []string{"foo", "bar", "bob", "harry"}},
	}
	r := &Ranger{}
	for _, tt := range tests {
		got := r.parseList(tt.args.quantifier)
		assert.Equal(t, tt.want, got)
	}
}
