package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Name = "Google form XLSX to JSON Converter for tools.tldr.run!"
	app.Version = version
	app.EnableBashCompletion = true
	app.Authors = []*cli.Author{{Name: "Madhu Akula"}}
	app.Usage = "An utility to perform custom google form xlsx to json convertion for tools.tldr.run."

	app.Before = func(context *cli.Context) error {
		if context.Bool("verbose") {
			log.SetFlags(0)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}
	app.Compiled = time.Now()
	app.Commands = []*cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "Convert XLSX to JSON",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "input",
					Aliases: []string{"i"},
					Usage:   "Specify input XLSX file path",
				},
				// &cli.StringFlag{
				// 	Name:    "output",
				// 	Aliases: []string{"o"},
				// 	Usage:   "Specify output JSON file path",
				// },
			},
			Action: handleError(checkRequired(convertXLSXToJSON, "input")),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		// Run with '--help' for usage.
		log.Fatal(err)
	}
}
