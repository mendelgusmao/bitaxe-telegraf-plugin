package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	path                  = "/api/system/info"
	endpoint              = "http://%s" + path
	fetchError            = "bitaxe.Fetch: %v"
	unexpectedStatusError = "bitaxe.Fetch: unexpected status `%s`"
)

type bitaxefetcher struct{}

func NewFetcher() *bitaxefetcher {
	return &bitaxefetcher{}
}

func (h *bitaxefetcher) Fetch(address string) (*MinerInfo, error) {
	address = fmt.Sprintf(endpoint, address)
	response, err := http.Get(address)

	if err != nil {
		return nil, fmt.Errorf(fetchError, err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(unexpectedStatusError, response.Status)
	}

	miner := &MinerInfo{}

	if err := json.NewDecoder(response.Body).Decode(miner); err != nil {
		return nil, fmt.Errorf(fetchError, err)
	}

	return miner, nil
}
