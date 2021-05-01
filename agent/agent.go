package agent

import (
	"encoding/json"
	"github.com/pengcheng789/dumpus/log"
)

var logger = log.Logger

type InfoStat struct {
	Cpu        CpuStat         `json:"cpu"`
	Mem        MemStat         `json:"mem"`
	Net        NetStat         `json:"net"`
	Partitions []PartitionStat `json:"partitions"`
	Host       HostStat        `json:"host"`
}

func Info() InfoStat {
	cpu, err := Cpu()
	if err != nil {
		logger.WithError(err).Error()
	}

	mem, err := Mem()
	if err != nil {
		logger.WithError(err).Error()
	}

	net, err := Net()
	if err != nil {
		logger.WithError(err).Error()
	}

	Partitions, err := Disk()
	if err != nil {
		logger.WithError(err).Error()
	}

	host, err := Host()
	if err != nil {
		logger.WithError(err).Error()
	}

	return InfoStat{
		Cpu:        cpu,
		Mem:        mem,
		Net:        net,
		Partitions: Partitions,
		Host:       host,
	}
}

func (i InfoStat) String() string {
	s, _ := json.Marshal(i)
	return string(s)
}
