package bitaxe

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/unit"
	"github.com/stretchr/testify/require"
)

func TestFetchSystemInfo(t *testing.T) {
	response := map[string]any{
		"power":             11.979999542236328,
		"voltage":           5171.25,
		"current":           2323.75,
		"fanSpeedRpm":       5654,
		"temp":              41,
		"hashRate":          570.041489032058,
		"bestDiff":          "258M",
		"bestSessionDiff":   "2.88M",
		"freeHeap":          165372,
		"coreVoltage":       1200,
		"coreVoltageActual": 1213,
		"frequency":         550,
		"ssid":              "network",
		"hostname":          "bitaxe",
		"wifiStatus":        "Connected!",
		"sharesAccepted":    4963,
		"sharesRejected":    4,
		"uptimeSeconds":     44688,
		"ASICModel":         "BM1366",
		"stratumURL":        "cool-pool.whatever",
		"stratumPort":       3333,
		"stratumUser":       "luckyperson.001",
		"version":           "v2.1.8",
		"boardVersion":      "204",
		"runningPartition":  "ota_0",
		"flipscreen":        1,
		"invertscreen":      0,
		"invertfanpolarity": 1,
		"autofanspeed":      1,
		"fanspeed":          100,
	}

	s := serve(t, systemResource, http.StatusOK, response)
	defer s.Listener.Close()

	miner, err := NewSystemFetcher().Fetch(s.Listener.Addr().String())

	require.NoError(t, err)

	bestDiff := unit.SuffixedNumber(258_000_000)
	bestSessionDiff := unit.SuffixedNumber(2_880_000)

	expected := SystemInfo{
		Power:             11.979999542236328,
		Voltage:           5171.25,
		Current:           2323.75,
		FanSpeedRpm:       5654,
		Temp:              41,
		HashRate:          570.041489032058,
		BestDiff:          &bestDiff,
		BestSessionDiff:   &bestSessionDiff,
		FreeHeap:          165372,
		CoreVoltage:       1200,
		CoreVoltageActual: 1213,
		Frequency:         550,
		Hostname:          "bitaxe",
		SharesAccepted:    4963,
		SharesRejected:    4,
		UptimeSeconds:     44688,
		ASICModel:         "BM1366",
		StratumURL:        "cool-pool.whatever",
		StratumPort:       3333,
		StratumUser:       "luckyperson.001",
		Version:           "v2.1.8",
		BoardVersion:      "204",
		AutoFanSpeed:      1,
		FanSpeed:          100,
	}

	require.Equal(t, expected, *miner, "systemInfo is different from expected")
}

func TestFetchSystemInfoWrongAddress(t *testing.T) {
	_, err := NewSystemFetcher().Fetch("127.0.0.1:1")
	require.Error(t, err)
}

func TestFetchSystemInfoWithUnexpectedHTTPStatus(t *testing.T) {
	s := serve(t, systemResource, http.StatusInternalServerError, nil)
	defer s.Listener.Close()

	_, err := NewSystemFetcher().Fetch(s.Listener.Addr().String())
	require.Error(t, err)
}

func serve(t *testing.T, path string, httpStatus int, response any) *httptest.Server {
	body, err := json.Marshal(response)

	require.NoError(t, err)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, r.URL.Path, path)

		w.WriteHeader(httpStatus)
		w.Write(body)
	}))

	return server
}
