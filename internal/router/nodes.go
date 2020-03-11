package router

import (
	"github.com/LTitan/BloomFilter/pkg/queue"
)

// IPPair .
type IPPair struct {
	IP      string
	MemFree float64
}

// Less .
func (a *IPPair) Less(b interface{}) bool {
	tmp := b.(IPPair)
	return a.MemFree < tmp.MemFree
}

var globalQueue *queue.PriorityQueue

func init() {
	globalQueue = queue.NewPriorityQueue()
}

func loadFromDB() {
	
}
