package service

import (
	"testing"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/stretchr/testify/require"
)

func TestInitHttpClientLeavesTimeoutDisabledWhenRelayTimeoutZero(t *testing.T) {
	oldRelayTimeout := common.RelayTimeout
	oldClient := httpClient
	t.Cleanup(func() {
		common.RelayTimeout = oldRelayTimeout
		httpClient = oldClient
	})

	common.RelayTimeout = 0
	InitHttpClient()

	require.NotNil(t, GetHttpClient())
	require.Equal(t, time.Duration(0), GetHttpClient().Timeout)
}

func TestNewProxyHttpClientLeavesTimeoutDisabledWhenRelayTimeoutZero(t *testing.T) {
	oldRelayTimeout := common.RelayTimeout
	ResetProxyClientCache()
	t.Cleanup(func() {
		common.RelayTimeout = oldRelayTimeout
		ResetProxyClientCache()
	})

	common.RelayTimeout = 0
	client, err := NewProxyHttpClient("http://127.0.0.1:8080")

	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, time.Duration(0), client.Timeout)
}

func TestNewProxyHttpClientAppliesConfiguredRelayTimeout(t *testing.T) {
	oldRelayTimeout := common.RelayTimeout
	ResetProxyClientCache()
	t.Cleanup(func() {
		common.RelayTimeout = oldRelayTimeout
		ResetProxyClientCache()
	})

	common.RelayTimeout = 60
	client, err := NewProxyHttpClient("http://127.0.0.1:8080")

	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, 60*time.Second, client.Timeout)
}
