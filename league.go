package main

import (
	"encoding/json"
	"log"
	"strings"
)

// LeagueListDTO is returned by the call to
// /lol/league/v3/challengerleagues/by-queue/{queue}
type LeagueListDTO struct {
	Tier    string
	Queue   string
	Name    string
	Entries []*LeaguePositionDTO
}

// LeaguePositionDTO is an entry to a team or player in LeagueListDTO
type LeaguePositionDTO struct {
	Rank             string
	HotStreak        bool
	MiniSeries       []*MiniSeriesDTO
	Wins             int
	Veteran          bool
	Losses           int
	FreshBlood       bool
	PlayerOrTeamName string
	Inactive         bool
	PlayerOrTeamID   string
	LeaguePoints     int
}

// MiniSeriesDTO is an entry to a mini series for a player or team
type MiniSeriesDTO struct {
	Wins     int
	Losses   int
	Target   int
	Progress string
}

func GetChallengerLeagueByQueue(reg, queue string) LeagueListDTO {
	url := "https://" + reg +
		".api.riotgames.com/lol/league/v3/challengerleagues/by-queue/" + queue
	resp := get(url)
	return decodeLeague(resp)
}

func GetMasterLeagueByQueue(reg, queue string) LeagueListDTO {
	url := "https://" + reg +
		".api.riotgames.com/lol/league/v3/masterleagues/by-queue/" + queue
	resp := get(url)
	return decodeLeague(resp)
}

func decodeLeague(resp []byte) LeagueListDTO {
	dec := json.NewDecoder(strings.NewReader(string(resp)))
	var league LeagueListDTO
	for dec.More() {
		err := dec.Decode(&league)
		if err != nil {
			log.Fatal(err)
		}
	}
	return league
}
