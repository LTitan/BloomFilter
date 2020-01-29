package handle

import (
	"fmt"
	"log"

	appdata "github.com/LTitan/BloomFilter/pkg/app"
	"github.com/LTitan/BloomFilter/pkg/datastruct"
	"github.com/mackerelio/go-osstat/memory"
)

type colorController struct {
	address map[string]*datastruct.BloomFilter
}

func addValues(req *appdata.AddRequest) {
	for _, each := range req.Strings {
		cc.address[req.Key].Add(each)
	}
}

func getSystemAvalible(size uint64) error {
	m, err := memory.Get()
	if err != nil {
		return err
	}
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
		log.Fatalln(err)
		return
	}
	for key, value := range ret {
		cc.address[key] = value
	}
}
