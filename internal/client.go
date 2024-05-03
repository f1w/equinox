package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"

	jsonv2 "github.com/go-json-experiment/json"
	"github.com/rs/zerolog"

	"github.com/Kyagara/equinox/api"
	"github.com/Kyagara/equinox/cache"
	"github.com/Kyagara/equinox/ratelimit"
)

var (
	ErrMaxRetries     = errors.New("max retries reached")
	ErrContextIsNil   = errors.New("context must be non-nil")
	ErrKeyNotProvided = errors.New("api key not provided")
)

var (
	apiHeaders = http.Header{
		"X-Riot-Token": {""},
		"Accept":       {"application/json"},
		"Content-Type": {"application/json"},
	}
)

type Client struct {
	http               *http.Client
	cache              *cache.Cache
	ratelimit          *ratelimit.RateLimit
	loggers            loggers
	key                string
	maxRetries         int
	jitter             time.Duration
	IsCacheEnabled     bool
	IsRateLimitEnabled bool
	IsRetryEnabled     bool
}

func NewInternalClient(config api.EquinoxConfig, h *http.Client, c *cache.Cache, r *ratelimit.RateLimit) (*Client, error) {
	if config.Key == "" {
		return nil, ErrKeyNotProvided
	}

	if h == nil {
		h = &http.Client{}
	}
	if c == nil {
		c = &cache.Cache{TTL: 0}
	}
	if r == nil {
		r = &ratelimit.RateLimit{Enabled: false}
	}

	client := &Client{
		key:  config.Key,
		http: h,
		loggers: loggers{
			main:    NewLogger(config),
			methods: make(map[string]zerolog.Logger, 1),
			mutex:   sync.Mutex{},
		},
		cache:              c,
		ratelimit:          r,
		maxRetries:         config.Retry.MaxRetries,
		jitter:             config.Retry.Jitter,
		IsCacheEnabled:     c.TTL > 0,
		IsRateLimitEnabled: r.Enabled,
		IsRetryEnabled:     config.Retry.MaxRetries > 0,
	}

	apiHeaders.Set("X-Riot-Token", config.Key)
	return client, nil
}

// Creates a new 'EquinoxRequest' object for the 'Execute' and 'ExecuteBytes' methods.
func (c *Client) Request(ctx context.Context, logger zerolog.Logger, httpMethod string, urlComponents []string, methodID string, body any) (api.EquinoxRequest, error) {
	logger.Trace().Msg("Request")

	if ctx == nil {
		return api.EquinoxRequest{}, ErrContextIsNil
	}

	var bodyReader io.Reader
	if body != nil {
		json, err := jsonv2.Marshal(body)
		if err != nil {
			logger.Error().Err(err).Msg("Error marshalling body")
			return api.EquinoxRequest{}, err
		}
		bodyReader = bytes.NewReader(json)
	}

	url := strings.Join(urlComponents, "")

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, bodyReader)
	if err != nil {
		logger.Error().Err(err).Msg("Error creating HTTP request")
		return api.EquinoxRequest{}, err
	}

	equinoxReq := api.EquinoxRequest{
		Logger:   logger,
		URL:      url,
		Route:    urlComponents[1],
		MethodID: methodID,
		Request:  request,
	}

	request.Header = apiHeaders

	return equinoxReq, nil
}

// Executes a 'EquinoxRequest', checks cache and unmarshals the response into 'target'.
func (c *Client) Execute(ctx context.Context, equinoxReq api.EquinoxRequest, target any) error {
	equinoxReq.Logger.Trace().Msg("Execute")

	if ctx == nil {
		return ErrContextIsNil
	}

	url := GetURLWithAuthorizationHash(equinoxReq)

	if c.IsCacheEnabled && equinoxReq.Request.Method == http.MethodGet {
		item, err := c.cache.Get(ctx, url)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error retrieving cached response")
		}

		if item != nil {
			equinoxReq.Logger.Debug().Msg("Cache hit")

			// Only valid json is cached, so unmarshal shouldn't fail
			_ = jsonv2.Unmarshal(item, target)

			return nil
		}
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Do failed")
		return err
	}
	defer response.Body.Close()

	if !c.IsCacheEnabled {
		err := jsonv2.UnmarshalRead(response.Body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
			return err
		}

		return nil
	}

	// Cache is enabled

	if equinoxReq.Request.Method == http.MethodGet {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error reading response")
			return err
		}

		err = jsonv2.Unmarshal(body, target)
		if err != nil {
			equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling body")
			return err
		}

		// Only cache valid json responses

		err = c.cache.Set(ctx, url, body)
		if err == nil {
			equinoxReq.Logger.Debug().Msg("Cache set")
		} else {
			equinoxReq.Logger.Error().Err(err).Msg("Error caching item")
		}

		return nil
	}

	err = jsonv2.UnmarshalRead(response.Body, target)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error unmarshalling response")
		return err
	}

	return nil
}

// Executes a 'EquinoxRequest', skips cache check and returns []byte.
func (c *Client) ExecuteBytes(ctx context.Context, equinoxReq api.EquinoxRequest) ([]byte, error) {
	equinoxReq.Logger.Trace().Msg("ExecuteBytes")

	if ctx == nil {
		return nil, ErrContextIsNil
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Reserve(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID)
		if err != nil {
			return nil, err
		}
	}

	response, err := c.Do(ctx, equinoxReq)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Do failed")
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		equinoxReq.Logger.Error().Err(err).Msg("Error reading response body")
		return nil, err
	}

	return body, nil
}

// Sends the request, retries if enabled.
func (c *Client) Do(ctx context.Context, equinoxReq api.EquinoxRequest) (*http.Response, error) {
	equinoxReq.Logger.Trace().Msg("Do")

	var httpErr error

	// MaxRetries+1 to run this loop at least once
	for i := 0; i < c.maxRetries+1; i++ {
		response, err := c.http.Do(equinoxReq.Request)
		if err != nil {
			// Stop if the http.Client itself returns any error
			return nil, err
		}

		delay, retryable, err := c.checkResponse(ctx, equinoxReq, response)
		if err == nil && delay == 0 {
			return response, nil
		}

		if !c.IsRetryEnabled || !retryable {
			return nil, err
		}

		if i < c.maxRetries {
			// Exponential backoff with jitter
			wait := delay*time.Duration(math.Pow(2, float64(i))) + c.jitter
			equinoxReq.Logger.Warn().Str("status_code", response.Status).Dur("wait", wait).Int("retries", i).Msg("Retrying request")
			err := ratelimit.WaitN(ctx, time.Now().Add(wait), wait)
			if err != nil {
				return nil, err
			}
		}

		httpErr = err
	}

	return nil, fmt.Errorf("%w: %w", ErrMaxRetries, httpErr)
}

// Checks the response, check if it contains a 'Retry-After' header and whether it should be retried (StatusCode within range 429-599).
func (c *Client) checkResponse(ctx context.Context, equinoxReq api.EquinoxRequest, response *http.Response) (time.Duration, bool, error) {
	// Delay in milliseconds
	var retryAfter time.Duration

	if response.StatusCode == http.StatusTooManyRequests {
		retryAfter = ratelimit.GetRetryAfterHeader(ratelimit.RETRY_AFTER_HEADER)
	}

	if c.IsRateLimitEnabled {
		err := c.ratelimit.Update(ctx, equinoxReq.Logger, equinoxReq.Route, equinoxReq.MethodID, response.Header, retryAfter)
		if err != nil {
			return 0, false, err
		}
	}

	// 2xx responses
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return 0, false, nil
	}

	err := api.StatusCodeToError(response.StatusCode)
	if err != nil {
		// 429 and 5xx responses will be retried
		if response.StatusCode == http.StatusTooManyRequests || (response.StatusCode >= 500 && response.StatusCode < 600) {
			return retryAfter, true, err
		}

		return 0, false, err
	}

	return 0, false, fmt.Errorf("unexpected status code: %d", response.StatusCode)
}

// Returns an URL with a hashed Authorization key if it exists.
func GetURLWithAuthorizationHash(req api.EquinoxRequest) string {
	auth := req.Request.Header.Get("Authorization")
	if auth == "" {
		return req.URL
	}

	hash := sha256.New()
	_, _ = hash.Write([]byte(auth))
	hashedAuth := hash.Sum(nil)

	return fmt.Sprintf("%s-%x", req.URL, hashedAuth)
}
