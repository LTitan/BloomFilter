package register

import (
	"sync"
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

func Register(uuid string, address string) {
	uuidMap.Store(uuid, address)
}
func GetRegistedHost(uuid string) (string, bool) {
	return uuidMap.Load(uuid)
}
func DelRegisterHost(uuid string) {
	uuidMap.Delete(uuid)
}
