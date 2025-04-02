package bitaxe

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
)

var (
	//go:embed bitaxe.conf
	sampleConfig         string
	gatherError          = "plugin.Gather: %v"
	emptyDevicesError    = "at least one device address should be specified"
	deviceSkippedMessage = "device skipped and removed: %s\n"
)

type systemFetcher interface {
	Fetch(string) (*bitaxelib.SystemInfo, error)
}

type plugin struct {
	Devices       []string      `toml:"devices"`
	Timeout       time.Duration `toml:"timeout"`
	systemFetcher systemFetcher
}

func (p *plugin) Init() error {
	if len(p.Devices) == 0 {
		return errors.New(emptyDevicesError)
	}

	p.systemFetcher = bitaxelib.NewSystemFetcher(p.Timeout)

	return nil
}

func (p *plugin) Gather(acc telegraf.Accumulator) error {
	var (
		wg           sync.WaitGroup
		gatherErrors chan error
	)

	wg.Add(len(p.Devices))

	for index, deviceAddress := range p.Devices {
		go func() {
			defer wg.Done()

			systemInfo, err := p.systemFetcher.Fetch(deviceAddress)

			if err != nil {
				gatherErrors <- fmt.Errorf(gatherError, err)
			}

			if systemInfo == nil {
				log.Printf(deviceSkippedMessage, deviceAddress)
				p.Devices = append(p.Devices[:index], p.Devices[index+1:]...)
				return
			}

			metric := bitaxeMetric(*systemInfo)
			acc.AddFields("bitaxe", metric.Fields(), metric.Tags())

		}()
	}

	wg.Wait()

	select {
	case err := <-gatherErrors:
		return err
	case <-time.After(1 * time.Second):
		return nil
	}
}

func (*plugin) SampleConfig() string {
	return sampleConfig
}

func init() {
	inputs.Add("bitaxe", func() telegraf.Input {
		timeout, _ := time.ParseDuration("5s")

		return &plugin{
			Devices: []string{},
			Timeout: timeout,
		}
	})
}
