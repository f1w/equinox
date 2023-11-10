// Automatically generated package.
package lor

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

// LorDeckV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [lor-deck-v1]: https://developer.riotgames.com/apis#lor-deck-v1
type DeckV1 struct {
	internalClient *internal.InternalClient
}

// Get a list of the calling user's decks.
//
// # Parameters
//   - `route` - Route to query.
//   - `authorization` (required, in header)
//
// # Riot API Reference
//
// [lor-deck-v1.getDecks]
//
// Note: this method is automatically generated.
//
// [lor-deck-v1.getDecks]: https://developer.riotgames.com/api-methods/#lor-deck-v1/GET_getDecks
func (e *DeckV1) Decks(route api.RegionalRoute, authorization string) ([]DeckV1DTO, error) {
  logger := e.internalClient.Logger("LOR", "DeckV1", "Decks")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/lor/deck/v1/decks/me", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]DeckV1DTO), err
	}
  if authorization == "" {
    return *new([]DeckV1DTO), fmt.Errorf("'authorization' header is required")
  }
  request.Header.Set("authorization", fmt.Sprint(authorization))
  var data []DeckV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]DeckV1DTO), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// Create a new deck for the calling user.
//
// # Parameters
//   - `route` - Route to query.
//   - `authorization` (required, in header)
//
// # Riot API Reference
//
// [lor-deck-v1.createDeck]
//
// Note: this method is automatically generated.
//
// [lor-deck-v1.createDeck]: https://developer.riotgames.com/api-methods/#lor-deck-v1/POST_createDeck
func (e *DeckV1) CreateDeck(route api.RegionalRoute, body *NewDeckV1DTO, authorization string) (string, error) {
  logger := e.internalClient.Logger("LOR", "DeckV1", "CreateDeck")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodPost, route, "/lor/deck/v1/decks/me", body)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new(string), err
	}
  if authorization == "" {
    return *new(string), fmt.Errorf("'authorization' header is required")
  }
  request.Header.Set("authorization", fmt.Sprint(authorization))
  var data string
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new(string), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// LorInventoryV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [lor-inventory-v1]: https://developer.riotgames.com/apis#lor-inventory-v1
type InventoryV1 struct {
	internalClient *internal.InternalClient
}

// Return a list of cards owned by the calling user.
//
// # Parameters
//   - `route` - Route to query.
//   - `authorization` (required, in header)
//
// # Riot API Reference
//
// [lor-inventory-v1.getCards]
//
// Note: this method is automatically generated.
//
// [lor-inventory-v1.getCards]: https://developer.riotgames.com/api-methods/#lor-inventory-v1/GET_getCards
func (e *InventoryV1) Cards(route api.RegionalRoute, authorization string) ([]CardV1DTO, error) {
  logger := e.internalClient.Logger("LOR", "InventoryV1", "Cards")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/lor/inventory/v1/cards/me", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]CardV1DTO), err
	}
  if authorization == "" {
    return *new([]CardV1DTO), fmt.Errorf("'authorization' header is required")
  }
  request.Header.Set("authorization", fmt.Sprint(authorization))
  var data []CardV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]CardV1DTO), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// LorMatchV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [lor-match-v1]: https://developer.riotgames.com/apis#lor-match-v1
type MatchV1 struct {
	internalClient *internal.InternalClient
}

// Get a list of match ids by PUUID
//
// # Parameters
//   - `route` - Route to query.
//   - `puuid` (required, in path)
//
// # Riot API Reference
//
// [lor-match-v1.getMatchIdsByPUUID]
//
// Note: this method is automatically generated.
//
// [lor-match-v1.getMatchIdsByPUUID]: https://developer.riotgames.com/api-methods/#lor-match-v1/GET_getMatchIdsByPUUID
func (e *MatchV1) ListByPUUID(route api.RegionalRoute, puuid string) ([]string, error) {
  logger := e.internalClient.Logger("LOR", "MatchV1", "ListByPUUID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/lor/match/v1/matches/by-puuid/%v/ids", puuid), nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return *new([]string), err
	}
  var data []string
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return *new([]string), err
  }
  logger.Debug("Method executed successfully")
  return data, nil
}

// Get match by id
//
// # Parameters
//   - `route` - Route to query.
//   - `match_id` (required, in path)
//
// # Riot API Reference
//
// [lor-match-v1.getMatch]
//
// Note: this method is automatically generated.
//
// [lor-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#lor-match-v1/GET_getMatch
func (e *MatchV1) ByID(route api.RegionalRoute, matchId string) (*MatchV1DTO, error) {
  logger := e.internalClient.Logger("LOR", "MatchV1", "ByID")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, fmt.Sprintf("/lor/match/v1/matches/%v", matchId), nil)
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

// LorRankedV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [lor-ranked-v1]: https://developer.riotgames.com/apis#lor-ranked-v1
type RankedV1 struct {
	internalClient *internal.InternalClient
}

// Get the players in Master tier.
//
// # Parameters
//   - `route` - Route to query.
//
// # Riot API Reference
//
// [lor-ranked-v1.getLeaderboards]
//
// Note: this method is automatically generated.
//
// [lor-ranked-v1.getLeaderboards]: https://developer.riotgames.com/api-methods/#lor-ranked-v1/GET_getLeaderboards
func (e *RankedV1) Leaderboards(route api.RegionalRoute) (*LeaderboardV1DTO, error) {
  logger := e.internalClient.Logger("LOR", "RankedV1", "Leaderboards")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/lor/ranked/v1/leaderboards", nil)
  if err != nil {
    logger.Error("Error creating request", zap.Error(err))
    return nil, err
	}
  var data LeaderboardV1DTO
  err = e.internalClient.Execute(request, &data)
  if err != nil {
    logger.Error("Error executing request", zap.Error(err))
    return nil, err
  }
  logger.Debug("Method executed successfully")
  return &data, nil
}

// LorStatusV1 endpoints handle.
//
// Note: this struct is automatically generated.
//
// [lor-status-v1]: https://developer.riotgames.com/apis#lor-status-v1
type StatusV1 struct {
	internalClient *internal.InternalClient
}

// Get Legends of Runeterra status for the given platform.
//
// # Parameters
//   - `route` - Route to query.
//
// # Riot API Reference
//
// [lor-status-v1.getPlatformData]
//
// Note: this method is automatically generated.
//
// [lor-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#lor-status-v1/GET_getPlatformData
func (e *StatusV1) Platform(route api.RegionalRoute) (*PlatformDataV1DTO, error) {
  logger := e.internalClient.Logger("LOR", "StatusV1", "Platform")
  logger.Debug("Method started execution")
  request, err := e.internalClient.Request(api.BaseURLFormat, http.MethodGet, route, "/lor/status/v1/platform-data", nil)
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
