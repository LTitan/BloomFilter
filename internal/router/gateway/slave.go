package gateway

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/LTitan/BloomFilter/internal/router/register"
	"github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// ApplyMemory godoc
// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func ApplyMemory(ctx *gin.Context) {
	bf := response.APP{C: ctx}
	var recv app.ApplyRequest
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	resp, err := applyHandler(recv.Size)
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", map[string]interface{}{"yes": resp.GetRecv(), "key": resp.GetKey()})
	return
}

func applyHandler(size uint64) (recv *rpc.ApplyReply, err error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	host := register.GetHosts()
	logs.Logger.Infof("all register hosts: %v", host)
	var address string
	if len(host) == 1 {
		address = host[0]
	} else {
		address = host[r.Intn(len(host))]
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	recv, err = client.Apply(ctx, &rpc.ApplyRequest{Size: size})
	if err != nil {
		return
	}
	if recv.GetKey() != "" {
		register.Register(recv.GetKey(), address)
	}
	return
}

// QueryValue .
func QueryValue(ctx *gin.Context) {
	key := ctx.DefaultQuery("key", "-1")
	value := ctx.DefaultQuery("value", "-1")
	bf := response.APP{C: ctx}
	address, found := register.GetRegistedHost(key)
	if !found {
		bf.Response(http.StatusNotFound, response.ParamsError, "not found key", nil)
		return
	}
	res, err := querySingleHandler(key, value, address.(string))
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", res)
	return
}

func querySingleHandler(key, value, address string) (res bool, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	recv, err := client.QuerySingle(ctx, &rpc.QueryRequest{Key: key, Value: value})
	if err != nil {
		return
	}
	return recv.GetRecv(), nil
}

//AddValues .
func AddValues(ctx *gin.Context) {
	var recv app.AddRequest
	bf := response.APP{C: ctx}
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	address, found := register.GetRegistedHost(recv.Key)
	if !found {
		bf.Response(http.StatusNotFound, response.ParamsError, "not found key", nil)
		return
	}
	if err := addHandler(&recv, address.(string)); err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", nil)
	return
}

func addHandler(recv *app.AddRequest, address string) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = client.Add(ctx, &rpc.AddRequest{Key: recv.Key, Values: recv.Strings})
	if err != nil {
		return
	}
	return nil
}

// QueryMany .
// TODO:
func QueryMany(ctx *gin.Context) {
	var recv app.AddRequest
	bf := response.APP{C: ctx}
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	address, found := register.GetRegistedHost(recv.Key)
	if !found {
		bf.Response(http.StatusNotFound, response.ParamsError, "not found key", nil)
		return
	}
	res, err := queryManyHandler(&recv, address.(string))
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", map[string]interface{}{"finish": res.GetHas(), "results": res.GetResult()})
	return
}

func queryManyHandler(recv *app.AddRequest, address string) (res *rpc.QueryManyReply, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err = client.QueryAll(ctx, &rpc.QueryManyRequest{Key: recv.Key, Values: recv.Strings})
	return
}
