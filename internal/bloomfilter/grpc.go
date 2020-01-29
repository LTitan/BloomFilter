package bloomfilter

import (
	"context"
	"time"

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
				Cpu:         uint32(currCPU.CPUCount),
				Memory:      memoryInfo.Total / 1048576,
				CpuUsage:    float32(currCPU.System-preCPU.System) / float32(currCPU.Total-preCPU.Total) * 100,
				MemoryUsage: float32((memoryInfo.Used/1048576)) / float32((memoryInfo.Total/1048576)) * 100,
			})
			cancel()
			preCPU = currCPU
		}
	}
}
