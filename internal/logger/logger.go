// Package logger provides fast, leveled, structured logging
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New instanciates a logger
func New(options ...Option) (*zap.Logger, error) {

	c := config{
		level:    ErrorLevel,
		encoding: ConsoleEncoding,
	}

	for _, opt := range options {
		c = opt(c)
	}

	l, err := c.level.convertToZapLevel()
	if err != nil {
		return nil, err
	}
	if err = c.encoding.Valid(); err != nil {
		return nil, err
	}
	config := zap.NewProductionConfig()
	config.Encoding = string(c.encoding)
	config.Level = zap.NewAtomicLevelAt(l)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	return config.Build()
}
