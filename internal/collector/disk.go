package collector

import (
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/lilith/agent/internal/collector/model"
)

func bytesToGB(b uint64) float64 {
	return float64(b) / (1024 * 1024 * 1024)
}

func CollectDiskInfo() []model.DiskInfo {
	partitions, _ := disk.Partitions(false)
	var disks []model.DiskInfo

	for _, p := range partitions {
		u, err := disk.Usage(p.Mountpoint)
		if err != nil {
			continue
		}

		disks = append(disks, model.DiskInfo{
			Device:  p.Device,
			FSType:  p.Fstype,
			TotalGB: bytesToGB(u.Total),
			UsedGB:  bytesToGB(u.Used),
			FreeGB:  bytesToGB(u.Free),
			Percent: u.UsedPercent,
		})
	}
	return disks
}
