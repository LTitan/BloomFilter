package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
)

// FE .
type FE struct{}

// QueryCPUMemory .
func (*FE) QueryCPUMemory(ctx *gin.Context) {
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
func (*FE) CreateUser(ctx *gin.Context) {
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
func (*FE) QueryHasUser(ctx *gin.Context) {
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

// QueryApplyPagination .
func (*FE) QueryApplyPagination(ctx *gin.Context) {
	app := response.APP{C: ctx}
	pageSize := ctx.Query("page_size")
	currSize := ctx.Query("current_page")
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	cs, err := strconv.Atoi(currSize)
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	ret, page, err := dao.GetApplyRecordsPagination(ps, cs)
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.ResponseWithPage(http.StatusOK, 20000, "ok", ret, page)
	return
}

// UpdateApplyRecord .
func (*FE) UpdateApplyRecord(ctx *gin.Context) {
	app := response.APP{C: ctx}
	var recv UpdateFeild
	if err := ctx.BindJSON(&recv); err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	tm, err := time.Parse("2006-01-02 15:04:05", recv.ExpirationAt)
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	update := map[string]interface{}{
		"forecast_cap":  recv.Size,
		"expiration_at": tm,
	}
	if err := dao.UpdateApplyRecord([]string{recv.UUID}, update); err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}

// DeleteApplyRecord .
func (*FE) DeleteApplyRecord(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "-1")
	app := response.APP{C: ctx}
	if err := dao.DeleteApplyRecord(id); err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}

// GetAliveHosts .
func (*FE) GetAliveHosts(ctx *gin.Context) {
	ret, err := dao.GetAliveHosts()
	app := response.APP{C: ctx}
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}

// GetSingleAddressInfo .
func (*FE) GetSingleAddressInfo(ctx *gin.Context) {
	address := ctx.DefaultQuery("address", "0:0")
	app := response.APP{C: ctx}
	ret, err := dao.GetSingleAddressInfo(address)
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}

// RegisterDistribution .
func (*FE) RegisterDistribution(ctx *gin.Context) {
	app := response.APP{C: ctx}
	ret, err := dao.RegisterDistribution()
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}

// GetRegisterMemoryInfo .
func (*FE) GetRegisterMemoryInfo(ctx *gin.Context) {
	address := ctx.DefaultQuery("address", "0:0")
	app := response.APP{C: ctx}
	ret, err := dao.GetRegisterMemoryInfo(address)
	if err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", ret)
	return
}
