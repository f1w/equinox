package internal_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestNewCache(t *testing.T) {
	cache, err := internal.NewCache(0)

	require.Nil(t, err, "expecting nil error")

	require.NotNil(t, cache, "expecting non-nil Cache")
}

func TestCache(t *testing.T) {
	cache, err := internal.NewCache(120)

	require.Nil(t, err, "expecting nil error")

	cache.Set("/", make([]byte, 1))

	res := cache.Get("/")

	require.Nil(t, err, "expecting nil err")

	require.NotNil(t, res, "expecting non-nil res")

	cache.Clear()

	res = cache.Get("/")

	require.Nil(t, err, "expecting nil err")

	assert.Nil(t, res, "expecting nil res")
}

func TestCacheHit(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.TTL = 240

	internalClient, err := internal.NewInternalClient(config)

	require.Nil(t, err, "expecting nil error")

	client := lol.NewLOLClient(internalClient)

	j := []byte("{\"freeChampionIds\":[2,12,27,29,32,35,42,54,72,78,84,98,136,164,223,777],\"freeChampionIdsForNewPlayers\":[222,254,427,82,131,147,54,17,18,37],\"maxNewPlayerLevel\":10}")

	data := &lol.ChampionRotationsDTO{}

	err = json.Unmarshal(j, data)

	require.Nil(t, err, "expecting nil error")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "br1")).
		Get(lol.ChampionURL).
		Reply(200).JSON(data)

	gock.New(fmt.Sprintf(api.BaseURLFormat, "br1")).
		Get(lol.ChampionURL).
		Reply(200).JSON(data)

	gotData, gotErr := client.Champion.Rotations(lol.BR1)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	require.Equal(t, data, gotData, fmt.Sprintf("want data %v, got %v", data, gotData))

	gotCache, gotErr := client.Champion.Rotations(lol.BR1)

	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	require.Equal(t, data, gotCache, fmt.Sprintf("want data %v, got %v", data, gotCache))
}
