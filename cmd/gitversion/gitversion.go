package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "gitversion"
	app.Usage = "determine semantic version from GIT history"
	app.Action = func(c *cli.Context) error {
		err := validateGitInstalled()
		if err != nil {
			return err
		}

		dir, err := os.Getwd()
		if err != nil {
			return err
		}

		err = validateGitRepo(dir)
		if err != nil {
			return err
		}

		log.Print("Success!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
