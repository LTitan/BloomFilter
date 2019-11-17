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
	tm := time.NewTicker(time.Minute * 3)
	for {
		select {
		case <-tm.C:
			conn, err := grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			c := rpc.NewGreeterClient(conn)
			cpuInfo, _ := cpu.Get()
			memoryInfo, _ := memory.Get()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			_, _ = c.HeartBeat(ctx, &rpc.MachineInfo{
				Cpu:    int32(cpuInfo.Total),
				Memory: int32(memoryInfo.Total),
			})
			cancel()
		}
	}
}
