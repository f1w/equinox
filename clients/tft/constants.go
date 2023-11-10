package tft

///////////////////////////////////////////////
//                                           //
//                     !                     //
//   This file is automatically generated!   //
//           Do not directly edit!           //
//                                           //
///////////////////////////////////////////////

// Spec version = 95a5cf31a385d91b952e19190af5a828d2e60ed8

// TFT ranked tiers, such as gold, diamond, challenger, etc.
type Tier string

const (
	Iron        Tier = "IRON"
	Bronze      Tier = "BRONZE"
	Silver      Tier = "SILVER"
	Gold        Tier = "GOLD"
	Platinum    Tier = "PLATINUM"
	Diamond     Tier = "DIAMOND"
	Master      Tier = "MASTER"
	Grandmaster Tier = "GRANDMASTER"
	Challenger  Tier = "CHALLENGER"
)

// TFT rank divisions, I, II, III, IV, and (deprecated) V.
type Division string

const (
	I   Division = "I"
	II  Division = "II"
	III Division = "III"
	IV  Division = "IV"
    // Deprecated
    V   Division = "V"
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

// LoL ranked queue types.
type QueueType string

const (
    // Ranked Teamfight Tactics games
    //
    RankedTft QueueType = "RANKED_TFT"
    // Ranked Teamfight Tactics (Hyper Roll) games
    //
    RankedTftTurbo QueueType = "RANKED_TFT_TURBO"
    // Ranked Teamfight Tactics (Double Up Workshop) games
    //
    // Deprecated in patch 12.11 in favor of queueId 1160
    RankedTftPairs QueueType = "RANKED_TFT_PAIRS"
    // Ranked Teamfight Tactics (Double Up Workshop) games
    //
    RankedTftDoubleUp QueueType = "RANKED_TFT_DOUBLE_UP"
)
