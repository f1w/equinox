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

// Spec version = d712d94a43004a22ad9f31b9ebfbcaa9e0820305

import "github.com/Kyagara/equinox/internal"

type Client struct {
	AccountV1 AccountV1
}

// Creates a new Riot Client using the internal.Client provided.
func NewRiotClient(client *internal.Client) *Client {
	return &Client{
		AccountV1: AccountV1{internal: client},
	}
}
