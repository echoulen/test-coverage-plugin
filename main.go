package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/urfave/cli"
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
			Name:   "build.link",
			Usage:  "build.link",
			EnvVar: "DRONE_BUILD_LINK",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}


}

func run(c *cli.Context) error {
	file := os.Getenv("PLUGIN_FILE")
	repo := os.Getenv("PLUGIN_REPO")
	host := os.Getenv("PLUGIN_HOST")
	port := os.Getenv("PLUGIN_PORT")
	outputReader, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("%s", outputReader)
	buildLink := c.String("build.link")

	fmt.Printf("========= Reporting =========\n")
	targetUrl := fmt.Sprintf("http://%s:%s/repos/%s", host, port, repo)

	client := &http.Client{}
	postValues := url.Values{}
	postValues.Add("doc", output)
	postValues.Add("url", buildLink)
	resp, err := client.PostForm(targetUrl, postValues)
	fmt.Printf("======= Report finished =======")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}

	defer resp.Body.Close()
	return nil
}