[![Downloads](https://img.shields.io/github/downloads/mendelgusmao/bitaxe-telegraf-plugin/total.svg)](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/releases)
[![Build](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/actions/workflows/build.yml/badge.svg)](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/actions/workflows/build.yml)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=mendelgusmao_bitaxe-telegraf-plugin&metric=coverage)](https://sonarcloud.io/summary/new_code?id=mendelgusmao_bitaxe-telegraf-plugin)

## About bitaxe-telegraf-plugin
This [Telegraf](https://github.com/influxdata/telegraf) input plugin gathers data from crypto miners running [BitAxe](https://bitaxe.org/) firmware.

### Installation
To install the plugin you have to download a suitable [release archive](https://github.com/mendelgusmao/bitaxe-telegraf-plugin/releases) and extract it or build it from source by cloning the repository and issueing a simple
```
make
```
To build the plugin, Go version 1.16 or higher is required. The resulting plugin binary will be written to **./build/bin**.
Copy the either extracted or built plugin binary to a location of your choice (e.g. /usr/local/bin/telegraf/).

### Configuration
This is an [external plugin](https://github.com/influxdata/telegraf/blob/master/docs/EXTERNAL_PLUGINS.md) which has to be integrated via Telegraf's [execd plugin](https://github.com/influxdata/telegraf/tree/master/plugins/inputs/execd).

To use it you have to create a plugin specific config file (e.g. /etc/telegraf/bitaxe.conf) with following template content:

```toml
# Fetch bitaxe data from compatible miners
[[inputs.bitaxe]]
  # miners is an array of the miners' hostnames or IP addresses 
  miners = ["192.168.1.1", "bitcoin-miner"]
```
The most important setting is the **miners** line. It defines the miners' IP addresses or hostnames to query. At least one address has to be defined.

To enable the plugin within your Telegraf instance, add the following section to your **telegraf.conf**
```toml
[[inputs.execd]]
  command = ["/usr/local/bin/telegraf/bitaxe-telegraf-plugin", "-config", "/etc/telegraf/bitaxe.conf", "-poll_interval", "60s"]
  signal = "none"
```

### TODO
Implement the gathering of swarm data. **Contributions are welcome!**

### License
This project is subject to the the MIT License.
See [LICENSE](./LICENSE) information for details.
