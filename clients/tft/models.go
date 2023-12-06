package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 2099d6a309eb237d31b9c5fc1071c36eca85d91f

// LeagueListDTO data object.
type LeagueListV1DTO struct {
	LeagueID string            `json:"leagueId,omitempty"`
	Tier     Tier              `json:"tier,omitempty"`
	Name     string            `json:"name,omitempty"`
	Queue    QueueType         `json:"queue,omitempty"`
	Entries  []LeagueItemV1DTO `json:"entries,omitempty"`
}

// LeagueItemDTO data object.
type LeagueItemV1DTO struct {
	SummonerName string   `json:"summonerName,omitempty"`
	Rank         Division `json:"rank,omitempty"`
	// Player's encrypted summonerId.
	SummonerID string          `json:"summonerId,omitempty"`
	MiniSeries MiniSeriesV1DTO `json:"miniSeries,omitempty"`
	// First placement.
	Wins         int32 `json:"wins,omitempty"`
	LeaguePoints int32 `json:"leaguePoints,omitempty"`
	// Second through eighth placement.
	Losses     int32 `json:"losses,omitempty"`
	FreshBlood bool  `json:"freshBlood,omitempty"`
	Inactive   bool  `json:"inactive,omitempty"`
	Veteran    bool  `json:"veteran,omitempty"`
	HotStreak  bool  `json:"hotStreak,omitempty"`
}

// MiniSeriesDTO data object.
type MiniSeriesV1DTO struct {
	Progress string `json:"progress,omitempty"`
	Losses   int32  `json:"losses,omitempty"`
	Target   int32  `json:"target,omitempty"`
	Wins     int32  `json:"wins,omitempty"`
}

// LeagueEntryDTO data object.
type LeagueEntryV1DTO struct {
	// Player Universal Unique Identifier. Exact length of 78 characters. (Encrypted)
	PUUID string `json:"puuid,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	LeagueID string `json:"leagueId,omitempty"`
	// Player's encrypted summonerId.
	SummonerID   string    `json:"summonerId,omitempty"`
	SummonerName string    `json:"summonerName,omitempty"`
	QueueType    QueueType `json:"queueType,omitempty"`
	// Only included for the RANKED_TFT_TURBO queueType.
	// (Legal values:  ORANGE,  PURPLE,  BLUE,  GREEN,  GRAY)
	RatedTier string `json:"ratedTier,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Tier Tier `json:"tier,omitempty"`
	// The player's division within a tier. Not included for the RANKED_TFT_TURBO queueType.
	Rank Division `json:"rank,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	MiniSeries MiniSeriesV1DTO `json:"miniSeries,omitempty"`
	// Only included for the RANKED_TFT_TURBO queueType.
	RatedRating int32 `json:"ratedRating,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	LeaguePoints int32 `json:"leaguePoints,omitempty"`
	// First placement.
	Wins int32 `json:"wins,omitempty"`
	// Second through eighth placement.
	Losses int32 `json:"losses,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	HotStreak bool `json:"hotStreak,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Veteran bool `json:"veteran,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	FreshBlood bool `json:"freshBlood,omitempty"`
	// Not included for the RANKED_TFT_TURBO queueType.
	Inactive bool `json:"inactive,omitempty"`
}

// TopRatedLadderEntryDTO data object.
type TopRatedLadderEntryV1DTO struct {
	SummonerID   string `json:"summonerId,omitempty"`
	SummonerName string `json:"summonerName,omitempty"`
	// (Legal values:  ORANGE,  PURPLE,  BLUE,  GREEN,  GRAY)
	RatedTier   string `json:"ratedTier,omitempty"`
	RatedRating int32  `json:"ratedRating,omitempty"`
	// First placement.
	Wins                         int32 `json:"wins,omitempty"`
	PreviousUpdateLadderPosition int32 `json:"previousUpdateLadderPosition,omitempty"`
}

// MatchDTO data object.
type MatchV1DTO struct {
	// Match metadata.
	Metadata MetadataV1DTO `json:"metadata,omitempty"`
	// Match info.
	Info InfoV1DTO `json:"info,omitempty"`
}

// MetadataDTO data object.
type MetadataV1DTO struct {
	// Match data version.
	DataVersion string `json:"data_version,omitempty"`
	// Match id.
	MatchID string `json:"match_id,omitempty"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants,omitempty"`
}

// InfoDTO data object.
type InfoV1DTO struct {
	TftGameType    string `json:"tft_game_type,omitempty"`
	TftSetCoreName string `json:"tft_set_core_name,omitempty"`
	// Game variation key. Game variations documented in TFT static data.
	GameVariation string `json:"game_variation,omitempty"`
	// Game client version.
	GameVersion  string             `json:"game_version,omitempty"`
	Participants []ParticipantV1DTO `json:"participants,omitempty"`
	// Unix timestamp.
	GameDatetime int64 `json:"game_datetime,omitempty"`
	// Game length in seconds.
	GameLength float32 `json:"game_length,omitempty"`
	// Please refer to the League of Legends documentation.
	QueueID int32 `json:"queue_id,omitempty"`
	// Teamfight Tactics set number.
	TftSetNumber int32 `json:"tft_set_number,omitempty"`
}

// ParticipantDTO data object.
type ParticipantV1DTO struct {
	// Participant's companion.
	Companion CompanionV1DTO `json:"companion,omitempty"`
	PUUID     string         `json:"puuid,omitempty"`
	Augments  []string       `json:"augments,omitempty"`
	// A complete list of traits for the participant's active units.
	Traits []TraitV1DTO `json:"traits,omitempty"`
	// A list of active units for the participant.
	Units          []UnitV1DTO `json:"units,omitempty"`
	PartnerGroupID int32       `json:"partner_group_id,omitempty"`
	// Gold left after participant was eliminated.
	GoldLeft int32 `json:"gold_left,omitempty"`
	// The round the participant was eliminated in. Note: If the player was eliminated in stage 2-1 their last_round would be 5.
	LastRound int32 `json:"last_round,omitempty"`
	// Participant Little Legend level. Note: This is not the number of active units.
	Level int32 `json:"level,omitempty"`
	// Participant placement upon elimination.
	Placement int32 `json:"placement,omitempty"`
	// Number of players the participant eliminated.
	PlayersEliminated int32 `json:"players_eliminated,omitempty"`
	// The number of seconds before the participant was eliminated.
	TimeEliminated float32 `json:"time_eliminated,omitempty"`
	// Damage the participant dealt to other players.
	TotalDamageToPlayers int32 `json:"total_damage_to_players,omitempty"`
}

// TraitDTO data object.
type TraitV1DTO struct {
	// Trait name.
	Name string `json:"name,omitempty"`
	// Number of units with this trait.
	NumUnits int32 `json:"num_units,omitempty"`
	// Current style for this trait. (0 = No style, 1 = Bronze, 2 = Silver, 3 = Gold, 4 = Chromatic)
	Style int32 `json:"style,omitempty"`
	// Current active tier for the trait.
	TierCurrent int32 `json:"tier_current,omitempty"`
	// Total tiers for the trait.
	TierTotal int32 `json:"tier_total,omitempty"`
}

// UnitDTO data object.
type UnitV1DTO struct {
	// This field was introduced in patch 9.22 with data_version 2.
	CharacterID string `json:"character_id,omitempty"`
	// If a unit is chosen as part of the Fates set mechanic, the chosen trait will be indicated by this field. Otherwise this field is excluded from the response.
	Chosen string `json:"chosen,omitempty"`
	// Unit name. This field is often left blank.
	Name      string   `json:"name,omitempty"`
	ItemNames []string `json:"itemNames,omitempty"`
	// A list of the unit's items. Please refer to the Teamfight Tactics documentation for item ids.
	Items []int32 `json:"items,omitempty"`
	// Unit rarity. This doesn't equate to the unit cost.
	Rarity int32 `json:"rarity,omitempty"`
	// Unit tier.
	Tier int32 `json:"tier,omitempty"`
}

// CompanionDTO data object.
type CompanionV1DTO struct {
	ContentID string `json:"content_ID,omitempty"`
	Species   string `json:"species,omitempty"`
	ItemID    int32  `json:"item_ID,omitempty"`
	SkinID    int32  `json:"skin_ID,omitempty"`
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
	// (Legal values:  scheduled,  in_progress,  complete)
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	// (Legal values:  info,  warning,  critical)
	IncidentSeverity string         `json:"incident_severity,omitempty"`
	CreatedAt        string         `json:"created_at,omitempty"`
	ArchiveAt        string         `json:"archive_at,omitempty"`
	UpdatedAt        string         `json:"updated_at,omitempty"`
	Titles           []ContentV1DTO `json:"titles,omitempty"`
	Updates          []UpdateV1DTO  `json:"updates,omitempty"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms []string `json:"platforms,omitempty"`
	ID        int32    `json:"id,omitempty"`
}

// ContentDTO data object.
type ContentV1DTO struct {
	Locale  string `json:"locale,omitempty"`
	Content string `json:"content,omitempty"`
}

// UpdateDTO data object.
type UpdateV1DTO struct {
	Author    string `json:"author,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	// (Legal values: riotclient, riotstatus, game)
	PublishLocations []string       `json:"publish_locations,omitempty"`
	Translations     []ContentV1DTO `json:"translations,omitempty"`
	ID               int32          `json:"id,omitempty"`
	Publish          bool           `json:"publish,omitempty"`
}

// SummonerDTO data object.
//
// represents a summoner
//
// Note: This struct is automatically generated
type SummonerV1DTO struct {
	// Encrypted account ID. Max length 56 characters.
	AccountID string `json:"accountId,omitempty"`
	// Summoner name.
	Name string `json:"name,omitempty"`
	// Encrypted summoner ID. Max length 63 characters.
	ID string `json:"id,omitempty"`
	// Encrypted PUUID. Exact length of 78 characters.
	PUUID string `json:"puuid,omitempty"`
	// Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: summoner name change, summoner level change, or profile icon change.
	RevisionDate int64 `json:"revisionDate,omitempty"`
	// Summoner level associated with the summoner.
	SummonerLevel int64 `json:"summonerLevel,omitempty"`
	// ID of the summoner icon associated with the summoner.
	ProfileIconID int32 `json:"profileIconId,omitempty"`
}
