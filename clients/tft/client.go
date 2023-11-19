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

// Spec version = b554b42a14de337810d5a510d533453eaf6de207

import "github.com/Kyagara/equinox/internal"

// Note: this struct is automatically generated.
type TFTClient struct {
	internalClient *internal.InternalClient
	LeagueV1       *LeagueV1
	MatchV1        *MatchV1
	StatusV1       *StatusV1
	SummonerV1     *SummonerV1
}

// Creates a new TFTClient using the InternalClient provided.
func NewTFTClient(client *internal.InternalClient) *TFTClient {
	return &TFTClient{
		internalClient: client,
		LeagueV1:       &LeagueV1{internalClient: client},
		MatchV1:        &MatchV1{internalClient: client},
		StatusV1:       &StatusV1{internalClient: client},
		SummonerV1:     &SummonerV1{internalClient: client},
	}
}
