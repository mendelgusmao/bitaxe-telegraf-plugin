package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	swarmResource = "/api/swarm/info"
	swarmEndpoint = "http://%s" + swarmResource
)

type swarmFetcher struct {
	timeout time.Duration
}

func NewSwarmFetcher(timeout time.Duration) *swarmFetcher {
	return &swarmFetcher{
		timeout: timeout,
	}
}

func (h *swarmFetcher) Fetch(address string) (SwarmInfo, error) {
	client := http.Client{
		Timeout: h.timeout,
	}
	address = fmt.Sprintf(swarmEndpoint, address)
	response, err := client.Get(address)

	if err != nil {
		return SwarmInfo{}, fmt.Errorf(fetchError, address, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return SwarmInfo{}, fmt.Errorf(unexpectedStatusError, address, response.Status)
	}

	swarmMiners := make(SwarmInfo, 0)

	if err := json.NewDecoder(response.Body).Decode(&swarmMiners); err != nil {
		return SwarmInfo{}, fmt.Errorf(fetchError, address, err)
	}

	return swarmMiners, nil
}
