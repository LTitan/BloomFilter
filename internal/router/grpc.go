package router

import (
	context "context"
	"fmt"
	"net"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/config"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port string

func init() {
	port = fmt.Sprintf(":%v", config.Conf.Get("router.grpc_port"))
}

type server struct{}

func (s *server) HeartBeat(ctx context.Context, req *rpc.MachineInfo) (*rpc.Reply, error) {
	var record sqldata.HostHealthy
	record.HostIP = req.GetHost()
	record.CPUNum = int(req.GetCpu())
	record.MemCap = int(req.GetMemory())
	record.MemUsage = req.GetMemoryUsage()
	record.CPUUsage = req.GetCpuUsage()
	fmt.Printf("from host: %s\n", record.HostIP)
	fmt.Printf("cpu num: %v,memory cap:%v\n", record.CPUNum, record.MemCap)
	fmt.Printf("curr cpu used: %.2f%%, memory used:%.2f%%\n", record.CPUUsage, record.MemUsage)
	_ = dao.CreatedOneRecord(&record)
	return &rpc.Reply{Recv: true}, nil
}

// RunServer .
func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	s := grpc.NewServer()
	rpc.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		panic("failed to serve:" + err.Error())
	}
}
