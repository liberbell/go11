package main

import "net/http"

func main() {
	url1 := "https://www.yahoo.co.jp"

	req, _ := http.NewRequest("GET", url1, nil)
	req.Header.Set("Authorization", "Bearer access-token")

	client := net(http.Client)
	resp, err := client.Do(req)
}
