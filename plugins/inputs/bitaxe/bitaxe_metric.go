package bitaxe

import (
	"fmt"
	"strconv"

	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
)

type bitaxemetric bitaxelib.MinerInfo

func (m bitaxemetric) Tags() map[string]string {
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

func (m bitaxemetric) Fields() map[string]any {
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
