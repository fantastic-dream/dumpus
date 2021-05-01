package agent

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	gopsCpu "github.com/shirou/gopsutil/cpu"
	"math"
	"time"
)

type CpuInfoStat struct {
	ModelName string `json:"modelName"`
	Cores     int32  `json:"cores"`
}

func CpuInfo() (CpuInfoStat, error) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second * math.MaxUint16)

	c, err := gopsCpu.InfoWithContext(ctx)
	if err != nil {
		return CpuInfoStat{}, errors.Wrap(err, "Get cpu info failure.")
	}

	return CpuInfoStat{
		ModelName: c[0].ModelName,
		Cores:     c[0].Cores,
	}, nil
}

func (c CpuInfoStat) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

// Cpu info and performance struct.
type CpuStat struct {
	CpuInfo       CpuInfoStat `json:"cpuInfo"`
	CpuPercent    float64     `json:"cpuPercent"`
	CpuPerPercent []float64   `json:"cpuPerPercent"`
}

func Cpu() (CpuStat, error) {
	p, err := gopsCpu.Percent(time.Second, false)
	if err != nil {
		return CpuStat{}, errors.Wrap(err, "Get cpu usage percent failure.")
	}

	cpuInfo, err := CpuInfo()
	if err != nil {
		return CpuStat{}, err
	}

	return CpuStat{
		CpuInfo:    cpuInfo,
		CpuPercent: p[0],
	}, nil
}

func (c CpuStat) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}
