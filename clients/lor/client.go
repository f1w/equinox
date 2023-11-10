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
	"github.com/Kyagara/equinox/internal"
)

// Note: this struct is automatically generated.
type LORClient struct {
	internalClient  *internal.InternalClient
	DeckV1  *DeckV1
	InventoryV1  *InventoryV1
	MatchV1  *MatchV1
	RankedV1  *RankedV1
	StatusV1  *StatusV1
}

// Creates a new LORClient using the InternalClient provided.
func NewLORClient(client *internal.InternalClient) *LORClient {
	return &LORClient{
        internalClient: client,
        DeckV1: &DeckV1{internalClient: client},
        InventoryV1: &InventoryV1{internalClient: client},
        MatchV1: &MatchV1{internalClient: client},
        RankedV1: &RankedV1{internalClient: client},
        StatusV1: &StatusV1{internalClient: client},
	}
}
