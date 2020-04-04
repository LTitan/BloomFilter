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
		Port     uint32
	}
	// UserInfo .
	UserInfo struct {
		*gorm.Model
		Name     string `json:"name"`
		Password string `json:"password"`
		ImageSrc string `json:"img_src" gorm:"column:img_src"`
	}
)

// TableName .
func (h *HostHealthy) TableName() string {
	return "host_healthy"
}

// TableName .
func (u *UserInfo) TableName() string {
	return "users"
}
