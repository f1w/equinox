// This package is used to interact with all VAL endpoints.
//   - ContentV1
//   - MatchV1
//   - RankedV1
//   - StatusV1
//
// Note: this package is automatically generated.
package val

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 6461993a9c4165ddca053929f19f6d0e3eb1ca14

import "github.com/Kyagara/equinox/v2/internal"

type Client struct {
	ContentV1 ContentV1
	MatchV1   MatchV1
	RankedV1  RankedV1
	StatusV1  StatusV1
}

// Creates a new VAL Client using the internal.Client provided.
func NewVALClient(client *internal.Client) *Client {
	return &Client{
		ContentV1: ContentV1{internal: client},
		MatchV1:   MatchV1{internal: client},
		RankedV1:  RankedV1{internal: client},
		StatusV1:  StatusV1{internal: client},
	}
}
