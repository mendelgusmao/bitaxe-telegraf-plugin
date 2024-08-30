package bitaxe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	systemResource = "/api/system/info"
	systemEndpoint = "http://%s" + systemResource
)

type systemFetcher struct {
	timeout time.Duration
}

func NewSystemFetcher(timeout time.Duration) *systemFetcher {
	return &systemFetcher{
		timeout: timeout,
	}
}

func (f *systemFetcher) Fetch(address string) (*SystemInfo, error) {
	client := http.Client{
		Timeout: f.timeout,
	}
	address = fmt.Sprintf(systemEndpoint, address)
	response, err := client.Get(address)

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
