package api

import (
	"os"
	"time"
)

// An config object for the EquinoxClient
type EquinoxConfig struct {
	// Riot API Key
	Key string
	// Debug mode. Default: false
	Debug bool
	// Timeout for http.Request in milliseconds, 0 disables it. Default: 2000 milliseconds
	Timeout time.Duration
	// Retry request if it returns a 429 status code. Default: true
	Retry bool
	// Retry count. Default: 1
	RetryCount int8
}

// Creates an EquinoxConfig for tests
func NewTestEquinoxConfig() *EquinoxConfig {
	return &EquinoxConfig{
		Key:        os.Getenv("RIOT_API_KEY"),
		Debug:      true,
		Timeout:    2000,
		Retry:      true,
		RetryCount: 1,
	}
}
