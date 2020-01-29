package main

import (
	"github.com/LTitan/BloomFilter/internal/bloomfilter"
)

func main() {
	go bloomfilter.RunClient()
	bloomfilter.InitRouter()
}
