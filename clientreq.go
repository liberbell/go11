package main

import "net/http"

func main() {
	url1 := "https://www.yahoo.co.jp"

	req, _ := http.NewRequest("GET", url1, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
}
