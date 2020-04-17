package router

import (
	context "context"
	"fmt"
	"net"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/register"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/config"
	"github.com/LTitan/BloomFilter/pkg/logs"
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
	record.Port = req.GetPort()
	logs.Logger.Infof("from host: %s:%d, cpu number:%d, memory size:%d, cpu used:%.2f%%, memory used:%.2f%%", record.HostIP, record.Port, record.CPUNum, record.MemCap, record.CPUUsage, record.MemUsage)
	go func() {
		register.SetHost(fmt.Sprintf("%v:%v", record.HostIP, record.Port))
	}()
	_ = dao.CreatedOneRecord(&record)
	return &rpc.Reply{Recv: true}, nil
}

func (s *server) CancelRegister(ctx context.Context, req *rpc.MachineInfo) (*rpc.Reply, error) {
	go register.DelHost(fmt.Sprintf("%v:%v", req.GetHost(), req.GetPort()))
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
