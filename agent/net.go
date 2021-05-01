package agent

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/net"
	"time"
)

type NetStat struct {
	BytesSend       uint64 `json:"bytesSend"`
	BytesRecv       uint64 `json:"bytesRecv"`
	BytesSendPerSec uint64 `json:"bytesSendPerSec"`
	BytesRecvPerSec uint64 `json:"bytesRecvPerSec"`
}

func Net() (NetStat, error) {
	netIOInfoOlder, err := net.IOCounters(false)
	if err != nil {
		return NetStat{}, errors.Wrap(err, "Get old net io info failure.")
	}

	time.Sleep(time.Second)

	netIOInfoNewer, err := net.IOCounters(false)
	if err != nil {
		return NetStat{}, errors.Wrap(err, "Get new net io info failure.")
	}

	return NetStat{
		BytesSend:       netIOInfoNewer[0].BytesSent,
		BytesRecv:       netIOInfoNewer[0].BytesRecv,
		BytesSendPerSec: netIOInfoNewer[0].BytesSent - netIOInfoOlder[0].BytesSent,
		BytesRecvPerSec: netIOInfoNewer[0].BytesRecv - netIOInfoOlder[0].BytesRecv,
	}, nil
}

func (n NetStat) String() string {
	s, _ := json.Marshal(n)
	return string(s)
}
