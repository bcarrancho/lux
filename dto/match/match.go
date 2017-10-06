package match

type MatchDTO struct {
	SeasonID              int
	QueueID               int
	GameID                uint64
	ParticipantIdentities []*ParticipantIdentityDTO
	GameVersion           string
	PlatformID            string
	GameMode              string
	MapID                 int
	GameType              string
	Teams                 []*TeamStatsDTO
	Participants          []*ParticipantDTO
	GameDuration          uint64
	GameCreation          uint64
}

type ParticipantIdentityDTO struct {
	Player        []*PlayerDTO
	ParticipantID int
}

type PlayerDTO struct {
	CurrentPlatformID string
	SummonerName      string
	MatchHistoryURI   string
	PlatformID        string
	CurrentAccountID  uint64
	ProfileIcon       int
	SummonerID        uint64
	AccountID         uint64
}

type TeamStatsDTO struct {
	FirstDragon          bool
	FirstInhibitor       bool
	Bans                 []*TeamBansDTO
	BaronKills           int
	FirstRiftHerald      bool
	FirstBaron           bool
	RiftHeraldKills      int
	FirstBlood           bool
	TeamID               int
	FirstTower           bool
	VilemawKills         int
	InhibitorKills       int
	TowerKills           int
	DominionVictoryScore int
	Win                  string
	DragonKills          int
}

type TeamBansDTO struct {
	PickTurn   int
	ChampionID int
}

type ParticipantDTO struct {
	Stats                     ParticipantStatsDTO
	ParticipantID             int
	Runes                     []*RuneDTO
	Timeline                  ParticipantTimelineDTO
	TeamID                    int
	Spell2ID                  int
	Masteries                 []*MasteryDTO
	HighestAchievedSeasonTier string
	Spell1ID                  int
	ChampionID                int
}

type ParticipantStatsDTO struct {
	PhysicalDamageDealt             uint64
	NeutralMinionsKilledTeamJungle  int
	MagicDamageDealt                uint64
	TotalPlayerScore                int
	Deaths                          int
	Win                             bool
	NeutralMinionsKilledEnemyJungle int
	AltarsCaptured                  int
	LargestCriticalStrike           int
	TotalDamageDealt                uint64
	MagicDamageDealtToChampions     uint64
	VisionWardsBoughtInGame         int
	DamageDealtToObjectives         uint64
	LargestKillingSpree             int
	Item1                           int
	QuadraKills                     int
	TeamObjective                   int
	TotalTimeCrowdControlDealt      int
	LongestTimeSpentLiving          int
	WardsKilled                     int
	FirstTowerAssist                bool
	FirstTowerKill                  bool
	Item2                           int
	Item3                           int
	Item0                           int
	FirstBloodAssist                bool
	VisionScore                     uint64
	WardsPlaced                     int
	Item4                           int
	Item5                           int
	Item6                           int
	TurretKills                     int
	TripleKills                     int
	DamageSelfMitigated             uint64
	ChampLevel                      int
	NodeNeutralizeAssist            int
	FirstInhibitorKill              bool
	GoldEarned                      int
	MagicalDamageTaken              uint64
	Kills                           int
	DoubleKills                     int
	NodeCaptureAssist               int
	TrueDamageTaken                 uint64
	NodeNeutralize                  int
	FirstInhibitorAssist            bool
	Assists                         int
	UnrealKills                     int
	NeutralMinionsKilled            int
	ObjectivePlayerScore            int
	CombatPlayerScore               int
	DamageDealtToTurrets            uint64
	AltarsNeutralized               int
	PhysicalDamageDealtToChampions  uint64
	GoldSpent                       int
	TrueDamageDealt                 uint64
	TrueDamageDealtToChampions      uint64
	ParticipantID                   int
	PentaKills                      int
	TotalHeal                       uint64
	TotalMinionsKilled              int
	FirstBloodKill                  bool
	NodeCapture                     int
	LargestMultiKill                int
	SightWardsBoughtInGame          int
	RotalDamageDealtToChampions     uint64
	TotalUnitsHealed                int
	TnhibitorKills                  int
	TotalScoreRank                  int
	TotalDamageTaken                uint64
	KillingSprees                   int
	TimeCCingOthers                 uint64
	PhysicalDamageTaken             uint64
}

type RuneDTO struct {
	RuneID int
	Rank   int
}

type ParticipantTimelineDTO struct {
	Lane                        string
	ParticipantID               int
	CSDiffPerMinDeltas          map[string]float64
	GoldPerMinDeltas            map[string]float64
	XPDiffPerMinDeltas          map[string]float64
	CreepsPerMinDeltas          map[string]float64
	XPPerMinDeltas              map[string]float64
	Role                        string
	DamageTakenDiffPerMinDeltas map[string]float64
	DamageTakenPerMinDeltas     map[string]float64
}

type MasteryDTO struct {
	MasteryID int
	Rank      int
}
