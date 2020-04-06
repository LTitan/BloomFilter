// Package bloomfilter internal package to achive CRUD handler
// this file to run self-server
package bloomfilter

import (
	"context"
	"net"
	"time"

	self "github.com/LTitan/BloomFilter/internal/bloomfilter/rpc"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/LTitan/BloomFilter/pkg/signal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunServer .
func RunServer(port string, _port uint32) {
	go self.DumpTricker()
	go signal.ExitBeautiful(func() {
		conn, _ := grpc.Dial(address, grpc.WithInsecure())
		c := rpc.NewGreeterClient(conn)
		ip := getLocalHost()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		_, err := c.CancelRegister(ctx, &rpc.MachineInfo{Host: ip, Port: _port})
		defer cancel()
		logs.Logger.Warnf("process will exit ...., cancel register send error: %v", err)
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
