package slave

import (
	"github.com/LTitan/BloomFilter/internal/slave/handle"
	"github.com/gin-gonic/gin"
)

// InitRouter init routers
func InitRouter() {
	router := gin.Default()
	router.GET("/", handle.HelloWorld)
	router.Run(":65220")
}
