package proxy

import (
	"time"

	slavec "github.com/LTitan/BloomFilter/internal/proxy/handler"
	"github.com/gin-gonic/gin"
)

var slave slavec.Slave

// RunServer .
func RunServer(port string) {
	go func() {
		tk := time.NewTicker(time.Second * 55)
		for {
			select {
			case <-tk.C:
				slave.ReceiveRouter()
			}
		}
	}()
	router := gin.Default()
	r := router.Group("/api/v1")
	{
		r.GET("/bloomfilter/query", slave.QueryValue)
		r.POST("/bloomfilter/query", slave.QueryMany)
		r.POST("/bloomfilter/add", slave.AddValues)
	}
	router.Run(port)
}
