package equinox_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/clients/lol"
	"github.com/Kyagara/equinox/clients/riot"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestNewEquinoxClient(t *testing.T) {
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		key     string
	}{
		{
			name: "success",
			want: &equinox.Equinox{},
			key:  "RGAPI-TEST",
		},
		{
			name: "nil key",
			key:  "",
			want: &equinox.Equinox{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClient(test.key)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))
			require.NotEmpty(t, client, "expecting non-nil Client")
			require.NotEmpty(t, client.Cache, "expecting non-nil Client")
			require.NotEmpty(t, client.LOL, "expecting nil Client")
			require.NotEmpty(t, client.LOR, "expecting nil Client")
			require.NotEmpty(t, client.TFT, "expecting nil Client")
			require.NotEmpty(t, client.VAL, "expecting nil Client")
			require.NotEmpty(t, client.Riot, "expecting nil Client")
			require.NotEmpty(t, client.CDragon, "expecting non-nil Client")
			require.NotEmpty(t, client.DDragon, "expecting non-nil Client")
		})
	}
}

func TestNewEquinoxClientWithConfig(t *testing.T) {
	emptyKeyConfig := equinox.NewTestEquinoxConfig()
	emptyKeyConfig.Key = ""
	tests := []struct {
		name    string
		want    *equinox.Equinox
		wantErr error
		config  *api.EquinoxConfig
	}{
		{
			name:   "success",
			want:   &equinox.Equinox{},
			config: equinox.NewTestEquinoxConfig(),
		},
		{
			name:    "nil config",
			wantErr: fmt.Errorf("equinox configuration not provided"),
			config:  nil,
		},
		{
			name:    "no cache",
			want:    &equinox.Equinox{},
			config:  equinox.NewTestEquinoxConfig(),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client, err := equinox.NewClientWithConfig(test.config)
			require.Equal(t, test.wantErr, err, fmt.Sprintf("want err %v, got %v", test.wantErr, err))

			if test.wantErr == nil {
				require.NotEmpty(t, client, "expecting non-nil client")
			}

			if test.name == "no cache" {
				require.Equal(t, client.Cache.TTL, time.Duration(0), "expecting cache disabled")
			}
		})
	}
}

func TestEquinoxClientClearCache(t *testing.T) {
	config, err := equinox.DefaultConfig("RGAPI-TEST")
	config.LogLevel = api.NOP_LOG_LEVEL
	config.Retry = 0
	require.Equal(t, nil, err, fmt.Sprintf("want err %v, got %v", nil, err))

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(t, err, "expecting nil error")
	require.Equal(t, client.Cache.TTL, 4*time.Minute, "expecting cache enabled")

	account := &riot.AccountV1DTO{
		PUUID:    "puuid",
		GameName: "gamename",
		TagLine:  "tagline",
	}

	delay := 2 * time.Second

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(200).
		JSON(account).Delay(delay)

	gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, api.AMERICAS)).
		Get("/riot/account/v1/accounts/by-puuid/puuid").
		Reply(404).
		JSON(account).Delay(delay)

	gotData, gotErr := client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	require.Equal(t, account, gotData)
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start := time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration := int(time.Since(start).Seconds())

	require.Equal(t, account, gotData, fmt.Sprintf("want data %v, got %v", account, gotData))
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	if duration >= 2 {
		t.Error(fmt.Errorf("request took more than 1s, took %ds, request not cached", duration))
	}

	gotErr = client.Cache.Clear()
	require.Equal(t, nil, gotErr, fmt.Sprintf("want err %v, got %v", nil, gotErr))

	start = time.Now()
	gotData, gotErr = client.Riot.AccountV1.ByPUUID(api.AMERICAS, "puuid")
	duration = int(time.Since(start).Seconds())

	if duration <= 1 {
		t.Error(fmt.Errorf("request took less than 1s, took %ds, cache not cleared", duration))
	}

	require.Nil(t, gotData)
	require.Equal(t, api.ErrNotFound, gotErr, fmt.Sprintf("want err %v, got %v", api.ErrNotFound, gotErr))
}

func TestRateLimitWithMock(t *testing.T) {
	headers := map[string]string{
		ratelimit.APP_RATE_LIMIT_HEADER:          "50:5",
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER:    "1:5",
		ratelimit.METHOD_RATE_LIMIT_HEADER:       "50:5",
		ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: "1:5",
	}

	for i := 1; i <= 50; i++ {
		gock.New(fmt.Sprintf(api.RIOT_API_BASE_URL_FORMAT, lol.BR1)).
			Get("/lol/summoner/v4/summoners/by-puuid/puuid").
			Reply(200).
			SetHeaders(headers).
			JSON(&lol.SummonerV4DTO{})
		headers[ratelimit.APP_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:5", i)
		headers[ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER] = fmt.Sprintf("%d:5", i)
	}

	config := equinox.NewTestEquinoxConfig()
	config.LogLevel = api.WARN_LOG_LEVEL
	config.Retry = 1

	client, err := equinox.NewClientWithConfig(config)
	require.Nil(t, err)

	for i := 0; i < 50; i++ {
		_, err := client.LOL.SummonerV4.ByPUUID(lol.BR1, "puuid")
		if i < 50 {
			require.Nil(t, err)
		} else {
			require.Equal(t, fmt.Errorf("app rate limit reached on 'br1' route for method 'summoner-v4.getByPUUID'"), err)
		}
	}
}
