//go:build !spec || test

package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	t.Run("Should return an error because encoding used is unsupported", func(t *testing.T) {
		expectedErr := fmt.Sprintf("unsupported log %s encoding, choice between %s or %s",
			"foo",
			string(JSONEncoding),
			string(ConsoleEncoding),
		)
		e := Encoding("foo")
		err := e.Valid()
		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
		}
	})
	t.Run("Should be ok with json encoding", func(t *testing.T) {
		e := JSONEncoding
		err := e.Valid()
		assert.NoError(t, err)
	})

	t.Run("Should be ok with console encoding", func(t *testing.T) {
		e := ConsoleEncoding
		err := e.Valid()
		assert.NoError(t, err)
	})
}
