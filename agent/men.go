package agent

import (
	"encoding/json"
	"github.com/pkg/errors"
	gopsMem "github.com/shirou/gopsutil/mem"
)

type MemStat struct {
	Total uint64 `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	SwapUsedPercent float64 `json:"swapUsedPercent"`
}

func Mem() (MemStat, error) {
	memInfo, err := gopsMem.VirtualMemory()
	if err != nil {
		return MemStat{}, errors.Wrap(err, "Get memory info failure.")
	}

	swapMemInfo, err := gopsMem.SwapMemory()
	if err != nil {
		return MemStat{}, errors.Wrap(err, "Get swap info failure.")
	}

	return MemStat{
		Total: memInfo.Total,
		Available:   memInfo.Available,
		Used:        memInfo.Used,
		UsedPercent: memInfo.UsedPercent,
		SwapUsedPercent: swapMemInfo.UsedPercent,
	}, nil
}

func (m MemStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
