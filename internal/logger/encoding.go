package logger

import "github.com/pkg/errors"

// Supported encoding.
const (
	JSONEncoding    Encoding = "json"
	ConsoleEncoding Encoding = "console"
)

// Encoding logger.
type Encoding string

// Valid validates the encoding value
func (e Encoding) Valid() error {
	switch e {
	case JSONEncoding:
		return nil
	case ConsoleEncoding:
		return nil
	default:
		return errors.Errorf("unsupported log %s encoding, choice between %s or %s",
			e,
			string(JSONEncoding),
			string(ConsoleEncoding),
		)
	}
}
