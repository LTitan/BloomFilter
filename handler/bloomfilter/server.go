// Package bloomfilter internal package to achive CRUD handler
// this file to run self-server
package bloomfilter

import (
	"net"

	self "github.com/LTitan/BloomFilter/internal/bloomfilter/rpc"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/LTitan/BloomFilter/pkg/signal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunServer .
func RunServer(port string) {
	go self.DumpTricker()
	go signal.ExitBeautiful(func() {
		logs.Logger.Warnf("process will exit ....")
		defer logs.Logger.Sync()
	})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logs.Logger.Fatalf("failed to listen(%v), error %v", port, err)
		panic("failed to listen:" + err.Error())
	}
	s := grpc.NewServer()
	rpc.RegisterSlaveServerServer(s, &self.Slave{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		logs.Logger.Fatalf("failed to serve, error %v", err)
		panic("failed to serve:" + err.Error())
	}
}
