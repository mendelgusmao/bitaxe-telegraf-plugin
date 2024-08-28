package bitaxe

import (
	_ "embed"
	"errors"
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
)

var (
	//go:embed bitaxe.conf
	sampleConfig      string
	gatherError       = "bitaxeinput.Gather: %v"
	emptyDevicesError = "at least one device address should be specified"
)

type systemFetcher interface {
	Fetch(address string) (*bitaxelib.SystemInfo, error)
}

type swarmFetcher interface {
	Fetch(address string) (bitaxelib.SwarmInfo, error)
}

type bitaxeinput struct {
	Devices        []string `toml:"devices"`
	AllowSwarmMode bool     `toml:"allow_swarm_mode"`
	systemFetcher  systemFetcher
	swarmFetcher   swarmFetcher
}

func (i *bitaxeinput) Init() error {
	if len(i.Devices) == 0 {
		return errors.New(emptyDevicesError)
	}

	i.systemFetcher = bitaxelib.NewSystemFetcher()
	i.swarmFetcher = bitaxelib.NewSwarmFetcher()

	return nil
}

func (i *bitaxeinput) Gather(acc telegraf.Accumulator) error {
	devices := append([]string{}, i.Devices...)

	if i.AllowSwarmMode {
		swarmInfo, err := i.swarmFetcher.Fetch(i.Devices[0])

		if err != nil {
			return fmt.Errorf(gatherError, err)
		}

		devices = swarmInfo.UniqueDevices(i.Devices)
	}

	for _, deviceAddress := range devices {
		systemInfo, err := i.systemFetcher.Fetch(deviceAddress)

		if err != nil {
			return fmt.Errorf(gatherError, err)
		}

		metric := bitaxeMetric(*systemInfo)
		acc.AddFields("bitaxe", metric.Fields(), metric.Tags())
	}

	return nil
}

func (*bitaxeinput) SampleConfig() string {
	return sampleConfig
}

func init() {
	inputs.Add("bitaxe", func() telegraf.Input {
		return &bitaxeinput{
			Devices:        []string{},
			AllowSwarmMode: false,
		}
	})
}
