package sqldata

import (
	"time"

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
	// ApplyRecord .
	ApplyRecord struct {
		*gorm.Model
		UUID         string    `json:"uuid" gorm:"column:uuid"`
		HostIP       string    `json:"address"`
		ForecastCap  uint64    `json:"forecast_cap"`
		Status       string    `json:"status"`
		ExpirationAt time.Time `json:"expirarion_time"`
		ErrorRate    float64   `json:"error_rate" gorm:"default:0"`
	}
	// Page .
	Page struct {
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
		CurrentSize int `json:"current_size"`
		TotalSize   int `json:"total_size"`
	}
)

//
const (
	StatusNormal                 = "normal"
	StatusTemporarilyUnavailable = "unavailable"
	StatusDead                   = "dead"
	StatusRelease                = "release"
)

// TableName .
func (h *HostHealthy) TableName() string {
	return "host_healthy"
}

// TableName .
func (u *UserInfo) TableName() string {
	return "users"
}
