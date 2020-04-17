package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/register"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/response"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// Slave .
type Slave struct{}

// ApplyMemory .
// @Security ApiKeyAuth
// @Description apply bloomfilter memory
// @Accept  json
// @Produce json
// @Tags Slave
// @Param   allocSize body app.ApplyRequest true "alloc size"
// @Success 200 {object} ApplyRes {"code":2000,"data":null,"message":""}
// @Failure 400 {object} ApplyRes {"code":4000,"data":null,"message":""}
// @Router /bloomfilter/apply [post]
func (*Slave) ApplyMemory(ctx *gin.Context) {
	bf := response.APP{C: ctx}
	var recv app.ApplyRequest
	if err := ctx.BindJSON(&recv); err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	expira, err := time.Parse("2006-01-02 15:04:05", recv.Expiration)
	if err != nil {
		bf.Response(http.StatusBadRequest, response.ParamsError, "params error", nil)
		return
	}
	resp, err := applyHandler(recv.Size, expira)
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", &ApplyRes{Yes: resp.GetRecv(), Key: resp.GetKey()})
	return
}

func applyHandler(size uint64, expira time.Time) (recv *rpc.ApplyReply, err error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	host := register.GetHosts()
	logs.Logger.Infof("all register hosts: %v", host)
	var address string
	if len(host) == 0 {
		return nil, fmt.Errorf("no register hosts")
	}
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
		register.Register(recv.GetKey(), address, size, expira)
		// timeout to delete
		go func(addr, uuid string) {
			now := time.Now()
			tm := time.NewTimer(expira.Sub(now))
			<-tm.C
			deleteHandler(addr, uuid)
			logs.Logger.Warnf("from %s, this key %s had been deleted.", addr, uuid)
		}(address, recv.GetKey())
	}
	return
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
func (*Slave) QueryValue(ctx *gin.Context) {
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
// @Description query single value
// @Produce json
// @Tags Slave
// @Param values body app.AddRequest true "add values"
// @Success 200 {bool} bool has or not
// @Failure 400 {bool} bool has or not
// @Router /bloomfilter/add [post]
func (*Slave) AddValues(ctx *gin.Context) {
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
// @Description query single value
// @Produce json
// @Tags Slave
// @Param values body app.AddRequest true "add values"
// @Success 200 {bool} bool has or not
// @Failure 400 {bool} bool has or not
// @Router /bloomfilter/query [post]
func (*Slave) QueryMany(ctx *gin.Context) {
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

// DeletKey .
// @Security ApiKeyAuth
// @Description delete key
// @Produce json
// @Tags Slave
// @Param uuid path string true "key(uuid)"
// @Success 200 {boolean} bool has or not
// @Failure 400 {boolean} bool has or not
// @Router /bloomfilter/del/{uuid} [delete]
func (*Slave) DeletKey(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	address, found := register.GetRegistedHost(uuid)
	bf := response.APP{C: ctx}
	if !found {
		bf.Response(http.StatusNotFound, response.ParamsError, "not found key", nil)
		return
	}
	if err := deleteHandler(address.(string), uuid); err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", nil)
	return
}

func deleteHandler(address, uuid string) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = client.Delete(ctx, &rpc.DeleteRequest{Key: uuid})
	if err == nil {
		go func() {
			dao.UpdateApplyRecord([]string{uuid}, map[string]interface{}{"status": sqldata.StatusRelease})
		}()
	}
	return
}

// BackupSlave .
// @Security ApiKeyAuth
// @Description backup address
// @Produce json
// @Tags Slave
// @Param address path string true "address(ip:port)"
// @Success 200 {boolean} bool has or not
// @Failure 400 {boolean} bool has or not
// @Router /bloomfilter/{address} [put]
func (*Slave) BackupSlave(ctx *gin.Context) {
	bf := response.APP{C: ctx}
	address := ctx.Param("address")
	uuids, err := backupHandler(address)
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	err = dao.UpdateApplyRecord(uuids, map[string]interface{}{"status": sqldata.StatusNormal})
	if err != nil {
		bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
		return
	}
	bf.Response(http.StatusOK, response.OK, "ok", nil)
	return
}

func backupHandler(address string) ([]string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := rpc.NewSlaveServerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	recv, err := client.DoBackup(ctx, &rpc.Reply{Recv: true})
	return recv.GetKeys(), err
}

//ReadUploadFile .
func (*Slave) ReadUploadFile(ctx *gin.Context) {
	fileName := ctx.Query("file_name")
	status := ctx.Query("status")
	key := ctx.Query("key")
	bf := response.APP{C: ctx}
	address, found := register.GetRegistedHost(key)
	if !found {
		bf.Response(http.StatusNotFound, response.ParamsError, "not found key", nil)
		return
	}
	if status == "1" {
		data, err := ioutil.ReadFile("/tmp/" + fileName)
		if err != nil {
			bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
			return
		}
		strs := strings.Split(string(data), "\n")
		if err = addHandler(&app.AddRequest{Key: key, Strings: strs}, address.(string)); err != nil {
			bf.Response(http.StatusBadGateway, response.ServeError, err.Error(), nil)
			return
		}
	}
	bf.Response(http.StatusOK, response.OK, "ok", nil)
	return
}
