package agent

import (
	"encoding/json"
	"github.com/pkg/errors"
	gopsHost "github.com/shirou/gopsutil/host"
	gopsNet "github.com/shirou/gopsutil/net"
	"os"
	"path/filepath"
)

type NetInterface struct {
	Name    string   `json:"name"`
	IpAddrs []string `json:"ipAddrs"`
}

type HostStat struct {
	Hostname      string         `json:"hostname"`
	BootTime      uint64         `json:"bootTime"`
	UpTime        uint64         `json:"upTime"`
	ProcessNum    uint64         `json:"processNum"`
	Os            string         `json:"os"`
	Platform      string         `json:"platform"`
	PlatVersion   string         `json:"platVersion"`
	NetInterfaces []NetInterface `json:"netInterfaces"`
	AgentPath     string         `json:"agentPath"`
}

func Host() (HostStat, error) {
	hostInfo, err := gopsHost.Info()
	if err != nil {
		return HostStat{}, errors.Wrap(err, "Get host info failure.")
	}

	netInterfacesInfo, err := gopsNet.Interfaces()
	if err != nil {
		return HostStat{}, errors.Wrap(err, "Get net interface info failure.")
	}
	var netInterfaces []NetInterface
	for _, netInterfaceInfo := range netInterfacesInfo {
		var ipAddrs []string
		for _, addr := range netInterfaceInfo.Addrs {
			ipAddrs = append(ipAddrs, addr.Addr)
		}
		netInterfaces = append(netInterfaces, NetInterface{
			Name:    netInterfaceInfo.Name,
			IpAddrs: ipAddrs,
		})
	}

	ex, err := os.Executable()
	if err != nil {
		return HostStat{}, errors.Wrap(err, "Get os executable failure.")
	}
	exPath := filepath.Dir(ex)

	return HostStat{
		Hostname:      hostInfo.Hostname,
		BootTime:      hostInfo.BootTime,
		UpTime:        hostInfo.Uptime,
		ProcessNum:    hostInfo.Procs,
		Os:            hostInfo.OS,
		Platform:      hostInfo.Platform,
		PlatVersion:   hostInfo.PlatformVersion,
		NetInterfaces: netInterfaces,
		AgentPath:     exPath,
	}, nil
}

func (h HostStat) String() string {
	s, _ := json.Marshal(h)
	return string(s)
}
