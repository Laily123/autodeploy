package main

import (
	"autodeploy/feature"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "autodeploy"
	app.Usage = "auto deploy project"
	app.Version = "0.1"

	serverCmd := cli.Command{
		Name:  "server",
		Usage: "start a server for webhook",
		Action: func(c *cli.Context) error {
			feature.Server(c.String("config"))
			return nil
		},
	}
	serverCmd.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "config file path for server",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "watch",
			Usage: "watch local project",
			Action: func(c *cli.Context) error {
				feature.StartWatcher(os.Args[0], c.Args())
				return nil
			},
		},
		serverCmd,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
