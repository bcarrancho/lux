package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mixxtokent/lux/dto/league"
)

func main() {
	reg := "br1"
	apikey := os.Getenv("APIKEY")
	lurl := "https://" + reg + ".api.riotgames.com/lol/league/v3/masterleagues/by-queue/RANKED_FLEX_SR"

	cli := http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest(http.MethodGet, lurl, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("X-Riot-Token", apikey)
	resp, err := cli.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	for dec.More() {
		var league league.LeagueListDTO
		err := dec.Decode(&league)
		if err != nil {
			log.Fatal(err)
		}

		for _, ent := range league.Entries {
			fmt.Println(ent.PlayerOrTeamID)
		}
	}
}
