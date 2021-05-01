package agent

import (
	"encoding/json"
	"github.com/pkg/errors"
	gopsDisk "github.com/shirou/gopsutil/disk"
)

type PartitionStat struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mountPoint"`
	FsType      string  `json:"fsType"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func Disk() ([]PartitionStat, error) {
	partitionsInfo, err := gopsDisk.Partitions(true)
	if err != nil {
		return nil, errors.Wrap(err, "Get disks' partitions info failure.")
	}

	partitionsStat := make([]PartitionStat, len(partitionsInfo))
	for _, partitionInfo := range partitionsInfo {
		partitionUsage, err := gopsDisk.Usage(partitionInfo.Mountpoint)
		if err != nil {
			return nil, errors.Wrap(err, "Get the usage of disks' partitions failure.")
		}

		partitionsStat = append(partitionsStat, PartitionStat{
			Device:      partitionInfo.Device,
			MountPoint:  partitionInfo.Mountpoint,
			FsType:      partitionInfo.Fstype,
			Total:       partitionUsage.Total,
			Free:        partitionUsage.Free,
			Used:        partitionUsage.Used,
			UsedPercent: partitionUsage.UsedPercent,
		})
	}

	return partitionsStat, nil
}

func (p PartitionStat) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}
