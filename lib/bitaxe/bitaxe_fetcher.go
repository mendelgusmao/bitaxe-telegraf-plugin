package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	path                  = "/api/system/info"
	endpoint              = "http://%s" + path
	fetchError            = "bitaxe.Fetch: fetching %s: %v"
	unexpectedStatusError = "bitaxe.Fetch: fetching %s: unexpected status `%s`"
)

type bitaxefetcher struct{}

func NewFetcher() *bitaxefetcher {
	return &bitaxefetcher{}
}

func (h *bitaxefetcher) Fetch(address string) (*MinerInfo, error) {
	address = fmt.Sprintf(endpoint, address)
	response, err := http.Get(address)

	if err != nil {
		return nil, fmt.Errorf(fetchError, address, err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(unexpectedStatusError, address, response.Status)
	}

	miner := &MinerInfo{}

	if err := json.NewDecoder(response.Body).Decode(miner); err != nil {
		return nil, fmt.Errorf(fetchError, address, err)
	}

	return miner, nil
}
