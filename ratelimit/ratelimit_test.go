package ratelimit_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/internal"
	"github.com/Kyagara/equinox/ratelimit"
	"github.com/Kyagara/equinox/test/util"
	"github.com/stretchr/testify/require"
)

func TestNewLimits(t *testing.T) {
	limits := ratelimit.NewLimits()
	require.NotNil(t, limits)
	require.NotEmpty(t, limits.App)
	require.NotNil(t, limits.Methods)
}

func TestRateLimitCheck(t *testing.T) {
	client, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NoError(t, err)
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	t.Run("buckets not created", func(t *testing.T) {
		r := &ratelimit.RateLimit{
			Region: make(map[any]*ratelimit.Limits),
		}
		require.Nil(t, r.Region[equinoxReq.Route])
		ctx := context.Background()
		err := r.Take(ctx, equinoxReq)
		require.NoError(t, err)
		require.NotNil(t, r.Region[equinoxReq.Route].Methods)
		require.NotNil(t, r.Region[equinoxReq.Route].Methods[equinoxReq.MethodID])
	})

	// These tests should take around 2 seconds each

	t.Run("app rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"19:2"},
		}
		ctx := context.Background()
		err := r.Take(ctx, equinoxReq)
		require.NoError(t, err)
		r.Update(equinoxReq, &headers)
		err = r.Take(ctx, equinoxReq)
		require.NoError(t, err)
	})

	t.Run("method rate limited", func(t *testing.T) {
		r := &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.METHOD_RATE_LIMIT_HEADER:       []string{"100:2,200:2"},
			ratelimit.METHOD_RATE_LIMIT_COUNT_HEADER: []string{"1:2,199:2"},
		}
		ctx := context.Background()
		err := r.Take(ctx, equinoxReq)
		require.NoError(t, err)
		r.Update(equinoxReq, &headers)
		err = r.Take(ctx, equinoxReq)
		require.NoError(t, err)
	})

	t.Run("waiting bucket to reset", func(t *testing.T) {
		r := &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)}
		headers := http.Header{
			ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
			ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"20:2"},
		}
		ctx := context.Background()
		err = r.Take(ctx, equinoxReq)
		require.NoError(t, err)
		r.Update(equinoxReq, &headers)
		err := r.Take(ctx, equinoxReq)
		require.NoError(t, err)
	})
}

func TestLimitsDontMatch(t *testing.T) {
	client, err := internal.NewInternalClient(util.NewTestEquinoxConfig())
	require.NoError(t, err)
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
		Logger:   client.Logger("client_endpoint_method"),
	}

	r := &ratelimit.RateLimit{Region: make(map[any]*ratelimit.Limits)}
	headers := http.Header{
		ratelimit.APP_RATE_LIMIT_HEADER:       []string{"20:2"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"1:2"},
	}

	ctx, c := context.WithTimeout(context.Background(), 1*time.Second)
	defer c()
	err = r.Take(ctx, equinoxReq)
	require.NoError(t, err)
	r.Update(equinoxReq, &headers)
	err = r.Take(ctx, equinoxReq)
	require.NoError(t, err)

	headers = http.Header{
		ratelimit.APP_RATE_LIMIT_HEADER:       []string{"1:2"},
		ratelimit.APP_RATE_LIMIT_COUNT_HEADER: []string{"1:2"},
	}
	r.Update(equinoxReq, &headers)
	err = r.Take(ctx, equinoxReq)
	// The buckets should've been updated, so this request should be rate limited now
	require.Equal(t, ratelimit.ErrContextDeadlineExceeded, err)
}

func TestCheckRetryAfter(t *testing.T) {
	r := &ratelimit.RateLimit{
		Region: map[any]*ratelimit.Limits{
			"route": {
				App: ratelimit.NewLimit(ratelimit.APP_RATE_LIMIT_TYPE),
				Methods: map[string]*ratelimit.Limit{
					"method": ratelimit.NewLimit(ratelimit.METHOD_RATE_LIMIT_TYPE),
				},
			},
		},
	}

	headers := http.Header{}
	equinoxReq := &api.EquinoxRequest{
		Route:    "route",
		MethodID: "method",
	}

	headers.Set(ratelimit.RETRY_AFTER_HEADER, "")
	_, err := r.CheckRetryAfter(equinoxReq, &headers)
	require.Equal(t, ratelimit.Err429ButNoRetryAfterHeader, err)

	headers.Set(ratelimit.RETRY_AFTER_HEADER, "10")
	delay, err := r.CheckRetryAfter(equinoxReq, &headers)
	require.NoError(t, err)
	require.Equal(t, 10*time.Second, delay)
}
