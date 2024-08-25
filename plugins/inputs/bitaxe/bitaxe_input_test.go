package bitaxe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/influxdata/telegraf/testutil"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/unit"
)

var (
	bestDiff        = unit.SuffixedNumber(258_000_000)
	bestSessionDiff = unit.SuffixedNumber(2_880_000)
	minerInfo       = &bitaxelib.MinerInfo{
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
		Autofanspeed:      1,
		Fanspeed:          100,
	}
)

type mockedFetcher struct{}

func (h *mockedFetcher) Fetch(_ string) (*bitaxelib.MinerInfo, error) {
	return minerInfo, nil
}

func TestFetch(t *testing.T) {
	bitaxe := &bitaxeinput{
		fetcher: &mockedFetcher{},
		Miners:  []string{""},
	}

	acc := &testutil.Accumulator{}
	err := bitaxe.Gather(acc)

	require.NoError(t, err)
	require.Equal(t, 16, acc.NFields())

	acc.AssertContainsTaggedFields(t, "bitaxe", minerInfo.Fields(), minerInfo.Tags())
}
