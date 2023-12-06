package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 2099d6a309eb237d31b9c5fc1071c36eca85d91f

// LoL and TFT ranked tiers, such as gold, diamond, challenger, etc.
type Tier string

const (
	IRON        Tier = "IRON"
	BRONZE      Tier = "BRONZE"
	SILVER      Tier = "SILVER"
	GOLD        Tier = "GOLD"
	PLATINUM    Tier = "PLATINUM"
	EMERALD     Tier = "EMERALD"
	DIAMOND     Tier = "DIAMOND"
	MASTER      Tier = "MASTER"
	GRANDMASTER Tier = "GRANDMASTER"
	CHALLENGER  Tier = "CHALLENGER"
)

// LoL and TFT rank divisions, I, II, III, IV, and (deprecated) V.
type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
	// Deprecated
	V Division = "V"
)

// Platform routes for League of Legends (LoL), Teamfight Tactics (TFT).
type PlatformRoute string

const (
	// Brazil.
	BR1 PlatformRoute = "br1"
	// Europe, Northeast.
	EUN1 PlatformRoute = "eun1"
	// Europe, West.
	EUW1 PlatformRoute = "euw1"
	// Japan.
	JP1 PlatformRoute = "jp1"
	// Korea.
	KR PlatformRoute = "kr"
	// Latin America, North.
	LA1 PlatformRoute = "la1"
	// Latin America, South.
	LA2 PlatformRoute = "la2"
	// North America.
	NA1 PlatformRoute = "na1"
	// Oceana.
	OC1 PlatformRoute = "oc1"
	// Philippines
	PH2 PlatformRoute = "ph2"
	// Russia
	RU PlatformRoute = "ru"
	// Singapore
	SG2 PlatformRoute = "sg2"
	// Thailand
	TH2 PlatformRoute = "th2"
	// Turkey
	TR1 PlatformRoute = "tr1"
	// Taiwan
	TW2 PlatformRoute = "tw2"
	// Vietnam
	VN2 PlatformRoute = "vn2"
	// Public Beta Environment, special beta testing platform. Located in North America.
	PBE1 PlatformRoute = "pbe1"
)

// TFT ranked queue types.
type QueueType string

const (
	// Ranked Teamfight Tactics games
	RANKED_TFT QueueType = "RANKED_TFT"
	// Ranked Teamfight Tactics (Hyper Roll) games
	RANKED_TFT_TURBO QueueType = "RANKED_TFT_TURBO"
	// Ranked Teamfight Tactics (Double Up Workshop) games
	//
	// Deprecated: Deprecated in patch 12.11 in favor of queueId 1160
	RANKED_TFT_PAIRS QueueType = "RANKED_TFT_PAIRS"
	// Ranked Teamfight Tactics (Double Up Workshop) games
	RANKED_TFT_DOUBLE_UP QueueType = "RANKED_TFT_DOUBLE_UP"
)
