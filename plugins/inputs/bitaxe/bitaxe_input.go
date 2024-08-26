package bitaxe

import (
	_ "embed"
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
)

var (
	//go:embed bitaxe.conf
	sampleConfig string
	gatherError  = "bitaxeinput.Gather: %v"
)

type bitaxeinput struct {
	Miners  []string `toml:"miners"`
	fetcher fetcher
}

type fetcher interface {
	Fetch(address string) (*bitaxelib.MinerInfo, error)
}

func (i *bitaxeinput) Init() error {
	i.fetcher = bitaxelib.NewFetcher()

	return nil
}

func (i *bitaxeinput) Gather(acc telegraf.Accumulator) error {
	for _, minerAddress := range i.Miners {
		miner, err := i.fetcher.Fetch(minerAddress)

		if err != nil {
			return fmt.Errorf(gatherError, err)
		}

		acc.AddFields("bitaxe", miner.Fields(), miner.Tags())
	}

	return nil
}

func (*bitaxeinput) SampleConfig() string {
	return sampleConfig
}

func init() {
	inputs.Add("bitaxe", func() telegraf.Input {
		return &bitaxeinput{
			Miners: []string{},
		}
	})
}
