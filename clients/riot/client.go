// This package is used to interact with all Riot Games endpoints.
//   - AccountV1
//
// Note: this package is automatically generated.
package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 2099d6a309eb237d31b9c5fc1071c36eca85d91f

import "github.com/Kyagara/equinox/internal"

// Note: this struct is automatically generated.
type RiotClient struct {
	AccountV1 *AccountV1
}

// Creates a new RiotClient using the InternalClient provided.
func NewRiotClient(client *internal.InternalClient) *RiotClient {
	return &RiotClient{
		AccountV1: &AccountV1{internal: client},
	}
}
