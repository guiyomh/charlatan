package logger

type config struct {
	level    Level    `required:"true" default:"error" values:"info,warning,debug,error" desc:"log level grapping"`
	encoding Encoding `required:"true" default:"json" values:"json,console" desc:"print method"`
}

type Option func(config config) config

func OptionLevel(level Level) Option {
	return func(c config) config {
		c.level = level

		return c
	}
}

func OptionEncoding(encoding Encoding) Option {
	return func(c config) config {
		c.encoding = encoding

		return c
	}
}
