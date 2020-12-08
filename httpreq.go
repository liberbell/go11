package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	url1 = "https://www.yahoo.co.jp"

	resp, _ := http.Get(url1)
	defer resp.Body.close()

	batArray, _ := ioutil.ReadAll(resp.Body)
}
