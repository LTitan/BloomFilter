package slave

import (
	"github.com/LTitan/BloomFilter/internal/slave/handle"
	"github.com/koding/kite"
)

var k *kite.Kite

func init() {
	k = kite.New("slave", "1.0.0")
}

// InitRouter init slave router
func InitRouter() {
	k.Config.Port = 65220
	k.HandleFunc("hello", handle.HelloWorld)
	k.Run()
}
