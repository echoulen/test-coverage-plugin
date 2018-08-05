package main

import (
	"github.com/urfave/cli"
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "test coverage plugin"
	app.Usage = "test coverage plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "file",
			Usage:  "files for coverage upload",
			EnvVar: "PLUGIN_FILES",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	file := c.String("file")
	fmt.Printf("file: %s", file)
	console, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("file content: %s", string(console))
	// TODO
	return err
}