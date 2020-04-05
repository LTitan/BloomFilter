package gateway

import (
	"net/http"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
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

// CreateUser .
func CreateUser(ctx *gin.Context) {
	app := response.APP{C: ctx}
	var user sqldata.UserInfo
	if err := ctx.BindJSON(&user); err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	if err := dao.CreateUser(&user); err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}

// QueryHasUser .
func QueryHasUser(ctx *gin.Context) {
	app := response.APP{C: ctx}
	var user sqldata.UserInfo
	if err := ctx.BindJSON(&user); err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	ret, err := dao.QueryHasUser(&user)
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}
