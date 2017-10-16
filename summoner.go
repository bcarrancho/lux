package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

type SummonerDTO struct {
	ProfileIconID int
	Name          string
	SummonerLevel int
	RevisionDate  int
	ID            uint64
	AccountID     uint64
}

func GetSummonerByID(reg string, sid uint64) SummonerDTO {
	url := "https://" + reg +
		".api.riotgames.com/lol/summoner/v3/summoners/" +
		strconv.FormatUint(sid, 10)
	resp := get(url)

	dec := json.NewDecoder(strings.NewReader(string(resp)))
	var summoner SummonerDTO
	for dec.More() {
		err := dec.Decode(&summoner)
		if err != nil {
			log.Fatal(err)
		}
	}
	return summoner
}
