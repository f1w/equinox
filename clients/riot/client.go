// This package is used to interact with all Riot endpoints.
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

// Spec version = 26952273dd9de767dc805b41d363fe8ff8cd0510

import "github.com/Kyagara/equinox/v2/internal"

type Client struct {
	AccountV1 AccountV1
}

// Creates a new Riot Client using the internal.Client provided.
func NewRiotClient(client *internal.Client) *Client {
	return &Client{
		AccountV1: AccountV1{internal: client},
	}
}
