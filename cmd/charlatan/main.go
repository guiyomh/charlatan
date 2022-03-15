// Copyright (C) 2018 Guillaume Camus <guillaume.camus@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"github.com/guiyomh/charlatan/internal/logger"
)

var (
	version string
	commit  string
	date    string
)

type DB struct {
	User     string `required:"" short:"u" help:"user used to connect to the db"`
	Password string `required:"" short:"p" help:"password to use with the given user to connect to the db"`
	Name     string `required:"" short:"d" help:"name of the schema"`
	Host     string `default:"127.0.0.1" help:"host of the database"`
	Port     int16  `default:"3306" help:"listen database port"`
}

func main() {
	var cli struct {
		DB      `embed:"" prefix:"db." group:"database"`
		Logging struct {
			Level string `enum:"debug,info,warn,error" default:"error" help:"level of logging (debug,info,warn,error)"`
		} `embed:"" prefix:"log."`
		Load LoadCmd `cmd:"" help:"Load fixtures from the paths"`
	}

	ctx := kong.Parse(
		&cli,
		kong.Name("charlatan"),
		kong.Description(fmt.Sprintf(
			"charlatan is a very fast fixtures loaders.\n\nversion: %s (%s) - %s",
			version,
			commit,
			date,
		)),
		kong.UsageOnError(),
	)

	logger, err := logger.New(
		logger.OptionLevel(level(cli.Logging.Level)),
	)
	if err != nil {
		ctx.FatalIfErrorf(err)
	}

	err = ctx.Run(logger, cli.DB)
	ctx.FatalIfErrorf(err)
}

func level(level string) logger.Level {
	switch level {
	case "debug":
		return logger.DebugLevel
	case "info":
		return logger.InfoLevel
	case "warn":
		return logger.WarnLevel
	case "error":
		return logger.ErrorLevel
	default:
		return logger.ErrorLevel
	}
}
