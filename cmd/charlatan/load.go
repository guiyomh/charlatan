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
	"strings"

	"go.uber.org/zap"

	"github.com/guiyomh/charlatan/internal/db"
	"github.com/guiyomh/charlatan/internal/loader"
)

type LoadCmd struct {
	Fixtures []string `arg:"" name:"path" help:"Path of the fixtures" type:"path"`
}

func (cmd LoadCmd) Run(logger *zap.Logger, conf DB) error {
	var err error
	var persistor db.Persistor

	logger.With(
		zap.Any("user", conf.User),
		zap.Any("password", strings.Repeat("***", len(conf.Password))),
		zap.Any("host", conf.Host),
		zap.Any("port", conf.Port),
	).
		Debug("Database config")

	persistor, err = db.NewMySQL(logger, conf.User, conf.Password, conf.Name, conf.Host, int(conf.Port))
	if err != nil {
		return err
	}
	defer func() {
		if err = persistor.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	l := loader.New(logger, persistor)
	err = l.Load(cmd.Fixtures...)
	if err != nil {
		return err
	}

	return nil
}
