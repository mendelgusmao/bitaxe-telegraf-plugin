package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/influxdata/telegraf/plugins/common/shim"

	_ "github.com/mendelgusmao/bitaxe-telegraf-plugin/plugins/inputs/bitaxe"
)

var (
	pollInterval = flag.Duration("poll_interval", 1*time.Second, "how often to send metrics")

	pollIntervalDisabled = flag.Bool(
		"poll_interval_disabled",
		false,
		"set to true to disable polling. You want to use this when you are sending metrics on your own schedule",
	)
	configFile = flag.String("config", "", "path to the config file for this plugin")
	err        error
)

func main() {
	flag.Parse()

	if *pollIntervalDisabled {
		*pollInterval = shim.PollIntervalDisabled
	}

	shimLayer := shim.New()

	if err = shimLayer.LoadConfig(configFile); err != nil {
		fmt.Fprintf(os.Stderr, "Err loading input: %s\n", err)
		os.Exit(1)
	}

	if err = shimLayer.Run(*pollInterval); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
}
