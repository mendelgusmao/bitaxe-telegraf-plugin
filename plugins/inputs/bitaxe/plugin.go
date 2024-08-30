package bitaxe

import (
	_ "embed"
	"errors"
	"fmt"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/set"
)

var (
	//go:embed bitaxe.conf
	sampleConfig      string
	gatherError       = "plugin.Gather: %v"
	emptyDevicesError = "at least one device address should be specified"
)

type systemFetcher interface {
	Fetch(string) (*bitaxelib.SystemInfo, error)
}

type swarmFetcher interface {
	Fetch(string) (bitaxelib.SwarmInfo, error)
}

type plugin struct {
	Devices        []string      `toml:"devices"`
	Timeout        time.Duration `toml:"timeout"`
	AllowSwarmMode bool          `toml:"allow_swarm_mode"`
	systemFetcher  systemFetcher
	swarmFetcher   swarmFetcher
}

func (p *plugin) Init() error {
	if len(p.Devices) == 0 {
		return errors.New(emptyDevicesError)
	}

	p.systemFetcher = bitaxelib.NewSystemFetcher(p.Timeout)
	p.swarmFetcher = bitaxelib.NewSwarmFetcher(p.Timeout)

	return nil
}

func (p *plugin) Gather(acc telegraf.Accumulator) error {
	devices := set.NewSet[string](p.Devices...)

	if p.AllowSwarmMode {
		swarmInfo, err := p.swarmFetcher.Fetch(p.Devices[0])

		if err != nil {
			return fmt.Errorf(gatherError, err)
		}

		for _, address := range swarmInfo.Addresses() {
			devices.Add(address)
		}
	}

	for _, deviceAddress := range devices.Values() {
		systemInfo, err := p.systemFetcher.Fetch(deviceAddress)

		if err != nil {
			return fmt.Errorf(gatherError, err)
		}

		metric := bitaxeMetric(*systemInfo)
		acc.AddFields("bitaxe", metric.Fields(), metric.Tags())
	}

	return nil
}

func (*plugin) SampleConfig() string {
	return sampleConfig
}

func init() {
	inputs.Add("bitaxe", func() telegraf.Input {
		timeout, _ := time.ParseDuration("5s")

		return &plugin{
			Devices:        []string{},
			Timeout:        timeout,
			AllowSwarmMode: false,
		}
	})
}
