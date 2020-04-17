package dao

import (
	"testing"
	"time"
)

func TestGetRecentlyHost(t *testing.T) {
	now := time.Now()
	host, err := GetRecentlyHost(&now, time.Hour*48)
	t.Error(host, err)
}
