package model

type AgentInfo struct {
	AgentID  string `json:"agent_id"`
	Hostname  string `json:"hostname"`
	OS      string `json:"operating_system"`
	Platform     string `json:"platform"`
	Architecture string `json:"architecture"`
	BootTime     uint64 `json:"boot_time"`
	TotalRAMGB   uint64 `json:"total_ram"`
}

type CPUInfo struct {
	Manufacturer   string    `json:"manufacturer"`
	Model   string    `json:"model"`
	CPUSpeedMHz  float64   `json:"cpu_speed_mhz"`
	Cores  int32     `json:"cores"`
	LogicalProcessors int       `json:"logical_processors"`
	HyperThreading bool      `json:"hyperthread"`
	UsagePerCore   []float64 `json:"usage_per_core"`
	OverallCPUUsage   float64   `json:"overall_usage"`
}

type DiskInfo struct {
	Device string  `json:"device"`
	FSType string  `json:"fstype"`
	TotalGB float64 `json:"total"`
	UsedGB float64 `json:"used"`
	FreeGB float64 `json:"free"`
	Percent float64 `json:"percent"`
}

type NetworkInfo struct {
	LocalIPs []string `json:"local_ips"`
	MACs     []string `json:"mac_addresses"`
	PublicIP string   `json:"public_ip"`
}

type Payload struct {
	Agent  AgentInfo   `json:"agent-agentinfo"`
	CPU   CPUInfo     `json:"agent-cpu"`
	Disks []DiskInfo  `json:"agent-disks"`
	Network NetworkInfo `json:"agent-network"`
}
