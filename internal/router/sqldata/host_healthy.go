package sqldata

import (
	"github.com/jinzhu/gorm"
)

type (
	// HostHealthy .
	HostHealthy struct {
		*gorm.Model
		HostIP   string
		CPUNum   int
		MemCap   int
		CPUUsage float32
		MemUsage float32
	}
)

// TableName .
func (h *HostHealthy)TableName()(string){
	return "host_healthy"
}