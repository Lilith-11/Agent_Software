package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/lilith/agent/internal/collector/model"
)

func CollectCPUInfo() model.CPUInfo {
	info, _ := cpu.Info()
	usagePerCore, _ := cpu.Percent(time.Second, true)
	overall, _ := cpu.Percent(time.Second, false)
	physical, _ := cpu.Counts(false)
	logical, _ := cpu.Counts(true)

	return model.CPUInfo{
		Manufacturer:      info[0].VendorID,
		Model:             info[0].ModelName,
		CPUSpeedMHz:       info[0].Mhz,
		Cores:             int32(physical),
		LogicalProcessors: logical,
		HyperThreading:    logical > int(physical),
		UsagePerCore:      usagePerCore,
		OverallCPUUsage:   overall[0],
	}
}
