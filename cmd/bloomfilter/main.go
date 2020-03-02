package main

import (
	"fmt"

	"github.com/LTitan/BloomFilter/internal/bloomfilter"
	"github.com/LTitan/BloomFilter/pkg/config"
)

func main() {
	port := fmt.Sprintf(":%v", config.Conf.Get("bloomfilter.port"))
	if port == ":" {
		panic("slave port is null")
	}
	go bloomfilter.RunClient()

	bloomfilter.InitRouter(port)
}
