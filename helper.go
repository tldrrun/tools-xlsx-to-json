package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func handleError(action cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		err := action(c)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
}

func checkRequired(action cli.ActionFunc, args ...string) cli.ActionFunc {
	return func(c *cli.Context) error {
		for _, arg := range args {
			value := c.String(arg)
			if value == "" {
				return errors.Errorf("Missing '-%s' and it is required.", arg)
			}
		}
		return action(c)
	}
}

func stringTrimSplit(s string) []string {
	output := strings.Split(s, ",")
	for i := range output {
		output[i] = strings.TrimSpace(output[i])
	}
	return output
}

func safeCleanName(s string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	safe := reg.ReplaceAllString(s, "-")
	safe = strings.ToLower(strings.Trim(safe, "-"))
	return safe
}
