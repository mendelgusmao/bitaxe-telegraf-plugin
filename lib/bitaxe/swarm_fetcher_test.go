package bitaxe

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	response := []map[string]any{
		{
			"ip": "10.0.0.1",
		},
	}

	s := serve(t, swarmResource, http.StatusOK, response)
	defer s.Listener.Close()

	swarmInfo, err := NewSwarmFetcher(1 * time.Second).Fetch(s.Listener.Addr().String())

	require.NoError(t, err)

	expected := SwarmInfo{
		{
			IP: "10.0.0.1",
		},
	}

	require.Equal(t, expected, swarmInfo, "swarmInfo' slice is different from expected")
}

func TestFetchWrongAddress(t *testing.T) {
	_, err := NewSwarmFetcher(1 * time.Second).Fetch("127.0.0.1:1")
	require.Error(t, err)
}

func TestFetchWithUnexpectedHTTPStatus(t *testing.T) {
	s := serve(t, swarmResource, http.StatusInternalServerError, nil)
	defer s.Listener.Close()

	_, err := NewSwarmFetcher(1 * time.Second).Fetch(s.Listener.Addr().String())
	require.Error(t, err)
}
