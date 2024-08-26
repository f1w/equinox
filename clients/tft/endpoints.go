package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 3261d1c333d2269147205cdd87e62d64b898e005

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/internal"
)

// # Riot API Reference
//
// [tft-league-v1]
//
// [tft-league-v1]: https://developer.riotgames.com/apis#tft-league-v1
type LeagueV1 struct {
	internal *internal.Client
}

// Get league with given ID, including inactive entries.
//
// # Parameters
//   - route: Route to query.
//   - leagueId: The UUID of the league.
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueById]
//
// [tft-league-v1.getLeagueById]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueById
func (endpoint *LeagueV1) ByID(ctx context.Context, route PlatformRoute, leagueId string) (*LeagueListV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_ByID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/leagues/", leagueId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getLeagueById", nil)
	if err != nil {
		return nil, err
	}
	var data LeagueListV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get the challenger league.
//
// # Parameters
//   - route: Route to query.
//   - queue (optional): Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getChallengerLeague]
//
// [tft-league-v1.getChallengerLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getChallengerLeague
func (endpoint *LeagueV1) ChallengerByQueue(ctx context.Context, route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_ChallengerByQueue")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/challenger"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getChallengerLeague", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if queue != "" {
		values.Set("queue", queue)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data LeagueListV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get all the league entries.
//
// # Parameters
//   - route: Route to query.
//   - tier
//   - division
//   - queue (optional): Defaults to RANKED_TFT.
//   - page (optional): Defaults to 1. Starts with page 1.
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueEntries]
//
// [tft-league-v1.getLeagueEntries]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueEntries
func (endpoint *LeagueV1) Entries(ctx context.Context, route PlatformRoute, tier Tier, division string, queue string, page int32) ([]LeagueEntryV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_Entries")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/entries/", tier.String(), "/", division}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getLeagueEntries", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if page != -1 {
		values.Set("page", strconv.FormatInt(int64(page), 10))
	}
	if queue != "" {
		values.Set("queue", queue)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data []LeagueEntryV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Get the grandmaster league.
//
// # Parameters
//   - route: Route to query.
//   - queue (optional): Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getGrandmasterLeague]
//
// [tft-league-v1.getGrandmasterLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getGrandmasterLeague
func (endpoint *LeagueV1) GrandmasterByQueue(ctx context.Context, route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_GrandmasterByQueue")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/grandmaster"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getGrandmasterLeague", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if queue != "" {
		values.Set("queue", queue)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data LeagueListV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get the master league.
//
// # Parameters
//   - route: Route to query.
//   - queue (optional): Defaults to RANKED_TFT.
//
// # Riot API Reference
//
// [tft-league-v1.getMasterLeague]
//
// [tft-league-v1.getMasterLeague]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getMasterLeague
func (endpoint *LeagueV1) MasterByQueue(ctx context.Context, route PlatformRoute, queue string) (*LeagueListV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_MasterByQueue")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/master"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getMasterLeague", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if queue != "" {
		values.Set("queue", queue)
	}
	request.Request.URL.RawQuery = values.Encode()
	var data LeagueListV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get league entries for a given summoner ID.
//
// # Parameters
//   - route: Route to query.
//   - summonerId
//
// # Riot API Reference
//
// [tft-league-v1.getLeagueEntriesForSummoner]
//
// [tft-league-v1.getLeagueEntriesForSummoner]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getLeagueEntriesForSummoner
func (endpoint *LeagueV1) SummonerEntries(ctx context.Context, route PlatformRoute, summonerId string) ([]LeagueEntryV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_SummonerEntries")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/entries/by-summoner/", summonerId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getLeagueEntriesForSummoner", nil)
	if err != nil {
		return nil, err
	}
	var data []LeagueEntryV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Get the top rated ladder for given queue
//
// # Parameters
//   - route: Route to query.
//   - queue
//
// # Riot API Reference
//
// [tft-league-v1.getTopRatedLadder]
//
// [tft-league-v1.getTopRatedLadder]: https://developer.riotgames.com/api-methods/#tft-league-v1/GET_getTopRatedLadder
func (endpoint *LeagueV1) TopRatedLadder(ctx context.Context, route PlatformRoute, queue QueueType) ([]LeagueTopRatedLadderEntryV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_LeagueV1_TopRatedLadder")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/league/v1/rated-ladders/", queue.String(), "/top"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-league-v1.getTopRatedLadder", nil)
	if err != nil {
		return nil, err
	}
	var data []LeagueTopRatedLadderEntryV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// # Riot API Reference
//
// [tft-match-v1]
//
// [tft-match-v1]: https://developer.riotgames.com/apis#tft-match-v1
type MatchV1 struct {
	internal *internal.Client
}

// Get a match by match id
//
// # Parameters
//   - route: Route to query.
//   - matchId
//
// # Riot API Reference
//
// [tft-match-v1.getMatch]
//
// [tft-match-v1.getMatch]: https://developer.riotgames.com/api-methods/#tft-match-v1/GET_getMatch
func (endpoint *MatchV1) ByID(ctx context.Context, route api.RegionalRoute, matchId string) (*MatchV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_MatchV1_ByID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/match/v1/matches/", matchId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-match-v1.getMatch", nil)
	if err != nil {
		return nil, err
	}
	var data MatchV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get a list of match ids by PUUID
//
// # Parameters
//   - route: Route to query.
//   - puuid
//   - start (optional): Defaults to 0. Start index.
//   - endTime (optional): Epoch timestamp in seconds.
//   - startTime (optional): Epoch timestamp in seconds. The matchlist started storing timestamps on June 16th, 2021. Any matches played before June 16th, 2021 won't be included in the results if the startTime filter is set.
//   - count (optional): Defaults to 20. Number of match ids to return.
//
// # Riot API Reference
//
// [tft-match-v1.getMatchIdsByPUUID]
//
// [tft-match-v1.getMatchIdsByPUUID]: https://developer.riotgames.com/api-methods/#tft-match-v1/GET_getMatchIdsByPUUID
func (endpoint *MatchV1) ListByPUUID(ctx context.Context, route api.RegionalRoute, puuid string, start int32, endTime int64, startTime int64, count int32) ([]string, error) {
	logger := endpoint.internal.Logger("TFT_MatchV1_ListByPUUID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/match/v1/matches/by-puuid/", puuid, "/ids"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-match-v1.getMatchIdsByPUUID", nil)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	if count != -1 {
		values.Set("count", strconv.FormatInt(int64(count), 10))
	}
	if endTime != -1 {
		values.Set("endTime", strconv.FormatInt(endTime, 10))
	}
	if start != -1 {
		values.Set("start", strconv.FormatInt(int64(start), 10))
	}
	if startTime != -1 {
		values.Set("startTime", strconv.FormatInt(startTime, 10))
	}
	request.Request.URL.RawQuery = values.Encode()
	data := make([]string, 0, 20)
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// # Riot API Reference
//
// [spectator-tft-v5]
//
// [spectator-tft-v5]: https://developer.riotgames.com/apis#spectator-tft-v5
type SpectatorV5 struct {
	internal *internal.Client
}

// Get current game information for the given puuid.
//
// # Parameters
//   - route: Route to query.
//   - encryptedPUUID: The puuid of the summoner.
//
// # Riot API Reference
//
// [spectator-tft-v5.getCurrentGameInfoByPuuid]
//
// [spectator-tft-v5.getCurrentGameInfoByPuuid]: https://developer.riotgames.com/api-methods/#spectator-tft-v5/GET_getCurrentGameInfoByPuuid
func (endpoint *SpectatorV5) CurrentGameInfoByPUUID(ctx context.Context, route PlatformRoute, encryptedPUUID string) (*SpectatorCurrentGameInfoV5DTO, error) {
	logger := endpoint.internal.Logger("TFT_SpectatorV5_CurrentGameInfoByPUUID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/spectator/tft/v5/active-games/by-puuid/", encryptedPUUID}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "spectator-tft-v5.getCurrentGameInfoByPuuid", nil)
	if err != nil {
		return nil, err
	}
	var data SpectatorCurrentGameInfoV5DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get list of featured games.
//
// # Parameters
//   - route: Route to query.
//
// # Riot API Reference
//
// [spectator-tft-v5.getFeaturedGames]
//
// [spectator-tft-v5.getFeaturedGames]: https://developer.riotgames.com/api-methods/#spectator-tft-v5/GET_getFeaturedGames
func (endpoint *SpectatorV5) Featured(ctx context.Context, route PlatformRoute) (*SpectatorFeaturedGamesV5DTO, error) {
	logger := endpoint.internal.Logger("TFT_SpectatorV5_Featured")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/lol/spectator/tft/v5/featured-games"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "spectator-tft-v5.getFeaturedGames", nil)
	if err != nil {
		return nil, err
	}
	var data SpectatorFeaturedGamesV5DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [tft-status-v1]
//
// [tft-status-v1]: https://developer.riotgames.com/apis#tft-status-v1
type StatusV1 struct {
	internal *internal.Client
}

// Get Teamfight Tactics status for the given platform.
//
// # Parameters
//   - route: Route to query.
//
// # Riot API Reference
//
// [tft-status-v1.getPlatformData]
//
// [tft-status-v1.getPlatformData]: https://developer.riotgames.com/api-methods/#tft-status-v1/GET_getPlatformData
func (endpoint *StatusV1) Platform(ctx context.Context, route PlatformRoute) (*StatusPlatformDataV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_StatusV1_Platform")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/status/v1/platform-data"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-status-v1.getPlatformData", nil)
	if err != nil {
		return nil, err
	}
	var data StatusPlatformDataV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// # Riot API Reference
//
// [tft-summoner-v1]
//
// [tft-summoner-v1]: https://developer.riotgames.com/apis#tft-summoner-v1
type SummonerV1 struct {
	internal *internal.Client
}

// Get a summoner by access token.
//
// # Parameters
//   - route: Route to query.
//   - accessToken: RSO access token.
//
// # Riot API Reference
//
// [tft-summoner-v1.getByAccessToken]
//
// [tft-summoner-v1.getByAccessToken]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByAccessToken
func (endpoint *SummonerV1) ByAccessToken(ctx context.Context, route PlatformRoute, accessToken string) (*SummonerV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_SummonerV1_ByAccessToken")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/summoner/v1/summoners/me"}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-summoner-v1.getByAccessToken", nil)
	if err != nil {
		return nil, err
	}
	request.Request.Header = request.Request.Header.Clone()
	request.Request.Header.Del("X-Riot-Token")
	headerValue := []string{"Bearer ", accessToken}
	request.Request.Header.Set("Authorization", strings.Join(headerValue, ""))
	var data SummonerV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get a summoner by account ID.
//
// # Parameters
//   - route: Route to query.
//   - encryptedAccountId
//
// # Riot API Reference
//
// [tft-summoner-v1.getByAccountId]
//
// [tft-summoner-v1.getByAccountId]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByAccountId
func (endpoint *SummonerV1) ByAccountID(ctx context.Context, route PlatformRoute, encryptedAccountId string) (*SummonerV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_SummonerV1_ByAccountID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/summoner/v1/summoners/by-account/", encryptedAccountId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-summoner-v1.getByAccountId", nil)
	if err != nil {
		return nil, err
	}
	var data SummonerV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get a summoner by PUUID.
//
// # Parameters
//   - route: Route to query.
//   - encryptedPUUID: Summoner ID
//
// # Riot API Reference
//
// [tft-summoner-v1.getByPUUID]
//
// [tft-summoner-v1.getByPUUID]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getByPUUID
func (endpoint *SummonerV1) ByPUUID(ctx context.Context, route PlatformRoute, encryptedPUUID string) (*SummonerV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_SummonerV1_ByPUUID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/summoner/v1/summoners/by-puuid/", encryptedPUUID}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-summoner-v1.getByPUUID", nil)
	if err != nil {
		return nil, err
	}
	var data SummonerV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// Get a summoner by summoner ID.
//
// # Parameters
//   - route: Route to query.
//   - encryptedSummonerId: Summoner ID
//
// # Riot API Reference
//
// [tft-summoner-v1.getBySummonerId]
//
// [tft-summoner-v1.getBySummonerId]: https://developer.riotgames.com/api-methods/#tft-summoner-v1/GET_getBySummonerId
func (endpoint *SummonerV1) BySummonerID(ctx context.Context, route PlatformRoute, encryptedSummonerId string) (*SummonerV1DTO, error) {
	logger := endpoint.internal.Logger("TFT_SummonerV1_BySummonerID")
	urlComponents := []string{"https://", route.String(), api.RIOT_API_BASE_URL_FORMAT, "/tft/summoner/v1/summoners/", encryptedSummonerId}
	request, err := endpoint.internal.Request(ctx, logger, http.MethodGet, urlComponents, "tft-summoner-v1.getBySummonerId", nil)
	if err != nil {
		return nil, err
	}
	var data SummonerV1DTO
	err = endpoint.internal.Execute(ctx, request, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
