package bitaxe

import (
	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/unit"
)

const (
	fetchError            = "bitaxe.Fetch: fetching %s: %v"
	unexpectedStatusError = "bitaxe.Fetch: fetching %s: unexpected status `%s`"
)

type SystemInfo struct {
	Power             float64              `json:"power"`
	Voltage           float64              `json:"voltage"`
	CoreVoltage       int                  `json:"coreVoltage"`
	CoreVoltageActual int                  `json:"coreVoltageActual"`
	Current           float64              `json:"current"`
	FanSpeed          int                  `json:"fanspeed"`
	FanSpeedRpm       int                  `json:"fanSpeedRpm"`
	Temp              int                  `json:"temp"`
	OverheatMode      int                  `json:"overheat_mode"`
	HashRate          float64              `json:"hashRate"`
	BestDiff          *unit.SuffixedNumber `json:"bestDiff"`
	BestSessionDiff   *unit.SuffixedNumber `json:"bestSessionDiff"`
	FreeHeap          int                  `json:"freeHeap"`
	Frequency         int                  `json:"frequency"`
	Ssid              string               `json:"-"`
	Hostname          string               `json:"hostname"`
	WifiStatus        string               `json:"-"`
	SharesAccepted    int                  `json:"sharesAccepted"`
	SharesRejected    int                  `json:"sharesRejected"`
	UptimeSeconds     int                  `json:"uptimeSeconds"`
	ASICModel         string               `json:"ASICModel"`
	ASICCount         int                  `json:"asicCount"`
	SmallCoreCount    int                  `json:"smallCoreCount"`
	StratumURL        string               `json:"stratumURL"`
	StratumPort       int                  `json:"stratumPort"`
	StratumUser       string               `json:"stratumUser"`
	Version           string               `json:"version"`
	BoardVersion      string               `json:"boardVersion"`
	RunningPartition  string               `json:"-"`
	FlipScreen        int                  `json:"-"`
	InvertScreen      int                  `json:"-"`
	InvertFanPolarity int                  `json:"-"`
	AutoFanSpeed      int                  `json:"autofanspeed"`
}

type SwarmInfo []struct {
	IP string `json:"ip"`
}

func (i SwarmInfo) Addresses() []string {
	h := make([]string, len(i))

	for index, v := range i {
		h[index] = v.IP
	}

	return h
}
