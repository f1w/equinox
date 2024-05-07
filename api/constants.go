package api

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 92f57e3e7279cc02ec6a5ce6665ca08354d6a178

// Regional routes, used in tournament services, Legends of Runeterra (LoR), and some other endpoints.
type RegionalRoute string

const (
	// North and South America.
	AMERICAS RegionalRoute = "americas"
	// Asia-Pacific, deprecated, for some old matches in `lor-match-v1`.
	//
	// Deprecated
	APAC RegionalRoute = "apac"
	// Asia, used for LoL matches (`match-v5`) and TFT matches (`tft-match-v1`).
	ASIA RegionalRoute = "asia"
	// Special esports platform for `account-v1`. Do not confuse with the `esports` Valorant platform route.
	ESPORTS RegionalRoute = "esports"
	// Europe.
	EUROPE RegionalRoute = "europe"
	// South East Asia, used for LoR, LoL matches (`match-v5`), and TFT matches (`tft-match-v1`).
	SEA RegionalRoute = "sea"
)

func (route RegionalRoute) String() string {
	switch route {
	case AMERICAS:
		return "americas"
	case APAC:
		return "apac"
	case ASIA:
		return "asia"
	case ESPORTS:
		return "esports"
	case EUROPE:
		return "europe"
	case SEA:
		return "sea"
	default:
		return string(route)
	}
}
