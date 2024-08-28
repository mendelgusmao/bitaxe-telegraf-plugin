package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	systemResource = "/api/system/info"
	systemEndpoint = "http://%s" + systemResource
)

type systemFetcher struct{}

func NewSystemFetcher() *systemFetcher {
	return &systemFetcher{}
}

func (h *systemFetcher) Fetch(address string) (*SystemInfo, error) {
	address = fmt.Sprintf(systemEndpoint, address)
	response, err := http.Get(address)

	if err != nil {
		return nil, fmt.Errorf(fetchError, address, err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(unexpectedStatusError, address, response.Status)
	}

	miner := &SystemInfo{}

	if err := json.NewDecoder(response.Body).Decode(miner); err != nil {
		return nil, fmt.Errorf(fetchError, address, err)
	}

	return miner, nil
}
