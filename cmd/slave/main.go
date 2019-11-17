package main

import (
	"github.com/LTitan/BloomFilter/internal/slave"
)

func main() {
	go slave.RunClient()
	slave.InitRouter()
}
