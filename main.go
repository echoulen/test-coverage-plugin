package main

import (
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	file := os.Getenv("PLUGIN_FILE")
	repo := os.Getenv("PLUGIN_REPO")
	host := os.Getenv("PLUGIN_HOST")
	port := os.Getenv("PLUGIN_PORT")
	outputReader, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	output := fmt.Sprintf("%s", outputReader)

	fmt.Printf("========= Reporting =========\n")
	targetUrl := fmt.Sprintf("http://%s:%s/repos/%s", host, port, repo)

	client := &http.Client{}
	postValues := url.Values{}
	postValues.Add("doc", output)
	resp, err := client.PostForm(targetUrl, postValues)
	fmt.Printf("======= Report finish =======")
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}