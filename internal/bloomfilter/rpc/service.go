package rpc

import (
	"fmt"

	"github.com/LTitan/BloomFilter/pkg/datastruct"
	"github.com/LTitan/BloomFilter/pkg/logs"
	"github.com/mackerelio/go-osstat/memory"
)

type colorController struct {
	address map[string]*datastruct.BloomFilter
}

func getSystemAvalible(size uint64) error {
	m, err := memory.Get()
	if err != nil {
		return err
	}
	logs.Logger.Infof("currency memory size: %v", m.Free)
	if m.Free/1048576 > size {
		return nil
	}
	return fmt.Errorf("no enough memory")
}

func dumpToFile() {
	for key, value := range cc.address {
		value.Dump(key)
	}
}

func loadFromFile() {
	ret, err := datastruct.Load()
	if err != nil {
		logs.Logger.Fatalf("load from file err: %v", err)
		return
	}
	for key, value := range ret {
		cc.address[key] = value
	}
}
