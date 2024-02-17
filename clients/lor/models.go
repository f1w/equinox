package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 48735a0c9d1c521d94a20ff0b0b9dc927ab430ca

// CardDto data object.
type CardV1DTO struct {
	Code  string `json:"code,omitempty"`
	Count string `json:"count,omitempty"`
}

// ContentDto data object.
type ContentV1DTO struct {
	Content string `json:"content,omitempty"`
	Locale  string `json:"locale,omitempty"`
}

// DeckDto data object.
type DeckV1DTO struct {
	Code string `json:"code,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// InfoDto data object.
type InfoV1DTO struct {
	// (Legal values:  standard,  eternal)
	GameFormat string `json:"game_format,omitempty"`
	// (Legal values:  Constructed,  Expeditions,  Tutorial)
	GameMode         string `json:"game_mode,omitempty"`
	GameStartTimeUtc string `json:"game_start_time_utc,omitempty"`
	// (Legal values:  Ranked,  Normal,  AI,  Tutorial,  VanillaTrial,  Singleton,  StandardGauntlet)
	GameType    string        `json:"game_type,omitempty"`
	GameVersion string        `json:"game_version,omitempty"`
	Players     []PlayerV1DTO `json:"players,omitempty"`
	// Total turns taken by both players.
	TotalTurnCount int32 `json:"total_turn_count,omitempty"`
}

// PlayerDto data object.
type LeaderboardPlayerV1DTO struct {
	Name string `json:"name,omitempty"`
	// League points.
	LP   int32 `json:"lp,omitempty"`
	Rank int32 `json:"rank,omitempty"`
}

// LeaderboardDto data object.
type LeaderboardV1DTO struct {
	// A list of players in Master tier.
	Players []LeaderboardPlayerV1DTO `json:"players,omitempty"`
}

// MatchDto data object.
type MatchV1DTO struct {
	// Match metadata.
	Metadata MetadataV1DTO `json:"metadata,omitempty"`
	// Match info.
	Info InfoV1DTO `json:"info,omitempty"`
}

// MetadataDto data object.
type MetadataV1DTO struct {
	// Match data version.
	DataVersion string `json:"data_version,omitempty"`
	// Match id.
	MatchID string `json:"match_id,omitempty"`
	// A list of participant PUUIDs.
	Participants []string `json:"participants,omitempty"`
}

// NewDeckDto data object.
type NewDeckV1DTO struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// PlatformDataDto data object.
type PlatformDataV1DTO struct {
	ID           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Incidents    []StatusV1DTO `json:"incidents,omitempty"`
	Locales      []string      `json:"locales,omitempty"`
	Maintenances []StatusV1DTO `json:"maintenances,omitempty"`
}

// PlayerDto data object.
type PlayerV1DTO struct {
	// Code for the deck played. Refer to LOR documentation for details on deck codes.
	DeckCode    string   `json:"deck_code,omitempty"`
	DeckID      string   `json:"deck_id,omitempty"`
	GameOutcome string   `json:"game_outcome,omitempty"`
	PUUID       string   `json:"puuid,omitempty"`
	Factions    []string `json:"factions,omitempty"`
	// The order in which the players took turns.
	OrderOfPlay int32 `json:"order_of_play,omitempty"`
}

// StatusDto data object.
type StatusV1DTO struct {
	ArchiveAt string `json:"archive_at,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	// (Legal values:  info,  warning,  critical)
	IncidentSeverity string `json:"incident_severity,omitempty"`
	// (Legal values:  scheduled,  in_progress,  complete)
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	// (Legal values: windows, macos, android, ios, ps4, xbone, switch)
	Platforms []string       `json:"platforms,omitempty"`
	Titles    []ContentV1DTO `json:"titles,omitempty"`
	Updates   []UpdateV1DTO  `json:"updates,omitempty"`
	ID        int32          `json:"id,omitempty"`
}

// UpdateDto data object.
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
