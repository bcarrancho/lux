package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const reg = "kr"

var apikey string
var cli http.Client
var sid, aid, mid chan uint64

func main() {

	apikey = os.Getenv("APIKEY")

	sid = make(chan uint64, 1000)
	aid = make(chan uint64, 1000)
	mid = make(chan uint64, 1000)
	cli = http.Client{
		Timeout: time.Second * 60,
	}
	l := NewLimiter(10, 10)
	l.Aquire()

	go report()

	league := GetChallengerLeagueByQueue(reg, "RANKED_FLEX_SR")
	for _, ent := range league.Entries {
		u, err := strconv.ParseUint(ent.PlayerOrTeamID, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		sid <- u
	}
	league = GetMasterLeagueByQueue(reg, "RANKED_FLEX_SR")
	for _, ent := range league.Entries {
		u, err := strconv.ParseUint(ent.PlayerOrTeamID, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		sid <- u
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func report() {
	for {
		fmt.Printf("sid: %d, aid: %d, mid: %d\n", len(sid), len(aid), len(mid))
		time.Sleep(1 * time.Second)
	}
}
