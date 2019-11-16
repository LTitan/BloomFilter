package slave

import (
	"github.com/LTitan/BloomFilter/internal/slave/handle"
	"github.com/gin-gonic/gin"
)

// InitRouter init routers
func InitRouter() {
	router := gin.Default()
	router.GET("/", handle.HelloWorld)
	r := router.Group("/api/v1")
	{
		r.POST("/add", handle.AddHandle)
		r.POST("/apply", handle.ApplyMemory)
		r.GET("/query", handle.QueryValue)
		r.DELETE("/delete/:key/:value", handle.DeleteValue)
		r.DELETE("/delete/:key", handle.DeleteKey)
	}
	router.Run(":65220")
}