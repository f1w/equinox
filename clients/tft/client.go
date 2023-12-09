// This package is used to interact with all Teamfight Tactics endpoints.
//   - LeagueV1
//   - MatchV1
//   - StatusV1
//   - SummonerV1
//
// Note: this package is automatically generated.
package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = cd204d7d764a025c280943766bc498278e439a6c

import "github.com/Kyagara/equinox/internal"

// Note: this struct is automatically generated.
type Client struct {
	LeagueV1   LeagueV1
	MatchV1    MatchV1
	StatusV1   StatusV1
	SummonerV1 SummonerV1
}

// Creates a new TFT Client using the internal.Client provided.
func NewTFTClient(client *internal.Client) *Client {
	return &Client{
		LeagueV1:   LeagueV1{internal: client},
		MatchV1:    MatchV1{internal: client},
		StatusV1:   StatusV1{internal: client},
		SummonerV1: SummonerV1{internal: client},
	}
}
