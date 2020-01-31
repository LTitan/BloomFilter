package dao

import (
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/sql"
)

// CreatedOneRecord .
func CreatedOneRecord(r *sqldata.HostHealthy)(err error){
	db := sql.OpenDB()
	tx := db.Begin()
	var cnt int
	if err = tx.Model(&sqldata.HostHealthy{}).Where("host_ip = ?", r.HostIP).Count(&cnt).Error;err!=nil{
		tx.Rollback()
		return
	}
	if cnt > 10 {
		var ids []sqldata.HostHealthy
		if err = tx.Model(&sqldata.HostHealthy{}).Select("id").Where("host_ip = ?", r.HostIP).
			Order("created_at").Limit(cnt-10).Find(&ids).Error;err!=nil{
				tx.Rollback()
				return
			}
		var id []uint
		for _, each := range ids{
			id = append(id, each.ID)
		}
		if err = tx.Debug().Where("id in (?)", id).Delete(&sqldata.HostHealthy{}).Error;err!=nil{
			tx.Rollback()
			return
		}
	}
	if err = tx.Model(&sqldata.HostHealthy{}).Create(r).Error;err!=nil{
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}