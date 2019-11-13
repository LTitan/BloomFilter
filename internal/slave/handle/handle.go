package handle

import (
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorld test
func HelloWorld(c *gin.Context) {
	app := response.APP{C: c}
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}
