package main

import (
	"github.com/LTitan/BloomFilter/internal/router"
	"github.com/LTitan/BloomFilter/internal/router/gateway"
)

func main() {
	go router.RunServer()
	gateway.InitRouter()
}
