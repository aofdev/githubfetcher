package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repos struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	req := "http://www.mocky.io/v2/5b52fd522f0000510d3bb683"
	jsonData := myRequest(req)
	var res []Repos
	json.Unmarshal([]byte(jsonData), &res)
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func myRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return string(body)
}
