package register

import (
	"time"

	"github.com/LTitan/BloomFilter/internal/router/dao"
	"github.com/LTitan/BloomFilter/pkg/logs"
)

func init() {
	go async()
}

func async() {
	tm := time.NewTimer(time.Minute)
	for {
		select {
		case <-tm.C:
			logs.Logger.Infof("async ticker begin ...")
			go dao.DeletedUpdateApplyRecord()
			now := time.Now()
			hosts, err := dao.GetRecentlyHost(&now, time.Minute*15)
			if err != nil {
				logs.Logger.Errorf("database async fail, error:", hosts)
			}
			for _, host := range hosts {
				if _, ok := globalMap.Load(host); !ok {
					globalMap.Store(host, true)
				}
			}
			tm.Reset(time.Minute * 20)
		}
	}
}
