package internal_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/lol"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

func TestInternalClient(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	assert.NotNil(t, client, "expecting non-nil InternalClient")
}

func TestInternalClientPut(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Put("/").
		Reply(200)

	err := client.Put("tests", "/", nil, "", "")

	assert.Nil(t, err, "expecting nil error")
}

func TestInternalClientRetries(t *testing.T) {
	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(429).SetHeader("Retry-After", "1").
		JSON(&api.PlatformDataDTO{})

	gock.New(fmt.Sprintf(api.BaseURLFormat, lol.BR1)).
		Get(lol.StatusURL).
		Reply(200).
		JSON(&api.PlatformDataDTO{})

	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	client := internal.NewInternalClient(config)

	res := api.PlatformDataDTO{}

	// This will take 1 second.
	err := client.Get(lol.BR1, lol.StatusURL, &res, "", "", "")

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, res, "expecting non-nil response")
}

func TestInternalClientFailingRetry(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	client := internal.NewInternalClient(config)

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429).SetHeader("Retry-After", "1")

	var object api.PlainTextResponse

	// This will take 2 seconds.
	gotErr := client.Get("tests", "/", &object, "", "", "")

	wantErr := fmt.Errorf("retried 2 times, stopping")

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

func TestInternalClientRetryHeader(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.Retry = true

	client := internal.NewInternalClient(config)

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(429)

	var object api.PlainTextResponse

	gotErr := client.Get("tests", "/", &object, "", "", "")

	wantErr := fmt.Errorf("rate limited but no Retry-After header was found, stopping")

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

// Testing if InternalClient.Post() can properly decode a plain text response.
func TestInternalClientPlainTextResponse(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Post("/").
		Reply(200).BodyString("response")

	var object api.PlainTextResponse

	err := client.Post("tests", "/", nil, &object, "", "", "")

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, object, "expecting non-nil response")
}

func TestInternalClientPost(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Post("/").
		Reply(200).JSON(&api.PlatformDataDTO{})

	var object *api.PlatformDataDTO

	err := client.Post("tests", "/", nil, &object, "", "", "")

	require.Nil(t, err, "expecting nil error")

	assert.NotNil(t, object, "expecting non-nil response")
}

// Testing if the InternalClient can properly handle a status code not specified in the Riot API.
func TestInternalClientHandleErrorResponse(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Get("/").
		Reply(418).BodyString("response")

	var object api.PlainTextResponse

	gotErr := client.Get("tests", "/", &object, "", "", "")

	wantErr := api.ErrorResponse{
		Status: api.Status{
			Message:    "Unknown error",
			StatusCode: 418,
		},
	}

	require.Equal(t, wantErr, gotErr, fmt.Sprintf("want err %v, got %v", wantErr, gotErr))
}

func TestInternalClientNewRequest(t *testing.T) {
	client := internal.NewInternalClient(internal.NewTestEquinoxConfig())

	tests := []struct {
		name    string
		want    *http.Request
		wantErr error
		method  string
		url     string
	}{

		{
			name: "invalid url",
			wantErr: &url.Error{
				Op:  "parse",
				URL: "https://tests.api.riotgames.com\\:invalid:/=",
				Err: url.InvalidHostError("\\"),
			},
			url: "\\:invalid:/=",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotErr := client.Get("tests", test.url, nil, "", "", "")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientErrorResponses(t *testing.T) {
	tests := []struct {
		name    string
		wantErr error
		setup   func()
		region  string
	}{
		{
			name:    "not found",
			wantErr: api.NotFoundError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(404)
			},
			region: "tests",
		},
		{
			name:    "rate limited with retry disabled",
			wantErr: api.RateLimitedError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(429)
			},
			region: "tests",
		},
		{
			name:    "bad request",
			wantErr: api.BadRequestError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(400)
			},
			region: "tests",
		},
		{
			name:    "unauthorized",
			wantErr: api.UnauthorizedError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(401)
			},
			region: "tests",
		},
		{
			name:    "forbidden",
			wantErr: api.ForbiddenError,
			setup: func() {
				gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
					Get("/").
					Reply(403)
			},
			region: "tests",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup()
			client := internal.NewInternalClient(&api.EquinoxConfig{
				Key:       "RIOT_API_KEY",
				Cluster:   api.AmericasCluster,
				LogLevel:  api.DebugLevel,
				TTL:       0,
				Timeout:   10,
				Retry:     false,
				RateLimit: false,
			})

			var gotData api.PlainTextResponse

			gotErr := client.Get(test.region, "/", &gotData, "", "", "")

			require.Equal(t, test.wantErr, gotErr, fmt.Sprintf("want err %v, got %v", test.wantErr, gotErr))
		})
	}
}

func TestInternalClientRateLimit(t *testing.T) {
	config := internal.NewTestEquinoxConfig()

	config.RateLimit = true

	client := internal.NewInternalClient(config)

	headers := map[string]string{}

	headers["X-App-Rate-Limit"] = "1:10,1:600"
	headers["X-App-Rate-Limit-Count"] = "1000:10,1000:600"

	headers["X-Method-Rate-Limit"] = "1:10,1:600"
	headers["X-Method-Rate-Limit-Count"] = "1000:10,1000:600"

	gock.New(fmt.Sprintf(api.BaseURLFormat, "tests")).
		Put("/").
		Reply(200).SetHeaders(headers)

	err := client.Put("tests", "/", nil, "", "")

	assert.Nil(t, err, "expecting nil error")

	err = client.Put("tests", "/", nil, "", "")

	require.Equal(t, api.RateLimitedError, err, fmt.Sprintf("want err %v, got %v", api.RateLimitedError, err))
}
