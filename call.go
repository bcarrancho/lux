package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func call(url string) (int, []byte) {
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

	switch resp.StatusCode {
	case 200:
		fmt.Println(resp.TransferEncoding)
		var json []byte
		resp.Body.Read(json)
		return 200, json
	case 404:
		fmt.Println("404 NOT FOUND")
		fmt.Println(resp.TransferEncoding)
		return 404, nil
	default:
		return 0, nil
	}

}

func GetMatch(reg string, matid uint64) (int, []byte) {
	url := "https://" + reg + ".api.riotgames.com" +
		"/lol/match/v3/matches/" + strconv.FormatUint(matid, 10)

	code, json := call(url)
	return code, json
}

func GetMatchlist(reg string, accid uint64, queue string) (int, []byte) {
	url := "https://" + reg + ".api.riotgames.com" +
		"/lol/match/v3/matchlists/by-account/" + strconv.FormatUint(accid, 10) +
		

	code, json := call(url)
	return code, json
}
