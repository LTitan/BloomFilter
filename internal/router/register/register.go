package register

import (
	"sync"
	"time"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/logs"
)

var globalMap *sync.Map
var uuidMap *sync.Map

func init() {
	globalMap = new(sync.Map)
	uuidMap = new(sync.Map)
}
func SetHost(address string) {
	globalMap.Store(address, true)
}

func GetHosts() []string {
	var res []string
	globalMap.Range(func(k, v interface{}) bool {
		res = append(res, k.(string))
		return true
	})
	return res
}
func DelHost(address string) {
	globalMap.Delete(address)
	err := dao.UpdateStatusApplyAddress(address, sqldata.StatusTemporarilyUnavailable)
	if err != nil {
		logs.Logger.Errorf("update apply records status fail, error:", err)
	}
}

func Register(uuid string, address string, size uint64, expira time.Time) {
	var record sqldata.ApplyRecord
	record.ErrorRate = 0
	record.ForecastCap = size
	record.ExpirationAt = expira
	record.HostIP = address
	record.UUID = uuid
	record.Status = sqldata.StatusNormal
	if err := dao.CreatedApplyRecord(&record); err != nil {
		logs.Logger.Errorf("created apply records fail, error:", err)
	}
}
func GetRegistedHost(uuid string) (interface{}, bool) {
	return dao.QueryApplyAddress(uuid)
}
