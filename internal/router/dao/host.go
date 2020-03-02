package dao

import (
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/internal/router/fe"
	"github.com/LTitan/BloomFilter/pkg/sql"
)

// CreatedOneRecord .
func CreatedOneRecord(r *sqldata.HostHealthy)(err error){
	db := sql.OpenDB()
	defer db.Close()
	tx := db.Begin()
	var cnt int
	if err = tx.Model(&sqldata.HostHealthy{}).Where("host_ip = ?", r.HostIP).Count(&cnt).Error;err!=nil{
		tx.Rollback()
		return
	}
	if cnt > 1000 {
		var ids []sqldata.HostHealthy
		if err = tx.Model(&sqldata.HostHealthy{}).Select("id").Where("host_ip = ?", r.HostIP).
			Order("created_at").Limit(cnt-1000).Find(&ids).Error;err!=nil{
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

// QueryHostHardwareInfo .
func QueryHostHardwareInfo()(ret *fe.CPUMemoryInfo, err error){
	db := sql.OpenDB()
	var hosts []sqldata.HostHealthy
	defer db.Close()
	if err = db.Model(&sqldata.HostHealthy{}).Select("host_ip, cpu_num, mem_cap").Group("host_ip").Find(&hosts).Error;err!=nil{
		return nil, err
	}
	ret = new(fe.CPUMemoryInfo)
	ret.Legend = []string{"cpu", "memory"}
	for i := range hosts{
		ret.YAxis = append(ret.YAxis, hosts[i].HostIP)
		ret.Series.CPU = append(ret.Series.CPU, hosts[i].CPUNum)
		ret.Series.Memory = append(ret.Series.Memory, hosts[i].MemCap)
	}
	return
}