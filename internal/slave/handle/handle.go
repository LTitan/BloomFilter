package handle

import (
	"github.com/koding/kite"
)

// HelloWorld this is a hello world test
func HelloWorld(req *kite.Request) (interface{}, error) {
	return "hello world", nil
}
