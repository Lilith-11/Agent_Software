package collector

import (
	"fmt"
	"runtime"

	"github.com/google/uuid"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/lilith/agent/internal/collector/model"
)

func CollectAgentInfo() model.AgentInfo {
	h, _ := host.Info()
	m, _ := mem.VirtualMemory()

	return model.AgentInfo{
		AgentID:   uuid.New().String(),
		Hostname:   h.Hostname,
		OS:     fmt.Sprintf("%s %s", h.Platform, h.PlatformVersion),
		Platform:  runtime.GOOS,
		Architecture: runtime.GOARCH,
		BootTime:   h.BootTime,
		TotalRAMGB:   m.Total / (1024 * 1024 * 1024),
	}
}
