// Automatically generated package.
package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 95a5cf31a385d91b952e19190af5a828d2e60ed8

import (
	"fmt"
	"net/http"
    
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"go.uber.org/zap"
)

// TftLeagueV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [tft-league-v1]: https://developer.riotgames.com/apis#tft-league-v1
type LeagueV1 struct {
	internalClient *internal.InternalClient
}

// Get the challenger league.
//
// # Parameters
//   - `route` - Route to query.
//   - `queue` (optional, in query) - Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getChallengerLeague]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getChallengerLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getChallengerLeague
func (e *LeagueV1) ChallengerByQueue(route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "ChallengerByQueue")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/tft/league/v1/challenger", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  values := request.URL.Query()
  if queue != "" {
    values.Set("queue", fmt.Sprint(queue))
  }
  request.URL.RawQuery = values.Encode()
  var data LeagueListV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get league entries for a given summoner ID.
//
// # Parameters
//   - `route` - Route to query.
//   - `summoner_id` (required, in path)
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueEntriesForSummoner]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getLeagueEntriesForSummoner]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueEntriesForSummoner
func (e *LeagueV1) SummonerEntries(route PlatformRoute, summonerId string) ([]LeagueEntryV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "SummonerEntries")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/league/v1/entries/by-summoner/%v", summonerId), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]LeagueEntryV1DTO), err
	}
  var data []LeagueEntryV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]LeagueEntryV1DTO), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// Get all the league entries.
//
// # Parameters
//   - `route` - Route to query.
//   - `tier` (required, in path)
//   - `division` (required, in path)
//   - `queue` (optional, in query) - Defaults to RANKED_TFT.
//   - `page` (optional, in query) - Defaults to 1. Starts with page 1.
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueEntries]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getLeagueEntries]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueEntries
func (e *LeagueV1) Entries(route PlatformRoute, tier Tier, division string, page int32, queue string) ([]LeagueEntryV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "Entries")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/league/v1/entries/%v/%v", tier, division), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]LeagueEntryV1DTO), err
	}
  values := request.URL.Query()
  if page != -1 {
    values.Set("page", fmt.Sprint(page))
  }
  if queue != "" {
    values.Set("queue", fmt.Sprint(queue))
  }
  request.URL.RawQuery = values.Encode()
  var data []LeagueEntryV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]LeagueEntryV1DTO), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// Get the grandmaster league.
//
// # Parameters
//   - `route` - Route to query.
//   - `queue` (optional, in query) - Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getGrandmasterLeague]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getGrandmasterLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getGrandmasterLeague
func (e *LeagueV1) GrandmasterByQueue(route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "GrandmasterByQueue")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/tft/league/v1/grandmaster", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  values := request.URL.Query()
  if queue != "" {
    values.Set("queue", fmt.Sprint(queue))
  }
  request.URL.RawQuery = values.Encode()
  var data LeagueListV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get league with given ID, including inactive entries.
//
// # Parameters
//   - `route` - Route to query.
//   - `league_id` (required, in path) - The UUID of the league.
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueById]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getLeagueById]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueById
func (e *LeagueV1) ByID(route PlatformRoute, leagueId string) (*LeagueListV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "ByID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/league/v1/leagues/%v", leagueId), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data LeagueListV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get the master league.
//
// # Parameters
//   - `route` - Route to query.
//   - `queue` (optional, in query) - Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getMasterLeague]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getMasterLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getMasterLeague
func (e *LeagueV1) MasterByQueue(route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "MasterByQueue")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/tft/league/v1/master", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  values := request.URL.Query()
  if queue != "" {
    values.Set("queue", fmt.Sprint(queue))
  }
  request.URL.RawQuery = values.Encode()
  var data LeagueListV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get the top rated ladder for given queue
//
// # Parameters
//   - `route` - Route to query.
//   - `queue` (required, in path)
//
// # Riot API Reference
//
// [tft-league-v1.getTopRatedLadder]
//
// Note: this method is automatically generated.
//
// [tft-league-v1.getTopRatedLadder]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getTopRatedLadder
func (e *LeagueV1) TopRatedLadder(route PlatformRoute, queue QueueType) ([]TopRatedLadderEntryV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "LeagueV1", "TopRatedLadder")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/league/v1/rated-ladders/%v/top", queue), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]TopRatedLadderEntryV1DTO), err
	}
  var data []TopRatedLadderEntryV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]TopRatedLadderEntryV1DTO), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// TftMatchV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [tft-match-v1]: https://developer.riotgames.com/apis#tft-match-v1
type MatchV1 struct {
	internalClient *internal.InternalClient
}

// Get a list of match ids by PUUID
//
// # Parameters
//   - `route` - Route to query.
//   - `puuid` (required, in path)
//   - `start` (optional, in query) - Defaults to 0. Start index.
//   - `end_time` (optional, in query) - Epoch timestamp in seconds.
//   - `start_time` (optional, in query) - Epoch timestamp in seconds. The matchlist started storing timestamps on June 16th, 2021. Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
//   - `count` (optional, in query) - Defaults to 20. Number of match ids to return.
//
// # Riot API Reference
//
// [tft-match-v1.getMatchIdsByPUUID]
//
// Note: this method is automatically generated.
//
// [tft-match-v1.getMatchIdsByPUUID]: https://developer.riotgames.com/api-methods/#tft-match-v1/GET_getMatchIdsByPUUID
func (e *MatchV1) ListByPUUID(route api.RegionalRoute, puuid string, count int32, endTime int64, start int32, startTime int64) ([]string, error) {
  logger := e.internalClient.Logger("TFT", "MatchV1", "ListByPUUID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/match/v1/matches/by-puuid/%v/ids", puuid), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]string), err
	}
  values := request.URL.Query()
  if count != -1 {
    values.Set("count", fmt.Sprint(count))
  }
  if endTime != -1 {
    values.Set("endTime", fmt.Sprint(endTime))
  }
  if start != -1 {
    values.Set("start", fmt.Sprint(start))
  }
  if startTime != -1 {
    values.Set("startTime", fmt.Sprint(startTime))
  }
  request.URL.RawQuery = values.Encode()
  var data []string
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]string), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// Get a match by match id
//
// # Parameters
//   - `route` - Route to query.
//   - `match_id` (required, in path)
//
// # Riot API Reference
//
// [tft-match-v1.getMatch]
//
// Note: this method is automatically generated.
//
// [tft-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#tft-match-v1/GET_getMatch
func (e *MatchV1) ByID(route api.RegionalRoute, matchId string) (*MatchV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "MatchV1", "ByID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/match/v1/matches/%v", matchId), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data MatchV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// TftStatusV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [tft-status-v1]: https://developer.riotgames.com/apis#tft-status-v1
type StatusV1 struct {
	internalClient *internal.InternalClient
}

// Get Teamfight Tactics status for the given platform.
//
// # Parameters
//   - `route` - Route to query.
//
// # Riot API Reference
//
// [tft-status-v1.getPlatformData]
//
// Note: this method is automatically generated.
//
// [tft-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#tft-status-v1/GET_getPlatformData
func (e *StatusV1) Platform(route PlatformRoute) (*PlatformDataV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "StatusV1", "Platform")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/tft/status/v1/platform-data", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data PlatformDataV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// TftSummonerV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [tft-summoner-v1]: https://developer.riotgames.com/apis#tft-summoner-v1
type SummonerV1 struct {
	internalClient *internal.InternalClient
}

// Get a summoner by account ID.
//
// # Parameters
//   - `route` - Route to query.
//   - `encrypted_account_id` (required, in path)
//
// # Riot API Reference
//
// [tft-summoner-v1.getByAccountId]
//
// Note: this method is automatically generated.
//
// [tft-summoner-v1.getByAccountId]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByAccountId
func (e *SummonerV1) ByAccountID(route PlatformRoute, encryptedAccountId string) (*SummonerV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "SummonerV1", "ByAccountID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/summoner/v1/summoners/by-account/%v", encryptedAccountId), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data SummonerV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get a summoner by summoner name.
//
// # Parameters
//   - `route` - Route to query.
//   - `summoner_name` (required, in path) - Summoner Name
//
// # Riot API Reference
//
// [tft-summoner-v1.getBySummonerName]
//
// Note: this method is automatically generated.
//
// [tft-summoner-v1.getBySummonerName]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getBySummonerName
func (e *SummonerV1) ByName(route PlatformRoute, summonerName string) (*SummonerV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "SummonerV1", "ByName")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/summoner/v1/summoners/by-name/%v", summonerName), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data SummonerV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get a summoner by PUUID.
//
// # Parameters
//   - `route` - Route to query.
//   - `encrypted_puuid` (required, in path) - Summoner ID
//
// # Riot API Reference
//
// [tft-summoner-v1.getByPUUID]
//
// Note: this method is automatically generated.
//
// [tft-summoner-v1.getByPUUID]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByPUUID
func (e *SummonerV1) ByPUUID(route PlatformRoute, encryptedPUUID string) (*SummonerV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "SummonerV1", "ByPUUID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/summoner/v1/summoners/by-puuid/%v", encryptedPUUID), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data SummonerV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get a summoner by access token.
//
// # Parameters
//   - `route` - Route to query.
//   - `authorization` (optional, in header) - Bearer token.
//
// # Riot API Reference
//
// [tft-summoner-v1.getByAccessToken]
//
// Note: this method is automatically generated.
//
// [tft-summoner-v1.getByAccessToken]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByAccessToken
func (e *SummonerV1) ByAccessToken(route PlatformRoute, authorization string) (*SummonerV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "SummonerV1", "ByAccessToken")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/tft/summoner/v1/summoners/me", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  if authorization == "" {
    return new(SummonerV1DTO), fmt.Errorf("'authorization' header is required")
  }
  request.Header.Set("authorization", fmt.Sprint(authorization))
  var data SummonerV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// Get a summoner by summoner ID.
//
// # Parameters
//   - `route` - Route to query.
//   - `encrypted_summoner_id` (required, in path) - Summoner ID
//
// # Riot API Reference
//
// [tft-summoner-v1.getBySummonerId]
//
// Note: this method is automatically generated.
//
// [tft-summoner-v1.getBySummonerId]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getBySummonerId
func (e *SummonerV1) BySummonerID(route PlatformRoute, encryptedSummonerId string) (*SummonerV1DTO, error) {
  logger := e.internalClient.Logger("TFT", "SummonerV1", "BySummonerID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/tft/summoner/v1/summoners/%v", encryptedSummonerId), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data SummonerV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}
