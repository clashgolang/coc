// Usage:  go run examples/player/main.go player -t <APITOKEN> -p <PLAYERTAG>
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/clashgolang/coc/coc"
	"github.com/urfave/cli/v2"
)

const (
	appName = "coc"
	usage   = "Clash of Clans go library"
)

var (
	commands = []*cli.Command{
		{
			Name:        "clanlabels",
			Usage:       "Lists clan labels",
			Description: "Lists clan labels",
			Action:      getClanLabels,
		},
	}

	// flags are the set of flags supported by the CoC application
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "token",
			Aliases:     []string{"t"},
			EnvVars:     []string{"COC_TOKEN"},
			Usage:       "API token to use for authentication with the Clash of Clans REST server (required)",
			DefaultText: " ",
			Required:    true,
		},
	}
)

func main() {
	app := &cli.App{
		Name:     appName,
		Commands: commands,
		Flags:    flags,
		Usage:    usage,
		Before: func(c *cli.Context) error {
			// Set the token
			token := c.String("token")
			coc.SetToken(token)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// getClanLabels lists clan labels
func getClanLabels(c *cli.Context) error {
	// Get the clan labels
	labels, err := coc.GetClanLabels(nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, league := range labels {
		fmt.Printf("League: %s, ID: %d\n", league.Name, league.ID)
	}

	return nil
}
