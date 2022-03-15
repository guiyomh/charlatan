package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangerMakeNumericList(t *testing.T) {
	r := ranger{}

	tests := []struct {
		name string
		args string
		want []string
	}{
		{
			name: "Should start from 1",
			args: "1..5",
			want: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "Should start from 2",
			args: "2..5",
			want: []string{"2", "3", "4", "5"},
		},
		{
			name: "Should start from 5",
			args: "5..10",
			want: []string{"5", "6", "7", "8", "9", "10"},
		},
		{
			name: "Should start from 5 to 5",
			args: "5..5",
			want: []string{"5"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list, err := r.makeNumericList(tt.args)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, list)
			}
		})
	}
}
