package bitaxe

import (
	"fmt"
	"strconv"

	bitaxelib "github.com/mendelgusmao/bitaxe-telegraf-plugin/lib/bitaxe"
)

type bitaxeMetric bitaxelib.SystemInfo

func (m bitaxeMetric) Tags() map[string]string {
	return map[string]string{
		"hostname":              m.Hostname,
		"asic_model":            m.ASICModel,
		"stratum_url":           fmt.Sprintf("%s:%d", m.StratumURL, m.StratumPort),
		"stratum_user":          m.StratumUser,
		"os_version":            m.Version,
		"board_version":         m.BoardVersion,
		"auto_fan_speed":        strconv.Itoa(m.AutoFanSpeed),
		"overheat_mode":         strconv.Itoa(m.OverheatMode),
		"asic_count":            strconv.Itoa(m.ASICCount),
		"asic_small_core_count": strconv.Itoa(m.SmallCoreCount),
	}
}

func (m bitaxeMetric) Fields() map[string]any {
	return map[string]any{
		"power":                   m.Power,
		"voltage":                 m.Voltage,
		"current":                 m.Current,
		"core_voltage":            m.CoreVoltage,
		"current_core_voltage":    m.CoreVoltageActual,
		"frequency":               m.Frequency,
		"fan_speed_rpm":           m.FanSpeedRpm,
		"fan_speed":               m.FanSpeed,
		"temperature":             m.Temp,
		"hash_rate":               m.HashRate,
		"best_difficulty":         int64(*m.BestDiff),
		"best_session_difficulty": int64(*m.BestSessionDiff),
		"free_heap":               m.FreeHeap,
		"shares_accepted":         m.SharesAccepted,
		"shares_rejected":         m.SharesRejected,
		"uptime":                  m.UptimeSeconds,
	}
}
