package handler

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/LTitan/BloomFilter/internal/proxy/util"
	"github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Slave .
type Slave struct {
	Register *sync.Map
}

var clientTLSKey credentials.TransportCredentials

func init() {
	var err error
	clientTLSKey, err = credentials.NewClientTLSFromFile("/opt/server.pem", "bloomfilter")
	if err != nil {
		panic(err)
	}
}

// QueryValue .
// @Description query single value
// @Produce json
// @Tags Slave
// @Param key query string true "apply key"
// @Param value query string true "query value"
// @Success 200 {bool} bool has or not
// @Failure 400 {bool} bool has or not
// @Router /bloomfilter/query [get]
func (s *Slave) QueryValue(ctx *gin.Context) {
	key := ctx.DefaultQuery("key", "-1")
	value := ctx.DefaultQuery("value", "-1")
	bf := response.APP{C: ctx}
	address, found := s.getRegistedHost(key)
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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(clientTLSKey))
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
// @Description query single value
// @Produce json
// @Tags Slave
// @Param values body app.AddRequest true "add values"
// @Success 200 {bool} bool has or not
// @Failure 400 {bool} bool has or not
// @Router /bloomfilter/add [post]
func (s *Slave) AddValues(ctx *gin.Context) {
	var recv app.AddRequest
	bf := response.APP{C: ctx}
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	address, found := s.getRegistedHost(recv.Key)
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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(clientTLSKey))
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
// @Description query single value
// @Produce json
// @Tags Slave
// @Param values body app.AddRequest true "add values"
// @Success 200 {bool} bool has or not
// @Failure 400 {bool} bool has or not
// @Router /bloomfilter/query [post]
func (s *Slave) QueryMany(ctx *gin.Context) {
	var recv app.AddRequest
	bf := response.APP{C: ctx}
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	address, found := s.getRegistedHost(recv.Key)
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
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(clientTLSKey))
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err = client.QueryAll(ctx, &rpc.QueryManyRequest{Key: recv.Key, Values: recv.Strings})
	return
}

func (s *Slave) getRegistedHost(key string) (interface{}, bool) {
	return s.Register.Load(key)
}

// ReceiveRouter .
func (s *Slave) ReceiveRouter() {
	ret := util.GetHosts()
	if ret == nil {
		return
	}
	s.Register = nil
	s.Register = new(sync.Map)
	for key, value := range ret.(map[string]interface{}) {
		s.Register.Store(key, value)
	}
}
