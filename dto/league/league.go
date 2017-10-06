package league

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
