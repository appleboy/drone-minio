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

		// rm specific flags.
		cli.BoolFlag{
			Name:   "recursive, r",
			Usage:  "Remove recursively.",
			EnvVar: "PLUGIN_RECURSIVE",
		},
		cli.BoolFlag{
			Name:   "force",
			Usage:  "Allow a recursive remove operation.",
			EnvVar: "PLUGIN_FORCE",
		},
		cli.BoolFlag{
			Name:   "dangerous",
			Usage:  "Allow site-wide removal of buckets and objects.",
			EnvVar: "PLUGIN_DANGEROUS",
		},
		cli.BoolFlag{
			Name:   "incomplete, I",
			Usage:  "Remove incomplete uploads.",
			EnvVar: "PLUGIN_INCOMPLETE",
		},
		cli.BoolFlag{
			Name:   "fake",
			Usage:  "Perform a fake remove operation.",
			EnvVar: "PLUGIN_FAKE",
		},
		cli.IntFlag{
			Name:   "older-than",
			Usage:  "Remove objects older than N days",
			EnvVar: "PLUGIN_ORDER_THAN",
		},
		cli.IntFlag{
			Name:   "newer-than",
			Usage:  "Remove objects newer than N days",
			EnvVar: "PLUGIN_NEWER_THAN",
		},
		cli.StringFlag{
			Name:   "encrypt-key",
			Usage:  "Encrypt object (using server-side encryption)",
			EnvVar: "PLUGIN_ENCRYPT_KEY",
		},
		cli.StringFlag{
			Name:   "path",
			Usage:  "object path",
			EnvVar: "PLUGIN_PATH",
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
			Actions:      c.StringSlice("actions"),
			URL:          c.String("url"),
			AccessKey:    c.String("access-key"),
			SecretKey:    c.String("secret-key"),
			IsQuiet:      c.Bool("quiet"),
			IsNoColor:    c.Bool("no-color"),
			IsDebug:      c.Bool("debug"),
			IsJSON:       c.Bool("json"),
			IsInsecure:   c.Bool("insecure"),
			IsForce:      c.Bool("force"),
			IsRecursive:  c.Bool("recursive"),
			IsDangerous:  c.Bool("dangerous"),
			IsIncomplete: c.Bool("incomplete"),
			EncryptKey:   c.String("encrypt-key"),
			OlderThan:    c.Int("older-than"),
			NewerThan:    c.Int("newer-than"),
			IsFake:       c.Bool("fake"),
			Path:         c.String("path"),
		},
	}

	return plugin.Exec()
}
