package bitaxe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/telegraf/testutil"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/unit"
)

var (
	bestDiff        = unit.SuffixedNumber(258_000_000)
	bestSessionDiff = unit.SuffixedNumber(2_880_000)
	systemInfo      = &bitaxelib.SystemInfo{
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
	swarmInfo = bitaxelib.SwarmInfo{
		{
			IP: "10.0.0.1",
		},
		{
			IP: "10.0.0.2",
		},
	}
)

type mockedSystemFetcher struct {
	error bool
}

func (f *mockedSystemFetcher) Fetch(_ string) (*bitaxelib.SystemInfo, error) {
	if f.error {
		return nil, fmt.Errorf("mocked fetcher error")
	}

	return systemInfo, nil
}

type mockedSwarmFetcher struct {
	error bool
}

func (f *mockedSwarmFetcher) Fetch(_ string) (bitaxelib.SwarmInfo, error) {
	if f.error {
		return bitaxelib.SwarmInfo{}, fmt.Errorf("mocked fetcher error")
	}

	return swarmInfo, nil
}

func TestGatherWithOneDevice(t *testing.T) {
	bitaxe := &plugin{
		systemFetcher: &mockedSystemFetcher{},
		Devices:       []string{"10.0.0.1"},
	}

	acc := &testutil.Accumulator{}
	err := bitaxe.Gather(acc)

	require.NoError(t, err)
	require.Equal(t, 16, acc.NFields())
	metric := bitaxeMetric(*systemInfo)

	acc.AssertContainsTaggedFields(t, "bitaxe", metric.Fields(), metric.Tags())
}

func TestGatherWithOneDeviceWithError(t *testing.T) {
	bitaxe := &plugin{
		systemFetcher: &mockedSystemFetcher{error: true},
		Devices:       []string{"10.0.0.1"},
	}

	acc := &testutil.Accumulator{}
	err := bitaxe.Gather(acc)

	require.Error(t, err)
}

func TestGatherWithDevicesInSwarm(t *testing.T) {
	bitaxe := &plugin{
		systemFetcher:  &mockedSystemFetcher{},
		swarmFetcher:   &mockedSwarmFetcher{},
		Devices:        []string{"10.0.0.1"},
		AllowSwarmMode: true,
	}

	acc := &testutil.Accumulator{}
	err := bitaxe.Gather(acc)

	require.NoError(t, err)
	require.Equal(t, 32, acc.NFields())
	metric := bitaxeMetric(*systemInfo)

	acc.AssertContainsTaggedFields(t, "bitaxe", metric.Fields(), metric.Tags())
}

func TestGatherWithDevicesInSwarmWithError(t *testing.T) {
	bitaxe := &plugin{
		systemFetcher:  &mockedSystemFetcher{},
		swarmFetcher:   &mockedSwarmFetcher{error: true},
		Devices:        []string{"10.0.0.1"},
		AllowSwarmMode: true,
	}

	acc := &testutil.Accumulator{}
	err := bitaxe.Gather(acc)

	require.Error(t, err)
}
