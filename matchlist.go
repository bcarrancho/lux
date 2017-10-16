package main

// MatchlistDTO is returned by a call to Matchlist API
type MatchlistDTO struct {
	Matches    []*MatchReferenceDTO
	TotalGames int
	StartIndex int
	EndIndex   int
}

// MatchReferenceDTO corresponds to one individual entry in MatchlistDTO
type MatchReferenceDTO struct {
	Lane       string
	GameID     uint64
	Champion   int
	PlatformID string
	Season     int
	Queue      int
	Role       string
	Timestamp  uint64
}
