package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = ed83574d1b85ef4c52f267ee5558e3c1c3ffb412

// ContentDTO data object.
type ContentV1DTO struct {
    Version string `json:"version"`
    Characters []ContentItemV1DTO `json:"characters"`
    Maps []ContentItemV1DTO `json:"maps"`
    Chromas []ContentItemV1DTO `json:"chromas"`
    Skins []ContentItemV1DTO `json:"skins"`
    SkinLevels []ContentItemV1DTO `json:"skinLevels"`
    Equips []ContentItemV1DTO `json:"equips"`
    GameModes []ContentItemV1DTO `json:"gameModes"`
    Sprays []ContentItemV1DTO `json:"sprays"`
    SprayLevels []ContentItemV1DTO `json:"sprayLevels"`
    Charms []ContentItemV1DTO `json:"charms"`
    CharmLevels []ContentItemV1DTO `json:"charmLevels"`
    PlayerCards []ContentItemV1DTO `json:"playerCards"`
    PlayerTitles []ContentItemV1DTO `json:"playerTitles"`
    Acts []ActV1DTO `json:"acts"`
    Ceremonies []ContentItemV1DTO `json:"ceremonies"`
    // Unknown type, this is a placeholder subject to change.
    Totems []string `json:"totems"`
}

// ContentItemDTO data object.
type ContentItemV1DTO struct {
    Name string `json:"name"`
    // This field is excluded from the response when a locale is set
    LocalizedNames LocalizedNamesV1DTO `json:"localizedNames"`
    Id string `json:"id"`
    AssetName string `json:"assetName"`
    // This field is only included for maps and game modes. These values are used in the match response.
    AssetPath string `json:"assetPath"`
}

// LocalizedNamesDTO data object.
type LocalizedNamesV1DTO struct {
    ArAE string `json:"ar-AE"`
    DeDE string `json:"de-DE"`
    EnGB string `json:"en-GB"`
    EnUS string `json:"en-US"`
    EsES string `json:"es-ES"`
    EsMX string `json:"es-MX"`
    FrFR string `json:"fr-FR"`
    IdID string `json:"id-ID"`
    ItIT string `json:"it-IT"`
    JaJP string `json:"ja-JP"`
    KoKR string `json:"ko-KR"`
    PlPL string `json:"pl-PL"`
    PtBR string `json:"pt-BR"`
    RuRU string `json:"ru-RU"`
    ThTH string `json:"th-TH"`
    TrTR string `json:"tr-TR"`
    ViVN string `json:"vi-VN"`
    ZhCN string `json:"zh-CN"`
    ZhTW string `json:"zh-TW"`
}

// ActDTO data object.
type ActV1DTO struct {
    Name string `json:"name"`
    // This field is excluded from the response when a locale is set
    LocalizedNames LocalizedNamesV1DTO `json:"localizedNames"`
    Id string `json:"id"`
    IsActive bool `json:"isActive"`
    ParentID string `json:"parentId"`
    Type_ string `json:"type"`
}

// MatchDTO data object.
type MatchV1DTO struct {
    MatchInfo MatchInfoV1DTO `json:"matchInfo"`
    Players []PlayerV1DTO `json:"players"`
    Coaches []CoachV1DTO `json:"coaches"`
    Teams []TeamV1DTO `json:"teams"`
    RoundResults []RoundResultV1DTO `json:"roundResults"`
}

// MatchInfoDTO data object.
type MatchInfoV1DTO struct {
    MatchID string `json:"matchId"`
    MapID string `json:"mapId"`
    GameLengthMillis int32 `json:"gameLengthMillis"`
    GameStartMillis int64 `json:"gameStartMillis"`
    ProvisioningFlowID string `json:"provisioningFlowId"`
    IsCompleted bool `json:"isCompleted"`
    CustomGameName string `json:"customGameName"`
    QueueID string `json:"queueId"`
    GameMode string `json:"gameMode"`
    IsRanked bool `json:"isRanked"`
    SeasonID string `json:"seasonId"`
}

// PlayerDTO data object.
type PlayerV1DTO struct {
    PUUID string `json:"puuid"`
    GameName string `json:"gameName"`
    TagLine string `json:"tagLine"`
    TeamID string `json:"teamId"`
    PartyID string `json:"partyId"`
    CharacterID string `json:"characterId"`
    Stats PlayerStatsV1DTO `json:"stats"`
    CompetitiveTier int32 `json:"competitiveTier"`
    PlayerCard string `json:"playerCard"`
    PlayerTitle string `json:"playerTitle"`
}

// PlayerStatsDTO data object.
type PlayerStatsV1DTO struct {
    Score int32 `json:"score"`
    RoundsPlayed int32 `json:"roundsPlayed"`
    Kills int32 `json:"kills"`
    Deaths int32 `json:"deaths"`
    Assists int32 `json:"assists"`
    PlaytimeMillis int32 `json:"playtimeMillis"`
    AbilityCasts AbilityCastsV1DTO `json:"abilityCasts"`
}

// AbilityCastsDTO data object.
type AbilityCastsV1DTO struct {
    GrenadeCasts int32 `json:"grenadeCasts"`
    Ability1Casts int32 `json:"ability1Casts"`
    Ability2Casts int32 `json:"ability2Casts"`
    UltimateCasts int32 `json:"ultimateCasts"`
}

// CoachDTO data object.
type CoachV1DTO struct {
    PUUID string `json:"puuid"`
    TeamID string `json:"teamId"`
}

// TeamDTO data object.
type TeamV1DTO struct {
    // This is an arbitrary string. Red and Blue in bomb modes. The puuid of the player in deathmatch.
    TeamID string `json:"teamId"`
    Won bool `json:"won"`
    RoundsPlayed int32 `json:"roundsPlayed"`
    RoundsWon int32 `json:"roundsWon"`
    // Team points scored. Number of kills in deathmatch.
    NumPoints int32 `json:"numPoints"`
}

// RoundResultDTO data object.
type RoundResultV1DTO struct {
    RoundNum int32 `json:"roundNum"`
    RoundResult string `json:"roundResult"`
    RoundCeremony string `json:"roundCeremony"`
    WinningTeam string `json:"winningTeam"`
    // PUUID of player
    BombPlanter string `json:"bombPlanter"`
    // PUUID of player
    BombDefuser string `json:"bombDefuser"`
    PlantRoundTime int32 `json:"plantRoundTime"`
    PlantPlayerLocations []PlayerLocationsV1DTO `json:"plantPlayerLocations"`
    PlantLocation LocationV1DTO `json:"plantLocation"`
    PlantSite string `json:"plantSite"`
    DefuseRoundTime int32 `json:"defuseRoundTime"`
    DefusePlayerLocations []PlayerLocationsV1DTO `json:"defusePlayerLocations"`
    DefuseLocation LocationV1DTO `json:"defuseLocation"`
    PlayerStats []PlayerRoundStatsV1DTO `json:"playerStats"`
    RoundResultCode string `json:"roundResultCode"`
}

// PlayerLocationsDTO data object.
type PlayerLocationsV1DTO struct {
    PUUID string `json:"puuid"`
    ViewRadians float32 `json:"viewRadians"`
    Location LocationV1DTO `json:"location"`
}

// LocationDTO data object.
type LocationV1DTO struct {
    X int32 `json:"x"`
    Y int32 `json:"y"`
}

// PlayerRoundStatsDTO data object.
type PlayerRoundStatsV1DTO struct {
    PUUID string `json:"puuid"`
    Kills []KillV1DTO `json:"kills"`
    Damage []DamageV1DTO `json:"damage"`
    Score int32 `json:"score"`
    Economy EconomyV1DTO `json:"economy"`
    Ability AbilityV1DTO `json:"ability"`
}

// KillDTO data object.
type KillV1DTO struct {
    TimeSinceGameStartMillis int32 `json:"timeSinceGameStartMillis"`
    TimeSinceRoundStartMillis int32 `json:"timeSinceRoundStartMillis"`
    // PUUID
    Killer string `json:"killer"`
    // PUUID
    Victim string `json:"victim"`
    VictimLocation LocationV1DTO `json:"victimLocation"`
    // List of PUUIDs
    Assistants []string `json:"assistants"`
    PlayerLocations []PlayerLocationsV1DTO `json:"playerLocations"`
    FinishingDamage FinishingDamageV1DTO `json:"finishingDamage"`
}

// FinishingDamageDTO data object.
type FinishingDamageV1DTO struct {
    DamageType string `json:"damageType"`
    DamageItem string `json:"damageItem"`
    IsSecondaryFireMode bool `json:"isSecondaryFireMode"`
}

// DamageDTO data object.
type DamageV1DTO struct {
    // PUUID
    Receiver string `json:"receiver"`
    Damage int32 `json:"damage"`
    Legshots int32 `json:"legshots"`
    Bodyshots int32 `json:"bodyshots"`
    Headshots int32 `json:"headshots"`
}

// EconomyDTO data object.
type EconomyV1DTO struct {
    LoadoutValue int32 `json:"loadoutValue"`
    Weapon string `json:"weapon"`
    Armor string `json:"armor"`
    Remaining int32 `json:"remaining"`
    Spent int32 `json:"spent"`
}

// AbilityDTO data object.
type AbilityV1DTO struct {
    GrenadeEffects string `json:"grenadeEffects"`
    Ability1Effects string `json:"ability1Effects"`
    Ability2Effects string `json:"ability2Effects"`
    UltimateEffects string `json:"ultimateEffects"`
}

// MatchlistDTO data object.
type MatchlistV1DTO struct {
    PUUID string `json:"puuid"`
    History []MatchlistEntryV1DTO `json:"history"`
}

// MatchlistEntryDTO data object.
type MatchlistEntryV1DTO struct {
    MatchID string `json:"matchId"`
    GameStartTimeMillis int64 `json:"gameStartTimeMillis"`
    QueueID string `json:"queueId"`
}

// RecentMatchesDTO data object.
type RecentMatchesV1DTO struct {
    CurrentTime int64 `json:"currentTime"`
    // A list of recent match ids.
    MatchIDs []string `json:"matchIds"`
}

// LeaderboardDTO data object.
type LeaderboardV1DTO struct {
    // The shard for the given leaderboard.
    Shard string `json:"shard"`
    // The act id for the given leaderboard. Act ids can be found using the val-content API.
    ActID string `json:"actId"`
    // The total number of players in the leaderboard.
    TotalPlayers int64 `json:"totalPlayers"`
    Players []PlayerV1DTO `json:"players"`
    ImmortalStartingPage int64 `json:"immortalStartingPage"`
    ImmortalStartingIndex int64 `json:"immortalStartingIndex"`
    TopTierRRThreshold int64 `json:"topTierRRThreshold"`
    TierDetails map[int64]TierDetailV1DTO `json:"tierDetails"`
    StartIndex int64 `json:"startIndex"`
    Query string `json:"query"`
}

// PlayerDTO data object.
type MatchPlayerV1DTO struct {
    // This field may be omitted if the player has been anonymized.
    PUUID string `json:"puuid"`
    // This field may be omitted if the player has been anonymized.
    GameName string `json:"gameName"`
    // This field may be omitted if the player has been anonymized.
    TagLine string `json:"tagLine"`
    LeaderboardRank int64 `json:"leaderboardRank"`
    RankedRating int64 `json:"rankedRating"`
    NumberOfWins int64 `json:"numberOfWins"`
    CompetitiveTier int64 `json:"competitiveTier"`
}

// TierDetailDTO data object.
type TierDetailV1DTO struct {
    RankedRatingThreshold int64 `json:"rankedRatingThreshold"`
    StartingPage int64 `json:"startingPage"`
    StartingIndex int64 `json:"startingIndex"`
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
type StatusContentV1DTO struct {
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