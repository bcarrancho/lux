package main

import (
    "os"
    "fmt"
    "net/http"
    "log"
    "time"
    "ioutil"
)

func main() {
    apikey := os.Getenv("APIKEY")
    fmt.Println(apikey)

    region := "kr"
    urlbase := "https://" + region + ".api.riotgames.com/lol/"
    // urllm := urlbase + "league/v3/challengerleagues/by-queue/"
    urllc := urlbase + "league/v3/masterleagues/by-queue/"
    queue := "RANKED_FLEX_SR"

    urlltest := urllc + queue
    cli := http.Client{
        Timeout: time.Second * 60,
    }
    req, err := http.NewRequest(http.MethodGet, urlltest, nil)
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Set("X-Riot-Token", apikey)
    res, getErr := cli.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }
    fmt.Println(res)
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }
}
