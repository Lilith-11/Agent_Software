package main

import (
	"encoding/json"
	"github.com/lilith/agent/internal/collector"

	"github.com/lilith/agent/internal/collector/model"

    "fmt"
	"os"

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
	file,err:=os.Create("payload.json")
	if err!=nil{
		fmt.Println("Error creating file:",err)
		return
	}
	defer file.Close()
	_,err=file.Write(data)
	if err!=nil{
		fmt.Println("Error writing to file:",err)
		return
	}
}