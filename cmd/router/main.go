package main

import (
	"fmt"

	"github.com/LTitan/BloomFilter/internal/router"
	"github.com/LTitan/BloomFilter/internal/router/gateway"
	"github.com/LTitan/BloomFilter/pkg/config"
)

func main() {
	port := fmt.Sprintf(":%v", config.Conf.Get("router.server_port"))
	if port == ":" {
		panic("slave port is null")
	}
	go router.RunServer()
	gateway.InitRouter(port)
}
