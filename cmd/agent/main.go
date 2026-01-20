package main

import (
	"encoding/json"
	"github.com/lilith/agent/internal/collector"

	"github.com/lilith/agent/internal/collector/model"

    "fmt"

)

func main() {
	payload := model.Payload{
		Agent:  collector.CollectAgentInfo(),
		CPU:  collector.CollectCPUInfo(),
		Disks:  collector.CollectDiskInfo(),
		Network: collector.CollectNetworkInfo(),
	}

	data, _ := json.MarshalIndent(payload, "", "  ")
	fmt.Println(string(data))
}
