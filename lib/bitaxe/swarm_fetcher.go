package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	swarmResource = "/api/swarm/info"
	swarmEndpoint = "http://%s" + swarmResource
)

type swarmFetcher struct{}

func NewSwarmFetcher() *swarmFetcher {
	return &swarmFetcher{}
}

func (h *swarmFetcher) Fetch(address string) (SwarmInfo, error) {
	address = fmt.Sprintf(swarmEndpoint, address)
	response, err := http.Get(address)

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
