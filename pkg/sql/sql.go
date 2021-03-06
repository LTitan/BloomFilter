package sql

import (
	"fmt"

	"github.com/LTitan/BloomFilter/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	host, port, user, passwd, dbName interface{}
)

func init() {
	host = config.Conf.Get("db.host")
	port = config.Conf.Get("db.port")
	user = config.Conf.Get("db.user")
	passwd = config.Conf.Get("db.password")
	dbName = config.Conf.Get("db.database")
	var err error
	DefaultDB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", user, passwd, host, port, dbName))
	if err != nil {
		panic(err)
	}
	DefaultDB.DB().SetMaxIdleConns(5)
	DefaultDB.DB().SetMaxOpenConns(10)
}

// DefaultDB gloabl db
var DefaultDB *gorm.DB
