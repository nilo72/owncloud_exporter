package main

import (
	"errors"
	"github.com/go-kit/kit/log"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
	"github.com/joho/godotenv"
	"github.com/nilo72/owncloud_exporter/pkg/action"
	"github.com/nilo72/owncloud_exporter/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// set during compile time
var version = "0.00"

var (
	// ErrMissingOwncloudURL defines the error if owncloud.url is empty.
	ErrMissingOwncloudURL = errors.New("Missing required owncloud.url")
)

func main() {
	cfg := config.Load()

	if env := os.Getenv("OWNCLOUD_EXPORTER_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:    "owncloud_exporter",
		Version: version,
		Usage:   "ownCloud Exporter",
		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
			{
				Name:  "Christian Bargmann",
				Email: "christian.bargmann@haw-hamburg.de",
			},
			{
				Name:  "Oliver Neumann",
				Email: "oliver.neumann@haw-hamurg.de",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log.level",
				Value:       "info",
				Usage:       "Only log messages with given severity",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_LOG_LEVEL"},
				Destination: &cfg.Logs.Level,
			},
			&cli.BoolFlag{
				Name:        "log.pretty",
				Value:       false,
				Usage:       "Enable pretty messages for logging",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_LOG_PRETTY"},
				Destination: &cfg.Logs.Pretty,
			},
			&cli.StringFlag{
				Name:        "web.address",
				Value:       "0.0.0.0:9507",
				Usage:       "Address to bind the metrics server",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_WEB_ADDRESS"},
				Destination: &cfg.Server.Addr,
			},
			&cli.StringFlag{
				Name:        "web.path",
				Value:       "/metrics",
				Usage:       "Path to bind the metrics server",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_WEB_PATH"},
				Destination: &cfg.Server.Path,
			},
			&cli.DurationFlag{
				Name:        "request.timeout",
				Value:       5 * time.Second,
				Usage:       "Request timeout as duration",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_REQUEST_TIMEOUT"},
				Destination: &cfg.Target.Timeout,
			},
			&cli.StringFlag{
				Name:        "owncloud.url",
				Value:       "",
				Usage:       "URL to access the ownCloud to scrape",
				EnvVars:     []string{"OWNCLOUD_EXPORTER_URL"},
				Destination: &cfg.Target.Address,
			},
		},
		Action: func(c *cli.Context) error {
			logger := setupLogger(cfg)

			if cfg.Target.Address == "" {
				level.Error(logger).Log(
					"msg", ErrMissingOwncloudURL,
				)

				return ErrMissingOwncloudURL
			}
			return action.Server(cfg, logger)
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show help",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}

func setupLogger(cfg *config.Config) log.Logger {
	var logger log.Logger

	if cfg.Logs.Pretty {
		logger = log.NewSyncLogger(
			log.NewLogfmtLogger(os.Stdout),
		)
	} else {
		logger = log.NewSyncLogger(
			log.NewJSONLogger(os.Stdout),
		)
	}

	switch strings.ToLower(cfg.Logs.Level) {
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "info":
		logger = level.NewFilter(logger, level.AllowInfo())
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	default:
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	return log.With(
		logger,
		"ts", log.DefaultTimestampUTC,
	)
}
