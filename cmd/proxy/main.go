package main

import (
	"fmt"
	"os"

	"github.com/LTitan/BloomFilter/internal/proxy"
	"github.com/spf13/pflag"
)

var port int

func init() {
	pflag.IntVar(&port, "port", 7361, "port")
}

func main() {
	pflag.Parse()
	p := os.Getenv("PROXY_PORT")
	if p != "" {
		p = ":" + p
	} else {
		p = fmt.Sprintf(":%d", port)
	}
	proxy.RunServer(p)
}
