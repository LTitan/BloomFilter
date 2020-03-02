package bloomfilter

import (
	"context"
	"fmt"
	"time"

	"github.com/LTitan/BloomFilter/pkg/config"
	"github.com/LTitan/BloomFilter/pkg/nets"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"google.golang.org/grpc"
)

var (
	address, network string
)

const (
	memBase       = 1048576
	cpuTickerTime = time.Second * 5
	heartBeatTime = time.Minute * 3
)

func init() {
	host := config.Conf.Get("router.host")
	port := config.Conf.Get("router.grpc_port")
	address = fmt.Sprintf("%v:%v", host, port)
	network = config.Conf.Get("bloomfilter.network").(string)
}

// RunClient .
func RunClient() {
	preCPU, _ := cpu.Get()
	tm := time.NewTicker(heartBeatTime)
	cpuTicker := time.NewTicker(cpuTickerTime)
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
				Host:        ip,
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
	ips, err := nets.GetIPv4ByInterface(network)
	if err != nil {
		return "unkonwn"
	}
	if len(ips) == 0 {
		return "unkonwn"
	}
	return ips[0]
}
