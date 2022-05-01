package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Kyagara/equinox/api"
)

type InternalClient struct {
	key   string
	debug bool
	retry bool
	http  *http.Client
	log   *Logger
}

const (
	LogRequestFormat = "[%s '%s'] %s\n"
)

// Creates an EquinoxConfig for tests.
func NewTestEquinoxConfig() *api.EquinoxConfig {
	return &api.EquinoxConfig{
		Key:     "RIOT_API_KEY",
		Debug:   true,
		Timeout: 10,
		Retry:   true,
	}
}

// Returns a new client using the API key provided.
func NewInternalClient(config *api.EquinoxConfig) *InternalClient {
	return &InternalClient{
		key:   config.Key,
		debug: config.Debug,
		retry: config.Retry,
		http:  &http.Client{Timeout: config.Timeout * time.Second},
		log:   NewLogger(),
	}
}

// Executes a http request.
func (c *InternalClient) Do(method string, route interface{}, endpoint string, body interface{}, object interface{}) error {
	baseUrl := fmt.Sprintf(api.BaseURLFormat, route)

	// Creating a new *http.Request.
	req, err := c.newRequest(method, fmt.Sprintf("%s%s", baseUrl, endpoint), body)

	if err != nil {
		return err
	}

	// Sending http request and returning the response.
	res, err := c.sendRequest(req, 0)

	if err != nil {
		return err
	}

	// Decoding the body into the endpoint method response object.
	if err = json.Unmarshal(res, &object); err != nil {
		return err
	}

	return nil
}

// Sends a http request.
func (c *InternalClient) sendRequest(req *http.Request, retryCount int8) ([]byte, error) {
	if c.retry && retryCount == 1 {
		msg := fmt.Sprintf(LogRequestFormat, req.Method, req.URL.Path, fmt.Sprintf("Failed %d time, stopping", retryCount))

		return nil, fmt.Errorf(msg)
	}

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, req.Method, req.URL.Path, "Requesting")
	}

	// Sending request.
	res, err := c.http.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusUnauthorized {
		if c.debug {
			c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, "Unauthorized")
		}

		return nil, api.UnauthorizedError
	}

	if res.StatusCode == http.StatusForbidden {
		if c.debug {
			c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, "Forbidden")
		}

		return nil, api.ForbiddenError
	}

	// If the API returns a 429 code.
	if res.StatusCode == http.StatusTooManyRequests {
		if c.debug {
			c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, "Too many requests")
		}

		// If Retry is disabled just return an error.
		if !c.retry {
			return nil, c.newErrorResponse(res)
		}

		retryAfter := res.Header.Get("Retry-After")

		// If the header isn't found, don't retry and return error.
		if retryAfter == "" {
			if c.debug {
				c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, "Retry-After header not found, not retrying")
			}

			return nil, fmt.Errorf("rate limited, status code 429")
		}

		seconds, err := strconv.Atoi(retryAfter)

		if err != nil {
			return nil, err
		}

		if c.debug {
			c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, fmt.Sprintf("Retrying in %ds", seconds))
		}

		time.Sleep(time.Duration(seconds) * time.Second)

		return c.sendRequest(req, retryCount+1)
	}

	// If the API returns a 404 code, return an error.
	if res.StatusCode == http.StatusNotFound {
		if c.debug {
			c.log.Warn.Printf(LogRequestFormat, req.Method, req.URL.Path, "Not Found")
		}

		return nil, api.NotFoundError
	}

	// If the status code is lower than 200 or higher than 400, return an error.
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusBadRequest {
		if c.debug {
			c.log.Error.Printf(LogRequestFormat, req.Method, req.URL.Path, "Returned an error")
		}

		return nil, c.newErrorResponse(res)
	}

	if c.debug {
		c.log.Info.Printf(LogRequestFormat, req.Method, req.URL.Path, "Request successful")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Creates a new *http.Request and sets headers.
func (c *InternalClient) newRequest(method string, url string, body interface{}) (*http.Request, error) {
	if c.key == "" {
		return nil, fmt.Errorf("API Key not provided")
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("X-Riot-Token", c.key)

	return req, nil
}

// Returns an error from the *http.Response provided.
func (c *InternalClient) newErrorResponse(res *http.Response) error {
	var errRes api.ErrorResponse

	err := json.NewDecoder(res.Body).Decode(&errRes)

	if err != nil {
		return errRes
	}

	return errRes
}
