package gateway

import (
	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// QueryCPUMemory .
func QueryCPUMemory(ctx *gin.Context) {
	app := response.APP{C: ctx}
	ret, err := dao.QueryHostHardwareInfo()
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}
