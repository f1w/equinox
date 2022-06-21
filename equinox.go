package equinox

import (
	"fmt"
	"strings"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/lor"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/clients/tft"
	"github.com/Kyagara/equinox/clients/val"
	"github.com/Kyagara/equinox/internal"
)

type Equinox struct {
	internalClient *internal.InternalClient
	Riot           *riot.RiotClient
	LOL            *lol.LOLClient
	TFT            *tft.TFTClient
	LOR            *lor.LORClient
	VAL            *val.VALClient
}

//	Creates a new Equinox client with a default configuration
//
//		- `Cluster`    : api.AmericasCluster
//		- `LogLevel`   : api.FatalLevel
//		- `Timeout`    : 10 Seconds
//		- `Retry`      : true
//		- `TTL`        : 240 Seconds
//		- `RateLimit`  : true
func NewClient(key string) (*Equinox, error) {
	if !strings.HasPrefix(key, "RGAPI-") {
		return nil, fmt.Errorf("API Key not provided")
	}

	config := &api.EquinoxConfig{
		Key:       key,
		Cluster:   api.AmericasCluster,
		LogLevel:  api.FatalLevel,
		Timeout:   10,
		Retry:     true,
		TTL:       240,
		RateLimit: true,
	}

	client, err := internal.NewInternalClient(config)

	if err != nil {
		return nil, err
	}

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
		TFT:            tft.NewTFTClient(client),
		LOR:            lor.NewLORClient(client),
		VAL:            val.NewVALClient(client),
	}

	return equinox, nil
}

//	Creates a new Equinox client using a custom configuration.
//
//	If you don't specify a Timeout this will disable the timeout for the http.Client.
func NewClientWithConfig(config *api.EquinoxConfig) (*Equinox, error) {
	if config == nil {
		return nil, fmt.Errorf("equinox configuration not provided")
	}

	if !strings.HasPrefix(config.Key, "RGAPI-") {
		return nil, fmt.Errorf("API Key not provided")
	}

	if config.Cluster == "" {
		return nil, fmt.Errorf("cluster not provided")
	}

	client, err := internal.NewInternalClient(config)

	if err != nil {
		return nil, err
	}

	equinox := &Equinox{
		internalClient: client,
		Riot:           riot.NewRiotClient(client),
		LOL:            lol.NewLOLClient(client),
		TFT:            tft.NewTFTClient(client),
		LOR:            lor.NewLORClient(client),
		VAL:            val.NewVALClient(client),
	}

	return equinox, nil
}

// Clears the cache
func (c *Equinox) ClearCache() error {
	err := c.internalClient.ClearInternalClientCache()

	return err
}
