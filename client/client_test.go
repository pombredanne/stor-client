package storclient_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/avast/stor-client/client"
	"github.com/stretchr/testify/assert"
)

func TestNewDefault(t *testing.T) {
	url, _ := url.Parse("http://stor.server.com")

	client, err := storclient.New(*url, "some_dir", storclient.StorClientOpts{})
	assert.NoError(t, err)

	assert.Equal(t, client.Max, 4)

	expectedTimeout, _ := time.ParseDuration("30s")
	assert.Equal(t, client.Timeout, expectedTimeout)
}

func TestNewInfinityTimeout(t *testing.T) {
	url, _ := url.Parse("http://stor.server.com")

	client, err := storclient.New(*url, "some_dir", storclient.StorClientOpts{Timeout: -1})
	assert.NoError(t, err)

	expectedTimeout, _ := time.ParseDuration("0s")
	assert.Equal(t, client.Timeout, expectedTimeout)
}
