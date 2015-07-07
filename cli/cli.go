package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/mikebeyer/clc-sdk/clc"
	"github.com/mikebeyer/clc-sdk/cli/server"
)

func main() {
	client := clc.New(clc.EnvConfig())

	app := cli.NewApp()
	app.Name = "clc"
	app.Usage = "clc v2 api cli"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mike Beyer",
			Email: "michael.beyer@centurylink.com",
		},
	}
	app.Commands = []cli.Command{
		server.Commands(client),
	}
	app.Run(os.Args)
}
