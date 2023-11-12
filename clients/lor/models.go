package lor

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = ed83574d1b85ef4c52f267ee5558e3c1c3ffb412

// DeckDTO data object.
type DeckV1DTO struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Code string `json:"code"`
}

// NewDeckDTO data object.
type NewDeckV1DTO struct {
    Name string `json:"name"`
    Code string `json:"code"`
}

// CardDTO data object.
type CardV1DTO struct {
    Code string `json:"code"`
    Count string `json:"count"`
}

// MatchDTO data object.
type MatchV1DTO struct {
    // Match metadata.
    Metadata MetadataV1DTO `json:"metadata"`
    // Match info.
    Info InfoV1DTO `json:"info"`
}

// MetadataDTO data object.
type MetadataV1DTO struct {
    // Match data version.
    Data_version string `json:"data_version"`
    // Match id.
    Match_id string `json:"match_id"`
    // A list of participant PUUIDs.
    Participants []string `json:"participants"`
}

// InfoDTO data object.
type InfoV1DTO struct {
    // (Legal values:  Constructed,  Expeditions,  Tutorial)
    Game_mode string `json:"game_mode"`
    // (Legal values:  Ranked,  Normal,  AI,  Tutorial,  VanillaTrial,  Singleton,  StandardGauntlet)
    Game_type string `json:"game_type"`
    Game_start_time_utc string `json:"game_start_time_utc"`
    Game_version string `json:"game_version"`
    Players []PlayerV1DTO `json:"players"`
    // Total turns taken by both players.
    Total_turn_count int32 `json:"total_turn_count"`
}

// PlayerDTO data object.
type PlayerV1DTO struct {
    PUUID string `json:"puuid"`
    Deck_id string `json:"deck_id"`
    // Code for the deck played. Refer to LOR documentation for details on deck codes.
    Deck_code string `json:"deck_code"`
    Factions []string `json:"factions"`
    Game_outcome string `json:"game_outcome"`
    // The order in which the players took turns.
    Order_of_play int32 `json:"order_of_play"`
}

// LeaderboardDTO data object.
type LeaderboardV1DTO struct {
    // A list of players in Master tier.
    Players []PlayerV1DTO `json:"players"`
}

// PlayerDTO data object.
type LeaderboardPlayerV1DTO struct {
    Name string `json:"name"`
    Rank int32 `json:"rank"`
    // League points.
    Lp int32 `json:"lp"`
}

// PlatformDataDTO data object.
type PlatformDataV1DTO struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Locales []string `json:"locales"`
    Maintenances []StatusV1DTO `json:"maintenances"`
    Incidents []StatusV1DTO `json:"incidents"`
}

// StatusDTO data object.
type StatusV1DTO struct {
    Id int32 `json:"id"`
    // (Legal values:  scheduled,  in_progress,  complete)
    Maintenance_status string `json:"maintenance_status"`
    // (Legal values:  info,  warning,  critical)
    Incident_severity string `json:"incident_severity"`
    Titles []ContentV1DTO `json:"titles"`
    Updates []UpdateV1DTO `json:"updates"`
    Created_at string `json:"created_at"`
    Archive_at string `json:"archive_at"`
    Updated_at string `json:"updated_at"`
    // (Legal values: windows, macos, android, ios, ps4, xbone, switch)
    Platforms []string `json:"platforms"`
}

// ContentDTO data object.
type ContentV1DTO struct {
    Locale string `json:"locale"`
    Content string `json:"content"`
}

// UpdateDTO data object.
type UpdateV1DTO struct {
    Id int32 `json:"id"`
    Author string `json:"author"`
    Publish bool `json:"publish"`
    // (Legal values: riotclient, riotstatus, game)
    Publish_locations []string `json:"publish_locations"`
    Translations []ContentV1DTO `json:"translations"`
    Created_at string `json:"created_at"`
    Updated_at string `json:"updated_at"`
}