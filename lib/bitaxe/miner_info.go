package bitaxe

import (
	"github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/unit"
)

type MinerInfo struct {
	Power             float64              `json:"power"`
	Voltage           float64              `json:"voltage"`
	Current           float64              `json:"current"`
	FanSpeedRpm       int                  `json:"fanSpeedRpm"`
	Temp              int                  `json:"temp"`
	HashRate          float64              `json:"hashRate"`
	BestDiff          *unit.SuffixedNumber `json:"bestDiff"`
	BestSessionDiff   *unit.SuffixedNumber `json:"bestSessionDiff"`
	FreeHeap          int                  `json:"freeHeap"`
	CoreVoltage       int                  `json:"coreVoltage"`
	CoreVoltageActual int                  `json:"coreVoltageActual"`
	Frequency         int                  `json:"frequency"`
	Ssid              string               `json:"-"`
	Hostname          string               `json:"hostname"`
	WifiStatus        string               `json:"-"`
	SharesAccepted    int                  `json:"sharesAccepted"`
	SharesRejected    int                  `json:"sharesRejected"`
	UptimeSeconds     int                  `json:"uptimeSeconds"`
	ASICModel         string               `json:"ASICModel"`
	StratumURL        string               `json:"stratumURL"`
	StratumPort       int                  `json:"stratumPort"`
	StratumUser       string               `json:"stratumUser"`
	Version           string               `json:"version"`
	BoardVersion      string               `json:"boardVersion"`
	RunningPartition  string               `json:"-"`
	Flipscreen        int                  `json:"-"`
	Invertscreen      int                  `json:"-"`
	Invertfanpolarity int                  `json:"-"`
	Autofanspeed      int                  `json:"autofanspeed"`
	Fanspeed          int                  `json:"fanspeed"`
}
