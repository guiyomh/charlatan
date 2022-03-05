package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var (
	version string
	commit  string
	date    string
)

func main() {
	var cli struct {
		Fixtures []string `arg:"" name:"path" help:"Path of the fixtures" type:"path"`
		DB       struct {
			User     string `required:"" short:"u" help:"user used to connect to the db"`
			Password string `required:"" short:"p" help:"password to use with the given user to connect to the db"`
			Name     string `required:"" short:"d" help:"name of the schema"`
			Host     string `default:"127.0.0.1" help:"host of the database"`
			Port     int16  `default:"3306" help:"listen database port"`
		} `embed:"" prefix:"db." group:"database"`
	}

	kong.Parse(
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
}
