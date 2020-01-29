package handle

import (
	"net/http"

	appdata "github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/datastruct"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var cc colorController

func init() {
	cc.address = make(map[string]*datastruct.BloomFilter, 0)
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
	key := uuid.Must(uuid.NewV4(), nil)
	cc.address[key.String()] = datastruct.New(uint(req.Size * 1048576))
	app.Response(http.StatusOK, 20000, "ok", map[string]interface{}{"key": key.String()})
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
	app := response.APP{C: c}
	value := c.Query("value")
	if cc.address[key] == nil {
		app.Response(http.StatusOK, 20000, "ok", nil)
		return
	}
	app.Response(http.StatusOK, 20000, "ok", map[string]interface{}{
		"has": cc.address[key].Has(value),
	})
	return
}

// DeleteValue .
func DeleteValue(c *gin.Context) {
	value := c.Param("value")
	key := c.Param("key")
	app := response.APP{C: c}
	if cc.address[key] == nil {
		app.Response(http.StatusOK, 20000, "no more", nil)
		return
	}
	cc.address[key].Delete(value)
	app.Response(http.StatusOK, 20000, "success", map[string]interface{}{"value": value})
	return
}

// DeleteKey .
func DeleteKey(c *gin.Context) {
	key := c.Param("key")
	app := response.APP{C: c}
	if cc.address[key] == nil {
		app.Response(http.StatusOK, 20000, "no more", nil)
		return
	}
	cc.address[key].Release()
	cc.address[key] = nil
	app.Response(http.StatusOK, 20000, "success", map[string]interface{}{"key": key})
	return
}
