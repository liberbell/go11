package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	url1 := "https://www.yahoo.co.jp"

	req, _ := http.NewRequest("GET", url1, nil)
	req.Header.Set("Authorization", "Bearer access-token")

	client := net(http.Client)
	resp, err := client.Do(req)

	dumpResp, _ := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", dumpResp)
}
