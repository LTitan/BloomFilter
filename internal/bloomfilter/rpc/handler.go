package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/LTitan/BloomFilter/pkg/datastruct"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/LTitan/BloomFilter/pkg/rpc"
	uuid "github.com/satori/go.uuid"
)

type Slave struct{}

var cc colorController

func init() {
	cc.address = make(map[string]*datastruct.BloomFilter, 0)
}

func (*Slave) Apply(ctx context.Context, ar *rpc.ApplyRequest) (*rpc.ApplyReply, error) {
	response := new(rpc.ApplyReply)
	response.Recv = false
	if err := getSystemAvalible(ar.GetSize()); err != nil {
		logs.Logger.Errorf("alloc bloomfilter fail, error: %v", err)
		return response, err
	}
	key := uuid.Must(uuid.NewV4(), nil)
	response.Key = key.String()
	cc.address[response.Key] = datastruct.New(uint(ar.GetSize() * 1048576))
	response.Recv = true
	return response, nil
}

func (*Slave) Add(ctx context.Context, ar *rpc.AddRequest) (*rpc.Reply, error) {
	response := new(rpc.Reply)
	key := ar.GetKey()
	if _, ok := cc.address[key]; !ok {
		logs.Logger.Warnf("not found uuid key is %v", ar.GetKey())
		response.Recv = false
		err := fmt.Errorf("not found uuid key")
		return response, err
	}
	str := ar.GetValues()
	logs.Logger.Infof("recv add strings are: %v", str)
	for i := range str {
		cc.address[key].Add(str[i])
	}
	response.Recv = true
	return response, nil
}

func (*Slave) Delete(ctx context.Context, rd *rpc.DeleteRequest) (*rpc.Reply, error) {
	key := rd.GetKey()
	response := new(rpc.Reply)
	if _, ok := cc.address[key]; !ok {
		response.Recv = false
		err := fmt.Errorf("not found uuid key")
		return response, err
	}
	cc.address[key].Release()
	cc.address[key] = nil
	delete(cc.address, key)
	response.Recv = true
	return response, nil
}

func (*Slave) QuerySingle(ctx context.Context, rq *rpc.QueryRequest) (*rpc.Reply, error) {
	key := rq.GetKey()
	response := new(rpc.Reply)
	if _, ok := cc.address[key]; !ok {
		response.Recv = false
		err := fmt.Errorf("not found uuid key")
		return response, err
	}
	value := rq.GetValue()
	logs.Logger.Infof("query single value key: %v, value : %v", key, value)
	response.Recv = cc.address[key].Has(value)
	return response, nil
}

func (*Slave) QueryAll(ctx context.Context, rq *rpc.QueryManyRequest) (*rpc.QueryManyReply, error) {
	response := new(rpc.QueryManyReply)
	key := rq.GetKey()
	if _, ok := cc.address[key]; !ok {
		response.Has = false
		err := fmt.Errorf("not found uuid key")
		return response, err
	}
	values := rq.GetValues()
	for i := range values {
		response.Result = append(response.Result, cc.address[key].Has(values[i]))
	}
	response.Has = true
	return response, nil
}

// DumpTricker .
func DumpTricker() {
	tm := time.NewTicker(time.Minute * 30)
	for {
		<-tm.C
		logs.Logger.Info("self dump backup will begin")
		dumpToFile()
		logs.Logger.Info("self dump backup success")
	}
}
