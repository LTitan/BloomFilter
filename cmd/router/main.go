package main

import (
	"fmt"

	"github.com/LTitan/BloomFilter/internal/router"
	"github.com/LTitan/BloomFilter/pkg/config"
)

// @Title bloomfilter router
// @Version 1.0
// @Description This is a sample server Petstore server.

// @License.name Apache 2.0
// @License.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @Host 127.0.0.1:65221
// @BasePath /api/v1
func main() {
	port := fmt.Sprintf(":%v", config.Conf.Get("router.server_port"))
	if port == ":" {
		panic("slave port is null")
	}
	go router.RunServer()
	router.InitRouter(port)
}
