package cmd

import (
	"fmt"

	"github.com/broodbear/tools/scripts"
	"github.com/urfave/cli/v2"
)

var Version = ""

var App = &cli.App{
	Name:  "bear",
	Usage: "single binary containing useful scripts",
	Action: func(*cli.Context) error {
		fmt.Printf("%s by broodbear\n", Version)
		return nil
	},
	Version: Version,
	Commands: []*cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "configure the environment",
			Action: func(cCtx *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "nmap",
			Aliases: []string{"n"},
			Usage:   "run nmap",
			Action: func(cCtx *cli.Context) error {
				fmt.Println(scripts.Nmap)
				return nil
			},
		},
	},
}
