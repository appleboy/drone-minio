package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"
)

// Version set at compile-time
var (
	Version  string
	BuildNum string
)

func main() {
	app := cli.NewApp()
	app.Name = "minio client plugin"
	app.Usage = "Drone plugin to upload or remove filesystems and object storage."
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:   "actions",
			Usage:  "a list of actions to have packer perform",
			EnvVar: "PLUGIN_ACTIONS",
		},
		cli.BoolFlag{
			Name:   "quiet, q",
			Usage:  "Disable progress bar display.",
			EnvVar: "PLUGIN_QUIET",
		},
		cli.BoolFlag{
			Name:   "no-color",
			Usage:  "Disable color theme.",
			EnvVar: "PLUGIN_NOCOLOR",
		},
		cli.BoolFlag{
			Name:   "json",
			Usage:  "Enable JSON formatted output.",
			EnvVar: "PLUGIN_JSON",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Enable debug output.",
			EnvVar: "PLUGIN_DEBUG",
		},
		cli.BoolFlag{
			Name:   "insecure",
			Usage:  "Disable SSL certificate verification.",
			EnvVar: "PLUGIN_INSECURE",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "s3 url",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name:   "access-key",
			Usage:  "s3 access key",
			EnvVar: "PLUGIN_ACCESS_KEY",
		},
		cli.StringFlag{
			Name:   "secret-key",
			Usage:  "s3 secret key",
			EnvVar: "PLUGIN_SECRET_KEY",
		},
	}

	app.Version = Version

	if BuildNum != "" {
		app.Version = app.Version + "+" + BuildNum
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("can't start app")
	}
}

func run(c *cli.Context) error {
	var version string
	if BuildNum != "" {
		version = Version + "+" + BuildNum
	} else {
		version = Version
	}
	log.Info().Str("revision", version).Msg("Drone Minio Plugin Version")

	plugin := Plugin{
		Config: Config{
			Actions:   c.StringSlice("actions"),
			URL:       c.String("url"),
			AccessKey: c.String("access-key"),
			SecretKey: c.String("secret-key"),
			Quiet:     c.Bool("quiet"),
			NoColor:   c.Bool("no-color"),
			Debug:     c.Bool("debug"),
			JSON:      c.Bool("json"),
			Insecure:  c.Bool("insecure"),
		},
	}

	return plugin.Exec()
}
