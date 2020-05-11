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
			Usage:   "-server or -s runs server",
			EnvVars: []string{"GPING_SERVER"},
		},
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Usage:   "-count or -c sets ping count",
			Value:   4,
			EnvVars: []string{"GPING_COUNT"},
		},
		&cli.IntFlag{
			Name:    "ttl",
			Aliases: []string{"t"},
			Usage:   "-ttl or -t sets time to live outgoing packets",
			Value:   64,
			EnvVars: []string{"GPING_TTL"},
		},
		&cli.IntFlag{
			Name:    "size",
			Usage:   "-size bytes (data + ICMP header)",
			Value:   64,
			EnvVars: []string{"GPING_SIZE"},
		},
		&cli.StringFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Usage:   "-intervale or -i sets time interval in format ns,us,ms,s",
			Value:   "1s",
			EnvVars: []string{"GPING_INTERVAL"},
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
			Value:   "0.0.0.0:3055",
			EnvVars: []string{"GPING_BIND"},
		},
		&cli.BoolFlag{
			Name:    "json",
			Usage:   "-json prints statistics in json format",
			EnvVars: []string{"GPING_JSON"},
		},
		&cli.BoolFlag{
			Name:    "silent",
			Usage:   "-silent prints just statistics",
			EnvVars: []string{"GPING_SILENT"},
		},
	}
	app := &cli.App{
		Flags: flags,
		Action: func(c *cli.Context) error {
			p = params{
				mode:     c.Bool("server"),
				bind:     c.String("bind"),
				remote:   c.String("remote"),
				count:    c.Int("count"),
				ttl:      c.Int("ttl"),
				size:     c.Int("size"),
				json:     c.Bool("json"),
				silent:   c.Bool("silent"),
				interval: c.String("interval"),
			}

			p.host = c.Args().Get(0)
			if c.NArg() < 1 && !p.mode {
				cli.ShowAppHelp(c)
				return errors.New("host not specified")
			}

			if len(p.remote) < 1 && !p.mode {
				cli.ShowAppHelp(c)
				return errors.New("remote not specified")
			}

			return nil
		},
		Version: "v0.2.0",
	}

	err := app.Run(os.Args)

	return p, err
}
