package main

import (
	"errors"
	"net"
	"os"
	"time"

	cli "github.com/urfave/cli/v2"
)

type params struct {
	mode       bool
	json       bool
	silent     bool
	privileged bool
	isLogs     bool
	hosts      []string
	src        string
	bind       string
	remote     string
	count      int
	ttl        int
	tos        int
	size       int
	interval   string
	timeout    string
}

func getCli() (params, error) {
	p := params{}

	cli.AppHelpTemplate = `
	gRPC Ping {{.Version}}

	usage:
	{{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .Commands}} host(s) {{end}} 
	{{if len .Authors}}
	COMMANDS:
	{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
	options:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright }}
	{{end}}
`

	flags := []cli.Flag{
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Usage:   "sets ping count",
			Value:   4,
			EnvVars: []string{"GPING_COUNT"},
		},
		&cli.IntFlag{
			Name:    "ttl",
			Aliases: []string{"t"},
			Usage:   "sets the IP Time to Live",
			Value:   64,
			EnvVars: []string{"GPING_TTL"},
		},
		&cli.IntFlag{
			Name:    "tos",
			Aliases: []string{"q"},
			Usage:   "sets quality of service in ICMP datagram",
			Value:   0,
			EnvVars: []string{"GPING_TOS"},
		},
		&cli.IntFlag{
			Name:    "size",
			Usage:   "sets the number of data bytes to be sent (data + ICMP header)",
			Value:   64,
			EnvVars: []string{"GPING_SIZE"},
		},
		&cli.StringFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Usage:   "sets wait between sending each packet in format ns,us,ms,s",
			Value:   "1s",
			EnvVars: []string{"GPING_INTERVAL"},
		},
		&cli.StringFlag{
			Name:    "timeout",
			Aliases: []string{"W"},
			Usage:   "sets time to wait for an ICMP reply in format ns,us,ms,s",
			Value:   "2s",
			EnvVars: []string{"GPING_TIMEOUT"},
		},
		&cli.StringFlag{
			Name:    "remote",
			Aliases: []string{"r"},
			Usage:   "sets remote server IP_ADDR:PORT",
			EnvVars: []string{"GPING_REMOTE"},
		},
		&cli.BoolFlag{
			Name:    "json",
			Usage:   "prints statistics in json format",
			EnvVars: []string{"GPING_JSON"},
		},
		&cli.BoolFlag{
			Name:    "silent",
			Usage:   "prints just statistics",
			EnvVars: []string{"GPING_SILENT"},
		},
		&cli.BoolFlag{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "runs server",
			EnvVars: []string{"GPING_SERVER"},
		},
		&cli.StringFlag{
			Name:    "bind",
			Aliases: []string{"b"},
			Usage:   "sets bind IP_ADDR:PORT [server]",
			Value:   "0.0.0.0:3055",
			EnvVars: []string{"GPING_BIND"},
		},
		&cli.BoolFlag{
			Name:    "privileged",
			Aliases: []string{"p"},
			Usage:   "enables ICMP privileged mode [server]",
			EnvVars: []string{"GPING_PRIVILEGED"},
		},
		&cli.BoolFlag{
			Name:    "logs",
			Usage:   "enables logging [server]",
			EnvVars: []string{"GPING_LOGS"},
		},
	}
	app := &cli.App{
		Flags: flags,
		Action: func(c *cli.Context) error {
			p = params{
				mode:       c.Bool("server"),
				bind:       c.String("bind"),
				remote:     c.String("remote"),
				count:      c.Int("count"),
				ttl:        c.Int("ttl"),
				tos:        c.Int("qos"),
				size:       c.Int("size"),
				json:       c.Bool("json"),
				silent:     c.Bool("silent"),
				privileged: c.Bool("privileged"),
				isLogs:     c.Bool("logs"),
				interval:   c.String("interval"),
				timeout:    c.String("timeout"),
			}

			p.hosts = c.Args().Slice()
			if c.NArg() < 1 && !p.mode {
				cli.ShowAppHelp(c)
				return errors.New("host not specified")
			}

			if len(p.remote) < 1 && !p.mode {
				cli.ShowAppHelp(c)
				return errors.New("remote not specified")
			}

			if _, err := time.ParseDuration(p.interval); err != nil {
				return err
			}

			if _, err := time.ParseDuration(p.timeout); err != nil {
				return err
			}

			if _, _, err := net.SplitHostPort(p.remote); err != nil && !p.mode {
				return err
			}

			if _, _, err := net.SplitHostPort(p.bind); err != nil {
				return err
			}

			return nil
		},
		Version: "v0.2.0",
	}

	err := app.Run(os.Args)

	return p, err
}
