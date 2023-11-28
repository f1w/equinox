// This package is used to interact with Community Dragon endpoints.
package cdragon

import (
	"github.com/Kyagara/equinox/internal"
)

// CDragon endpoint URLs.
const (
	VersionsURL = "/api/versions.json"
	ChampionURL = "/%s/champion/%v/data"
)

type CDragonClient struct {
	Version  *VersionEndpoint
	Champion *ChampionEndpoint
}

// Returns a new CDragonClient using the InternalClient provided.
func NewCDragonClient(client *internal.InternalClient) *CDragonClient {
	return &CDragonClient{
		Version:  &VersionEndpoint{client},
		Champion: &ChampionEndpoint{client},
	}
}
