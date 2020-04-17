package util

import (
	"net/http"

	"encoding/json"

	"github.com/LTitan/BloomFilter/pkg/logs"
)

var url string

func init() {
	url = "http://192.168.1.106:65221"
}

func GetHosts() interface{} {
	resp, err := http.Get(url + "/api/v1/host")
	if err != nil {
		logs.Logger.Errorf("[sync] [failed] error : %v", err)
		return nil
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logs.Logger.Errorf("[sync] [failed] error: response code is not 200")
		return nil
	}
	mp := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&mp); err != nil {
		logs.Logger.Errorf("[sync] [failed] error : %v", err)
		return nil
	}
	if _, ok := mp["data"]; ok {
		return mp["data"]
	}
	return nil
}
