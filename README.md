[![Downloads](https://img.shields.io/github/downloads/mendelgusmao/bitaxe-telegraf-plugin/total.svg)](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/releases)
[![Build](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/actions/workflows/build.yml/badge.svg)](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/actions/workflows/build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=mendelgusmao_bitaxe-telegraf-plugin&metric=coverage)](https://sonarcloud.io/summary/new_code?id=mendelgusmao_bitaxe-telegraf-plugin)

## About bitaxe-telegraf-plugin
This [Telegraf](https://github.com/influxdata/telegraf) input plugin gathers data from crypto miners running [Bitaxe](https://bitaxe.org/) firmware.

### Installation
To install the plugin you have to download a suitable [release archive](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/releases) and extract it or build it from source by cloning the repository and issueing a simple
```
make
```
To build the plugin, Go version 1.23 or higher is required. The resulting plugin binary will be written to **./build/bin**.
Copy the either extracted or built plugin binary to a location of your choice (e.g. /usr/local/bin/telegraf/).

### Configuration
This is an [external plugin](https://github.com/influxdata/telegraf/blob/master/docs/EXTERNAL_PLUGINS.md) which has to be integrated via Telegraf's [execd plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/execd).

To use it you have to create a plugin specific config file (e.g. /etc/telegraf/bitaxe.conf) with following template content:

```toml
[[inputs.bitaxe]]
  # devices is an array of the devices' hostnames or IP addresses 
  devices = ["192.168.1.1", "crypto-miner"]

  ## Amount of time allowed to complete the HTTP request
  # timeout = "5s"

  # allow_swarm_mode tells the gatherer to fetch swarm data from the
  # first device in devices list and then from the hosts in the response.
  # if the swarm list is empty, the gatherer falls back to the main devices list
  # allow_swarm_mode = false
```
The most important setting is the **devices** line. It defines the miners' IP addresses or hostnames to query. At least one address has to be defined.

To enable the plugin within your Telegraf instance, add the following section to your **telegraf.conf**
```toml
[[inputs.execd]]
  command = ["/usr/local/bin/telegraf/bitaxe-telegraf-plugin", "-config", "/etc/telegraf/bitaxe.conf", "-poll_interval", "60s"]
  signal = "none"
```

### License
This project is subject to the the MIT License.
See [LICENSE](./LICENSE) information for details.
