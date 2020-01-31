package bloomfilter

import (
	"context"
	"time"

	"github.com/LTitan/BloomFilter/pkg/nets"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:50051"
	memBase = 1048576
)

// RunClient .
func RunClient() {
	preCPU, _ := cpu.Get()
	tm := time.NewTicker(time.Minute * 3)
	cpuTicker := time.NewTicker(time.Second * 5)
	var sumCPUUsage, cnt float32
	for {
		select {
		case <-tm.C:
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			c := rpc.NewGreeterClient(conn)
			memoryInfo, _ := memory.Get()
			ip := getLocalHost()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			_, _ = c.HeartBeat(ctx, &rpc.MachineInfo{
				Cpu:         uint32(preCPU.CPUCount),
				Memory:      memoryInfo.Total / memBase,
				CpuUsage:    sumCPUUsage / cnt,
				MemoryUsage: float32((memoryInfo.Used / memBase)) / float32((memoryInfo.Total / memBase)) * 100,
				Host: ip,
			})
			cancel()
			cnt = 0
			sumCPUUsage = 0
		case <-cpuTicker.C:
			currCPU, _ := cpu.Get()
			sumCPUUsage += (float32(currCPU.System-preCPU.System) / float32(currCPU.Total-preCPU.Total) * 100)
			preCPU = currCPU
			cnt++
		}
	}
}

func getLocalHost() string {
	ips, err := nets.GetIPv4ByInterface("eth2")
	if err != nil {
		return "unkonwn"
	}
	if len(ips) == 0 {
		return "unkonwn"
	}
	return ips[0]
}
