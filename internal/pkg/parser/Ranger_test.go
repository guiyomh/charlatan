package parser

import (
	"reflect"
	"testing"
)

func TestRanger_BuildRecordName(t *testing.T) {
	type args struct {
		objectSetName string
		quantifier    string
	}
	tests := []struct {
		name string
		r    *Ranger
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ranger{}
			if got := r.BuildRecordName(tt.args.objectSetName, tt.args.quantifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.BuildRecordName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRanger_parseRange(t *testing.T) {
	type args struct {
		quantifier string
	}
	tests := []struct {
		name    string
		r       *Ranger
		args    args
		wantMin int
		wantMax int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ranger{}
			gotMin, gotMax := r.parseRange(tt.args.quantifier)
			if gotMin != tt.wantMin {
				t.Errorf("Ranger.parseRange() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("Ranger.parseRange() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func TestRanger_makeRange(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		r    *Ranger
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ranger{}
			if got := r.makeRange(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.makeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRanger_parseList(t *testing.T) {
	type args struct {
		quantifier string
	}
	tests := []struct {
		name string
		r    *Ranger
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ranger{}
			if got := r.parseList(tt.args.quantifier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ranger.parseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
