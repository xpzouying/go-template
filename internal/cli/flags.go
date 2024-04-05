package cli

import (
	cli "github.com/urfave/cli/v2"
)

var (
	// listenAddr is the address to bind the local server to.
	listenAddr string
	logLevel   string
)

var RootFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "listen",
		Aliases:     []string{"l"},
		Usage:       "Bind the local server to this address.",
		Destination: &listenAddr,
		Required:    false,
		Value:       "127.0.0.1:8080",
	},

	&cli.StringFlag{
		Name:        "log_level",
		Aliases:     []string{"ll"},
		Usage:       "Set the log level.",
		Destination: &logLevel,
		Required:    false,
		Value:       "debug",
	},
}
