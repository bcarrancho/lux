package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) []byte {
	fmt.Println(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Riot-Token", apikey)
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	deb := fmt.Sprintf("HTTP/%d: ", resp.StatusCode)
	switch resp.StatusCode {
	case 429:
		fmt.Println("*** RATE LIMITED ***")
		if lt, ok := resp.Header["X-Rate-Limit-Type"]; ok {
			deb += fmt.Sprintf("XRateLType: %s ", lt)
		}
		if ra, ok := resp.Header["Retry-After"]; ok {
			deb += fmt.Sprintf("RetryAfter: %s ", ra)
		}
	}
	deb += fmt.Sprintf("AppL: %s, AppLCount: %s, MetL: %s, MetLCount: %s",
		resp.Header["X-App-Rate-Limit"],
		resp.Header["X-App-Rate-Limit-Count"],
		resp.Header["X-Method-Rate-Limit"],
		resp.Header["X-Method-Rate-Limit-Count"])
	fmt.Println(deb)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
