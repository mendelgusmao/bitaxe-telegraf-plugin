package bitaxe

import (
	"fmt"
	"strconv"

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

func (m *MinerInfo) Tags() map[string]string {
	return map[string]string{
		"hostname":       m.Hostname,
		"asic_model":     m.ASICModel,
		"stratum_url":    fmt.Sprintf("%s:%d", m.StratumURL, m.StratumPort),
		"stratum_user":   m.StratumUser,
		"os_version":     m.Version,
		"board_version":  m.BoardVersion,
		"auto_fan_speed": strconv.Itoa(m.Autofanspeed),
	}
}

func (m *MinerInfo) Fields() map[string]any {
	return map[string]any{
		"power":                m.Power,
		"voltage":              m.Voltage,
		"current":              m.Current,
		"core_voltage":         m.CoreVoltage,
		"current_core_voltage": m.CoreVoltageActual,
		"frequency":            m.Frequency,
		"fan_speed_rpm":        m.FanSpeedRpm,
		"fan_speed":            m.Fanspeed,
		"temperature":          m.Temp,
		"hash_rate":            m.HashRate,
		"best_diff":            m.BestDiff,
		"best_session_diff":    m.BestSessionDiff,
		"free_heap":            m.FreeHeap,
		"shares_accepted":      m.SharesAccepted,
		"shares_rejected":      m.SharesRejected,
		"uptime":               m.UptimeSeconds,
	}
}
