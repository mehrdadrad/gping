package main

import (
	"errors"
	"os"

	cli "github.com/urfave/cli/v2"
)

func getCli() (params, error) {
	p := params{}

	cli.AppHelpTemplate = `
	gRPC Ping {{.Version}}

	usage:
	{{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .Commands}} host {{end}} 
	{{if len .Authors}}
	COMMANDS:
	{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
	options:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
	{{end}}
`

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "-server or -s",
			EnvVars: []string{"GPING_SERVER"},
		},
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Usage:   "-count or -c",
			Value:   4,
			EnvVars: []string{"GPING_COUNT"},
		},
		&cli.IntFlag{
			Name:    "ttl",
			Aliases: []string{"t"},
			Usage:   "-ttl or -t",
			Value:   64,
			EnvVars: []string{"GPING_TTL"},
		},
		&cli.StringFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Usage:   "-intervale or -i 2s",
			Value:   "1s",
			EnvVars: []string{"GPING_TTL"},
		},
		&cli.StringFlag{
			Name:    "remote",
			Aliases: []string{"r"},
			Usage:   "-remote 192.168.10.12:3055",
			EnvVars: []string{"GPING_REMOTE"},
		},
		&cli.StringFlag{
			Name:    "bind",
			Aliases: []string{"b"},
			Usage:   "-bind 192.168.10.12:3055",
			Value:   ":3055",
			EnvVars: []string{"GPING_BIND"},
		},
	}
	app := &cli.App{
		Flags: flags,
		Action: func(c *cli.Context) error {
			p.mode = c.Bool("server")
			p.count = c.Int("count")
			p.ttl = c.Int("ttl")
			p.interval = c.String("interval")

			p.host = c.Args().Get(0)
			if c.NArg() < 1 && !p.mode {
				cli.ShowAppHelp(c)
				return errors.New("host not specified")
			}

			return nil
		},
		Version: "v0.1.0",
	}

	err := app.Run(os.Args)

	return p, err
}
