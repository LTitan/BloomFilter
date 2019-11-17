package master

import (
	context "context"
	"fmt"
	"net"

	"github.com/LTitan/BloomFilter/pkg/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) HeartBeat(ctx context.Context, req *rpc.MachineInfo) (*rpc.Reply, error) {
	fmt.Printf("cpu:%v,memory:%v\n", req.GetCpu(), req.GetMemory())
	return &rpc.Reply{Recv: true}, nil
}

// RunServer .
func RunServer() {
	lis, err := net.Listen("tcp", ":50051")
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
