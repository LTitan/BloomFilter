package dao

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/LTitan/BloomFilter/internal/router/fe"
	"github.com/LTitan/BloomFilter/internal/router/sqldata"
	"github.com/LTitan/BloomFilter/pkg/sql"
)

// CreatedOneRecord .
func CreatedOneRecord(r *sqldata.HostHealthy) (err error) {
	db := sql.DefaultDB
	tx := db.Begin()
	var cnt int
	if err = tx.Model(&sqldata.HostHealthy{}).Where("host_ip = ? and port = ?", r.HostIP, r.Port).Count(&cnt).Error; err != nil {
		tx.Rollback()
		return
	}
	if cnt > 1000 {
		var ids []sqldata.HostHealthy
		if err = tx.Model(&sqldata.HostHealthy{}).Select("id").Where("host_ip = ? and port = ?", r.HostIP, r.Port).
			Order("created_at").Limit(cnt - 1000).Find(&ids).Error; err != nil {
			tx.Rollback()
			return
		}
		var id []uint
		for _, each := range ids {
			id = append(id, each.ID)
		}
		if err = tx.Debug().Where("id in (?)", id).Delete(&sqldata.HostHealthy{}).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	if err = tx.Model(&sqldata.HostHealthy{}).Create(r).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// QueryHostHardwareInfo .
func QueryHostHardwareInfo() (ret *fe.CPUMemoryInfo, err error) {
	db := sql.DefaultDB
	var hosts []sqldata.HostHealthy
	if err = db.Model(&sqldata.HostHealthy{}).Select("host_ip, port, cpu_num, mem_cap").Group("host_ip, port").Find(&hosts).Error; err != nil {
		return nil, err
	}
	ret = new(fe.CPUMemoryInfo)
	ret.Legend = []string{"cpu", "memory"}
	for i := range hosts {
		ret.YAxis = append(ret.YAxis, fmt.Sprintf("%s:%d", hosts[i].HostIP, hosts[i].Port))
		ret.Series.CPU = append(ret.Series.CPU, hosts[i].CPUNum)
		ret.Series.Memory = append(ret.Series.Memory, hosts[i].MemCap)
	}
	return
}

// GetRecentlyHost .
func GetRecentlyHost(now *time.Time, interval time.Duration) (ret []string, err error) {
	sub := now.Add(-interval)
	db := sql.DefaultDB
	var hosts []sqldata.HostHealthy
	if err = db.Model(&sqldata.HostHealthy{}).Select("host_ip, port").Where("created_at >= ?", sub.String()).Group("host_ip, port").Find(&hosts).Error; err != nil {
		return nil, err
	}
	for i := range hosts {
		ret = append(ret, fmt.Sprintf("%v:%v", hosts[i].HostIP, hosts[i].Port))
	}
	return
}

// CreatedApplyRecord .
func CreatedApplyRecord(ar *sqldata.ApplyRecord) (err error) {
	db := sql.DefaultDB
	tx := db.Begin().Model(&sqldata.ApplyRecord{})
	if err = tx.Create(ar).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// DeletedUpdateApplyRecord .
func DeletedUpdateApplyRecord() {
	db := sql.DefaultDB
	now := time.Now().Add(-time.Minute * 15)
	if err := db.Model(&sqldata.ApplyRecord{}).Where("status = ? and updated_at <= ?", sqldata.StatusTemporarilyUnavailable, now).Updates(map[string]interface{}{"status": sqldata.StatusDead}).Error; err != nil {
		return
	}
	if err := db.Model(&sqldata.ApplyRecord{}).Where("status = ? and updated_at <= ?", sqldata.StatusTemporarilyUnavailable, now).Delete(&sqldata.ApplyRecord{}).Error; err != nil {
		return
	}
}

// QueryApplyAddress .
func QueryApplyAddress(uuid string) (addr string, found bool) {
	db := sql.DefaultDB
	var ar sqldata.ApplyRecord
	if err := db.Model(&sqldata.ApplyRecord{}).First(&ar, "uuid = ?", uuid).Error; err != nil {
		return "", false
	}
	return ar.HostIP, true
}

// UpdateStatusApplyAddress .
func UpdateStatusApplyAddress(addr string, status string) (err error) {
	db := sql.DefaultDB
	tx := db.Begin()
	if err = tx.Model(&sqldata.ApplyRecord{}).Where("host_ip = ? and status = ?", addr, sqldata.StatusNormal).
		Updates(map[string]interface{}{"status": status}).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// GetApplyRecordsPagination .
func GetApplyRecordsPagination(pageSize, currentPage int) (ret []sqldata.ApplyRecord, page *sqldata.Page, err error) {
	page = new(sqldata.Page)
	db := sql.DefaultDB.Model(&sqldata.ApplyRecord{}).Unscoped()
	page.CurrentPage = currentPage
	currentPage--
	if err = db.Find(&ret).Count(&page.TotalSize).Error; err != nil {
		return
	}
	if err = db.Offset(currentPage * pageSize).Limit(pageSize).Order("id DESC").Find(&ret).Error; err != nil {
		return
	}
	page.TotalPage = int(math.Ceil(float64(page.TotalSize) / float64(pageSize)))
	page.CurrentSize = len(ret)
	return
}

// UpdateApplyRecord by uuid
func UpdateApplyRecord(uuid []string, updateField map[string]interface{}) (err error) {
	db := sql.DefaultDB
	tx := db.Begin().Model(&sqldata.ApplyRecord{})
	if err = tx.Where("uuid in (?)", uuid).Updates(updateField).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// DeleteApplyRecord by uuid
func DeleteApplyRecord(uuid string) (err error) {
	tx := sql.DefaultDB.Model(&sqldata.ApplyRecord{}).Unscoped().Begin()
	if err = tx.Where("uuid = ?", uuid).Delete(&sqldata.ApplyRecord{}).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// GetAliveHosts .
func GetAliveHosts() (ret map[string]interface{}, err error) {
	db := sql.DefaultDB.Model(&sqldata.ApplyRecord{}).Unscoped()
	var hosts []sqldata.ApplyRecord
	if err = db.Where("status = ? OR status = ?", sqldata.StatusNormal, sqldata.StatusTemporarilyUnavailable).Find(&hosts).Error; err != nil {
		return
	}
	ret = make(map[string]interface{}, len(hosts))
	for _, each := range hosts {
		ret[each.UUID] = each.HostIP
	}
	return
}

// GetSingleAddressInfo .
func GetSingleAddressInfo(address string) (ret fe.CPUMemoryInfo, err error) {
	db := sql.DefaultDB.Model(&sqldata.HostHealthy{})
	ips := strings.Split(address, ":")
	var hosts []sqldata.HostHealthy
	if err = db.Where("host_ip = ? AND port = ?", ips[0], ips[1]).Find(&hosts).Order("id").Limit(1000).Error; err != nil {
		return
	}
	mp := make(map[string][]sqldata.HostHealthy)
	for i := range hosts {
		tm := fmt.Sprintf("%d-%d-%d %02d:00:00", hosts[i].CreatedAt.Year(), hosts[i].CreatedAt.Month(), hosts[i].CreatedAt.Day(), hosts[i].CreatedAt.Hour())
		mp[tm] = append(mp[tm], hosts[i])
	}
	xAix := []string{}
	for key := range mp {
		xAix = append(xAix, key)
	}
	sort.Strings(xAix)
	ret.XAxis = xAix
	for _, x := range xAix {
		hs := mp[x]
		var cpu, memory float32
		for _, each := range hs {
			cpu += each.CPUUsage
			memory += each.MemUsage
		}
		cpu = cpu / float32(len(hs))
		memory = memory / float32(len(hs))
		ret.Series.CPU = append(ret.Series.CPU, cpu)
		ret.Series.Memory = append(ret.Series.Memory, memory)
	}
	ret.Legend = []string{"CPU", "Memoey"}
	mp = nil
	return
}

// RegisterDistribution .
func RegisterDistribution() (ret fe.ProductInfo, err error) {
	db := sql.DefaultDB
	var hosts []sqldata.HostHealthy
	if err = db.Model(&sqldata.HostHealthy{}).Group("host_ip, port").Find(&hosts).Error; err != nil {
		return
	}
	for _, host := range hosts {
		ret.Legend = append(ret.Legend, fmt.Sprintf("%v:%v", host.HostIP, host.Port))
	}
	for _, addr := range ret.Legend {
		cnt := 0
		if err = db.Model(&sqldata.ApplyRecord{}).Where("host_ip = ?", addr).Count(&cnt).Error; err != nil {
			return
		}
		ret.Series = append(ret.Series, fe.NameAndValue{Name: addr, Value: cnt})
	}
	return
}

// GetRegisterMemoryInfo .
func GetRegisterMemoryInfo(address string) (ret fe.ProductInfo, err error) {
	db := sql.DefaultDB.Model(&sqldata.ApplyRecord{}).Unscoped()
	var data []sqldata.ApplyRecord
	if err = db.Select("uuid, forecast_cap").Where("host_ip = ?", address).Group("uuid").Find(&data).Error; err != nil {
		return
	}
	for _, each := range data {
		ret.XAxis = append(ret.XAxis, each.UUID)
		ret.Series = append(ret.Series, each.ForecastCap)
	}
	data = nil
	return
}
