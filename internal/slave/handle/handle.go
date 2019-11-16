package handle

import (
	appdata "github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/datastruct"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
)

var cc colorController

func init() {
	cc.cnt = 1
	cc.address = make(map[uint32]*datastruct.BloomFilter, 0)
}

// HelloWorld test
func HelloWorld(c *gin.Context) {
	app := response.APP{C: c}
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}

// ApplyMemory .
func ApplyMemory(c *gin.Context) {
	app := response.APP{C: c}
	var key uint32
	var err error
	var req appdata.ApplyRequest
	if err = c.BindJSON(&req); err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	if err = getSystemAvalible(req.Size); err != nil {
		app.Response(http.StatusBadGateway, 50000, err.Error(), nil)
		return
	}
	key = cc.cnt
	atomic.AddUint32(&cc.cnt, 1)
	cc.address[key] = datastruct.New(uint(req.Size * 1048576))
	app.Response(http.StatusOK, 20000, "ok", map[string]interface{}{"key": key})
	return
}

// AddHandle .
func AddHandle(c *gin.Context) {
	var err error
	app := response.APP{C: c}
	var req appdata.AddRequest
	if err = c.BindJSON(&req); err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	if cc.address[req.Key] == nil {
		app.Response(http.StatusBadGateway, 50000, "not found key", nil)
		return
	}
	addValues(&req)
	app.Response(http.StatusOK, 20000, "ok", nil)
	return
}

// QueryValue .
func QueryValue(c *gin.Context) {
	key := c.Query("key")
	k, err := strconv.Atoi(key)
	app := response.APP{C: c}
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	value := c.Query("value")
	kk := uint32(k)
	if cc.address[kk] == nil {
		app.Response(http.StatusOK, 20000, "ok", nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", map[string]interface{}{
		"has": cc.address[kk].Has(value),
	})
	return
}

// DeleteValue .
func DeleteValue(c *gin.Context) {
	value := c.Param("value")
	key := c.Param("key")
	k, err := strconv.Atoi(key)
	app := response.APP{C: c}
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	kk := uint32(k)
	if cc.address[kk] == nil {
		app.Response(http.StatusOK, 20000, "no more", nil)
		return
	}
	cc.address[kk].Delete(value)
	app.Response(http.StatusOK, 20000, "success", map[string]interface{}{"value": value})
	return
}

// DeleteKey .
func DeleteKey(c *gin.Context) {
	key := c.Param("key")
	k, err := strconv.Atoi(key)
	app := response.APP{C: c}
	if err != nil {
		app.Response(http.StatusBadRequest, 40000, err.Error(), nil)
		return
	}
	kk := uint32(k)
	if cc.address[kk] == nil {
		app.Response(http.StatusOK, 20000, "no more", nil)
		return
	}
	cc.address[kk].Release()
	app.Response(http.StatusOK, 20000, "success", map[string]interface{}{"key": key})
	return
}
