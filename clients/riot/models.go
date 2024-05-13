package riot

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 6461993a9c4165ddca053929f19f6d0e3eb1ca14

// account-v1.ActiveShardDto
type AccountActiveShardV1DTO struct {
	ActiveShard string `json:"activeShard,omitempty"`
	Game        string `json:"game,omitempty"`
	PUUID       string `json:"puuid,omitempty"`
}

// account-v1.AccountDto
type AccountV1DTO struct {
	// This field may be excluded from the response if the account doesn't have a gameName.
	GameName string `json:"gameName,omitempty"`
	PUUID    string `json:"puuid,omitempty"`
	// This field may be excluded from the response if the account doesn't have a tagLine.
	TagLine string `json:"tagLine,omitempty"`
}
