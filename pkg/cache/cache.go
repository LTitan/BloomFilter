package cache

import (
	"time"

	gcache "github.com/patrickmn/go-cache"
)

// DefaultCache .
var DefaultCache *gcache.Cache

func init() {
	DefaultCache = gcache.New(time.Minute*20, time.Hour)
}
