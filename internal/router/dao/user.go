package dao

import (
	"crypto/md5"
	"fmt"

	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/sql"
	"github.com/jinzhu/gorm"
)

// CreateUser .
func CreateUser(u *sqldata.UserInfo) (err error) {
	db := sql.DefaultDB

	u.Password = md5String(u.Password)
	tx := db.Begin()
	if err = tx.Model(&sqldata.UserInfo{}).Create(&u).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// QueryHasUser .
func QueryHasUser(u *sqldata.UserInfo) (ret bool, err error) {
	ret = false
	db := sql.DefaultDB
	u.Password = md5String(u.Password)
	var temp sqldata.UserInfo
	if err = db.Model(&sqldata.UserInfo{}).Where("name = ?", u.Name).First(&temp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ret, nil
		}
		return
	}
	ret = temp.Password == u.Password
	return
}

func md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
