package gateway

import (
	"github.com/gin-gonic/gin"
)

// InitRouter .
func InitRouter(){
	router := gin.Default()
	r := router.Group("/api/v1")
	{
		r.GET("/host/info", QueryCPUMemory)
	}
	router.Run(":65221")
}