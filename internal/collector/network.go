package collector

import (
	"io"
	"net/http"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/lilith/agent/internal/collector/model"
)

func getPublicIP() string {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	ip, _ := io.ReadAll(resp.Body)
	return string(ip)
}

func CollectNetworkInfo() model.NetworkInfo {
	ifaces, _ := net.Interfaces()
	var ips, macs []string

	for _, i := range ifaces {
		if i.HardwareAddr != "" {
			macs = append(macs, i.HardwareAddr)
		}
		for _, a := range i.Addrs {
			ips = append(ips, a.Addr)
		}
	}

	return model.NetworkInfo{
		LocalIPs: ips,
		MACs:     macs,
		PublicIP: getPublicIP(),
	}
}
