package slave

import (
	"context"
	"time"

	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"google.golang.org/grpc"
)

const (
	address = "10.2.214.61:50051"
)

// RunClient .
func RunClient() {
	preCPU, _ := cpu.Get()
	tm := time.NewTicker(time.Minute * 3)
	for {
		select {
		case <-tm.C:
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			c := rpc.NewGreeterClient(conn)
			currCPU, _ := cpu.Get()
			memoryInfo, _ := memory.Get()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			_, _ = c.HeartBeat(ctx, &rpc.MachineInfo{
				Cpu:         int32(currCPU.CPUCount),
				Memory:      int32(memoryInfo.Total / 1048576),
				CpuUsage:    float32(currCPU.System-preCPU.System) / float32(currCPU.Total-preCPU.Total) * 100,
				MemoryUsage: float32(memoryInfo.Used/memoryInfo.Total) * 100,
			})
			cancel()
			preCPU = currCPU
		}
	}
}
