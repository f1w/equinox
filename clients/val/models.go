package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 2099d6a309eb237d31b9c5fc1071c36eca85d91f

// ContentDTO data object.
type ContentV1DTO struct {
	Version      string             `json:"version,omitempty"`
	Characters   []ContentItemV1DTO `json:"characters,omitempty"`
	Maps         []ContentItemV1DTO `json:"maps,omitempty"`
	Chromas      []ContentItemV1DTO `json:"chromas,omitempty"`
	Skins        []ContentItemV1DTO `json:"skins,omitempty"`
	SkinLevels   []ContentItemV1DTO `json:"skinLevels,omitempty"`
	Equips       []ContentItemV1DTO `json:"equips,omitempty"`
	GameModes    []ContentItemV1DTO `json:"gameModes,omitempty"`
	Sprays       []ContentItemV1DTO `json:"sprays,omitempty"`
	SprayLevels  []ContentItemV1DTO `json:"sprayLevels,omitempty"`
	Charms       []ContentItemV1DTO `json:"charms,omitempty"`
	CharmLevels  []ContentItemV1DTO `json:"charmLevels,omitempty"`
	PlayerCards  []ContentItemV1DTO `json:"playerCards,omitempty"`
	PlayerTitles []ContentItemV1DTO `json:"playerTitles,omitempty"`
	Acts         []ActV1DTO         `json:"acts,omitempty"`
	Ceremonies   []ContentItemV1DTO `json:"ceremonies,omitempty"`
	// Unknown type, this is a placeholder subject to change.
	Totems []string `json:"totems,omitempty"`
}

// ContentItemDTO data object.
type ContentItemV1DTO struct {
	Name string `json:"name,omitempty"`
	// This field is excluded from the response when a locale is set
	LocalizedNames LocalizedNamesV1DTO `json:"localizedNames,omitempty"`
	ID             string              `json:"id,omitempty"`
	AssetName      string              `json:"assetName,omitempty"`
	// This field is only included for maps and game modes. These values are used in the match response.
	AssetPath string `json:"assetPath,omitempty"`
}

// LocalizedNamesDTO data object.
type LocalizedNamesV1DTO struct {
	ArAe string `json:"ar-AE,omitempty"`
	DeDe string `json:"de-DE,omitempty"`
	EnGb string `json:"en-GB,omitempty"`
	EnUs string `json:"en-US,omitempty"`
	EsEs string `json:"es-ES,omitempty"`
	EsMx string `json:"es-MX,omitempty"`
	FrFr string `json:"fr-FR,omitempty"`
	IdID string `json:"id-ID,omitempty"`
	ItIt string `json:"it-IT,omitempty"`
	JaJp string `json:"ja-JP,omitempty"`
	KoKr string `json:"ko-KR,omitempty"`
	PlPl string `json:"pl-PL,omitempty"`
	PtBr string `json:"pt-BR,omitempty"`
	RuRu string `json:"ru-RU,omitempty"`
	ThTh string `json:"th-TH,omitempty"`
	TrTr string `json:"tr-TR,omitempty"`
	ViVn string `json:"vi-VN,omitempty"`
	ZhCn string `json:"zh-CN,omitempty"`
	ZhTw string `json:"zh-TW,omitempty"`
}

// ActDTO data object.
type ActV1DTO struct {
	Name string `json:"name,omitempty"`
	// This field is excluded from the response when a locale is set
	LocalizedNames LocalizedNamesV1DTO `json:"localizedNames,omitempty"`
	ID             string              `json:"id,omitempty"`
	IsActive       bool                `json:"isActive,omitempty"`
	ParentID       string              `json:"parentId,omitempty"`
	Type           string              `json:"type,omitempty"`
}

// MatchDTO data object.
type MatchV1DTO struct {
	MatchInfo    MatchInfoV1DTO     `json:"matchInfo,omitempty"`
	Players      []PlayerV1DTO      `json:"players,omitempty"`
	Coaches      []CoachV1DTO       `json:"coaches,omitempty"`
	Teams        []TeamV1DTO        `json:"teams,omitempty"`
	RoundResults []RoundResultV1DTO `json:"roundResults,omitempty"`
}

// MatchInfoDTO data object.
type MatchInfoV1DTO struct {
	MatchID            string `json:"matchId,omitempty"`
	MapID              string `json:"mapId,omitempty"`
	ProvisioningFlowID string `json:"provisioningFlowId,omitempty"`
	CustomGameName     string `json:"customGameName,omitempty"`
	QueueID            string `json:"queueId,omitempty"`
	GameMode           string `json:"gameMode,omitempty"`
	SeasonID           string `json:"seasonId,omitempty"`
	GameStartMillis    int64  `json:"gameStartMillis,omitempty"`
	GameLengthMillis   int32  `json:"gameLengthMillis,omitempty"`
	IsCompleted        bool   `json:"isCompleted,omitempty"`
	IsRanked           bool   `json:"isRanked,omitempty"`
}

// PlayerDTO data object.
type PlayerV1DTO struct {
	PUUID           string           `json:"puuid,omitempty"`
	GameName        string           `json:"gameName,omitempty"`
	TagLine         string           `json:"tagLine,omitempty"`
	TeamID          string           `json:"teamId,omitempty"`
	PartyID         string           `json:"partyId,omitempty"`
	CharacterID     string           `json:"characterId,omitempty"`
	Stats           PlayerStatsV1DTO `json:"stats,omitempty"`
	CompetitiveTier int32            `json:"competitiveTier,omitempty"`
	PlayerCard      string           `json:"playerCard,omitempty"`
	PlayerTitle     string           `json:"playerTitle,omitempty"`
}

// PlayerStatsDTO data object.
type PlayerStatsV1DTO struct {
	Score          int32             `json:"score,omitempty"`
	RoundsPlayed   int32             `json:"roundsPlayed,omitempty"`
	Kills          int32             `json:"kills,omitempty"`
	Deaths         int32             `json:"deaths,omitempty"`
	Assists        int32             `json:"assists,omitempty"`
	PlaytimeMillis int32             `json:"playtimeMillis,omitempty"`
	AbilityCasts   AbilityCastsV1DTO `json:"abilityCasts,omitempty"`
}

// AbilityCastsDTO data object.
type AbilityCastsV1DTO struct {
	GrenadeCasts  int32 `json:"grenadeCasts,omitempty"`
	Ability1Casts int32 `json:"ability1Casts,omitempty"`
	Ability2Casts int32 `json:"ability2Casts,omitempty"`
	UltimateCasts int32 `json:"ultimateCasts,omitempty"`
}

// CoachDTO data object.
type CoachV1DTO struct {
	PUUID  string `json:"puuid,omitempty"`
	TeamID string `json:"teamId,omitempty"`
}

// TeamDTO data object.
type TeamV1DTO struct {
	// This is an arbitrary string. Red and Blue in bomb modes. The puuid of the player in deathmatch.
	TeamID       string `json:"teamId,omitempty"`
	Won          bool   `json:"won,omitempty"`
	RoundsPlayed int32  `json:"roundsPlayed,omitempty"`
	RoundsWon    int32  `json:"roundsWon,omitempty"`
	// Team points scored. Number of kills in deathmatch.
	NumPoints int32 `json:"numPoints,omitempty"`
}

// RoundResultDTO data object.
type RoundResultV1DTO struct {
	RoundNum      int32  `json:"roundNum,omitempty"`
	RoundResult   string `json:"roundResult,omitempty"`
	RoundCeremony string `json:"roundCeremony,omitempty"`
	WinningTeam   string `json:"winningTeam,omitempty"`
	// PUUID of player
	BombPlanter string `json:"bombPlanter,omitempty"`
	// PUUID of player
	BombDefuser           string                  `json:"bombDefuser,omitempty"`
	PlantRoundTime        int32                   `json:"plantRoundTime,omitempty"`
	PlantPlayerLocations  []PlayerLocationsV1DTO  `json:"plantPlayerLocations,omitempty"`
	PlantLocation         LocationV1DTO           `json:"plantLocation,omitempty"`
	PlantSite             string                  `json:"plantSite,omitempty"`
	DefuseRoundTime       int32                   `json:"defuseRoundTime,omitempty"`
	DefusePlayerLocations []PlayerLocationsV1DTO  `json:"defusePlayerLocations,omitempty"`
	DefuseLocation        LocationV1DTO           `json:"defuseLocation,omitempty"`
	PlayerStats           []PlayerRoundStatsV1DTO `json:"playerStats,omitempty"`
	RoundResultCode       string                  `json:"roundResultCode,omitempty"`
}

// PlayerLocationsDTO data object.
type PlayerLocationsV1DTO struct {
	PUUID       string        `json:"puuid,omitempty"`
	ViewRadians float32       `json:"viewRadians,omitempty"`
	Location    LocationV1DTO `json:"location,omitempty"`
}

// LocationDTO data object.
type LocationV1DTO struct {
	X int32 `json:"x,omitempty"`
	Y int32 `json:"y,omitempty"`
}

// PlayerRoundStatsDTO data object.
type PlayerRoundStatsV1DTO struct {
	PUUID   string        `json:"puuid,omitempty"`
	Kills   []KillV1DTO   `json:"kills,omitempty"`
	Damage  []DamageV1DTO `json:"damage,omitempty"`
	Score   int32         `json:"score,omitempty"`
	Economy EconomyV1DTO  `json:"economy,omitempty"`
	Ability AbilityV1DTO  `json:"ability,omitempty"`
}

// KillDTO data object.
type KillV1DTO struct {
	TimeSinceGameStartMillis  int32 `json:"timeSinceGameStartMillis,omitempty"`
	TimeSinceRoundStartMillis int32 `json:"timeSinceRoundStartMillis,omitempty"`
	// PUUID
	Killer string `json:"killer,omitempty"`
	// PUUID
	Victim         string        `json:"victim,omitempty"`
	VictimLocation LocationV1DTO `json:"victimLocation,omitempty"`
	// List of PUUIDs
	Assistants      []string               `json:"assistants,omitempty"`
	PlayerLocations []PlayerLocationsV1DTO `json:"playerLocations,omitempty"`
	FinishingDamage FinishingDamageV1DTO   `json:"finishingDamage,omitempty"`
}

// FinishingDamageDTO data object.
type FinishingDamageV1DTO struct {
	DamageType          string `json:"damageType,omitempty"`
	DamageItem          string `json:"damageItem,omitempty"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode,omitempty"`
}

// DamageDTO data object.
type DamageV1DTO struct {
	// PUUID
	Receiver  string `json:"receiver,omitempty"`
	Damage    int32  `json:"damage,omitempty"`
	Legshots  int32  `json:"legshots,omitempty"`
	Bodyshots int32  `json:"bodyshots,omitempty"`
	Headshots int32  `json:"headshots,omitempty"`
}

// EconomyDTO data object.
type EconomyV1DTO struct {
	LoadoutValue int32  `json:"loadoutValue,omitempty"`
	Weapon       string `json:"weapon,omitempty"`
	Armor        string `json:"armor,omitempty"`
	Remaining    int32  `json:"remaining,omitempty"`
	Spent        int32  `json:"spent,omitempty"`
}

// AbilityDTO data object.
type AbilityV1DTO struct {
	GrenadeEffects  string `json:"grenadeEffects,omitempty"`
	Ability1Effects string `json:"ability1Effects,omitempty"`
	Ability2Effects string `json:"ability2Effects,omitempty"`
	UltimateEffects string `json:"ultimateEffects,omitempty"`
}

// MatchlistDTO data object.
type MatchlistV1DTO struct {
	PUUID   string                `json:"puuid,omitempty"`
	History []MatchlistEntryV1DTO `json:"history,omitempty"`
}

// MatchlistEntryDTO data object.
type MatchlistEntryV1DTO struct {
	MatchID             string `json:"matchId,omitempty"`
	GameStartTimeMillis int64  `json:"gameStartTimeMillis,omitempty"`
	QueueID             string `json:"queueId,omitempty"`
}

// RecentMatchesDTO data object.
type RecentMatchesV1DTO struct {
	CurrentTime int64 `json:"currentTime,omitempty"`
	// A list of recent match ids.
	MatchIDs []string `json:"matchIds,omitempty"`
}

// LeaderboardDTO data object.
type LeaderboardV1DTO struct {
	// The shard for the given leaderboard.
	Shard string `json:"shard,omitempty"`
	// The act id for the given leaderboard. Act ids can be found using the val-content API.
	ActID string `json:"actId,omitempty"`
	// The total number of players in the leaderboard.
	TotalPlayers          int64                     `json:"totalPlayers,omitempty"`
	Players               []PlayerV1DTO             `json:"players,omitempty"`
	ImmortalStartingPage  int64                     `json:"immortalStartingPage,omitempty"`
	ImmortalStartingIndex int64                     `json:"immortalStartingIndex,omitempty"`
	TopTierRrThreshold    int64                     `json:"topTierRRThreshold,omitempty"`
	TierDetails           map[int64]TierDetailV1DTO `json:"tierDetails,omitempty"`
	StartIndex            int64                     `json:"startIndex,omitempty"`
	Query                 string                    `json:"query,omitempty"`
}

// PlayerDTO data object.
type MatchPlayerV1DTO struct {
	// This field may be omitted if the player has been anonymized.
	PUUID string `json:"puuid,omitempty"`
	// This field may be omitted if the player has been anonymized.
	GameName string `json:"gameName,omitempty"`
	// This field may be omitted if the player has been anonymized.
	TagLine         string `json:"tagLine,omitempty"`
	LeaderboardRank int64  `json:"leaderboardRank,omitempty"`
	RankedRating    int64  `json:"rankedRating,omitempty"`
	NumberOfWins    int64  `json:"numberOfWins,omitempty"`
	CompetitiveTier int64  `json:"competitiveTier,omitempty"`
}

// TierDetailDTO data object.
type TierDetailV1DTO struct {
	RankedRatingThreshold int64 `json:"rankedRatingThreshold,omitempty"`
	StartingPage          int64 `json:"startingPage,omitempty"`
	StartingIndex         int64 `json:"startingIndex,omitempty"`
}

// PlatformDataDTO data object.
type PlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
}

// StatusDTO data object.
type StatusV1DTO struct {
	ID int32 `json:"id,omitempty"`
	// (Legal values:  scheduled,  in_progress,  complete)
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	// (Legal values:  info,  warning,  critical)
	IncidentSeverity string         `json:"incident_severity,omitempty"`
	Titles           []ContentV1DTO `json:"titles,omitempty"`
	Updates          []UpdateV1DTO  `json:"updates,omitempty"`
	CreatedAt        string         `json:"created_at,omitempty"`
	ArchiveAt        string         `json:"archive_at,omitempty"`
	UpdatedAt        string         `json:"updated_at,omitempty"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms []string `json:"platforms,omitempty"`
}

// ContentDTO data object.
type StatusContentV1DTO struct {
	Locale  string `json:"locale,omitempty"`
	Content string `json:"content,omitempty"`
}

// UpdateDTO data object.
type UpdateV1DTO struct {
	ID      int32  `json:"id,omitempty"`
	Author  string `json:"author,omitempty"`
	Publish bool   `json:"publish,omitempty"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string       `json:"publish_locations,omitempty"`
	Translations     []ContentV1DTO `json:"translations,omitempty"`
	CreatedAt        string         `json:"created_at,omitempty"`
	UpdatedAt        string         `json:"updated_at,omitempty"`
}
