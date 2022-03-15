//go:build !spec || test

package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestConvertToZapLevel(t *testing.T) {

	tests := []struct {
		name  string
		level Level
		want  zapcore.Level
	}{
		{
			name:  "Should be ok with error level",
			level: ErrorLevel,
			want:  zapcore.ErrorLevel,
		},
		{
			name:  "Should be ok with warn level",
			level: WarnLevel,
			want:  zapcore.WarnLevel,
		},
		{
			name:  "Should be ok with info level",
			level: InfoLevel,
			want:  zapcore.InfoLevel,
		},
		{
			name:  "Should be ok with debug level",
			level: DebugLevel,
			want:  zapcore.DebugLevel,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := tt.level.convertToZapLevel()
			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, out)
			}
		})
	}

	t.Run("Should return an error on unsupported level", func(t *testing.T) {
		l := Level("Foo")
		expectedErr := "unsupported log Foo level, choice between info, warning, debug or error"
		level, err := l.convertToZapLevel()
		if assert.Error(t, err) {
			assert.EqualError(t, err, expectedErr)
			assert.Empty(t, level)
		}
	})
}
