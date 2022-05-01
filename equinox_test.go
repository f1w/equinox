package equinox_test

import (
	"testing"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/api"
	"github.com/stretchr/testify/assert"
)

func TestEquinoxClient(t *testing.T) {
	client, err := equinox.NewClient("RIOT_API_KEY")

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}

func TestEquinoxClientWithConfig(t *testing.T) {
	config := &api.EquinoxConfig{
		Key:        "RIOT_API_KEY",
		Debug:      true,
		Timeout:    10,
		Retry:      true,
		RetryCount: 1,
	}

	client, err := equinox.NewClientWithConfig(config)

	assert.Nil(t, err, "expecting nil error")

	assert.NotNil(t, client, "expecting non-nil Equinox client")
}
