package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    db_toolsReq "github.com/flipped-aurora/gin-vue-admin/server/model/db_tools/request"
)

type DbInfoService struct {
}

// CreateDbInfo 创建DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService) CreateDbInfo(dbInfo *db_tools.DbInfo) (err error) {
	err = global.GVA_DB.Create(dbInfo).Error
	return err
}

// DeleteDbInfo 删除DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService)DeleteDbInfo(dbInfo db_tools.DbInfo) (err error) {
	err = global.GVA_DB.Delete(&dbInfo).Error
	return err
}

// DeleteDbInfoByIds 批量删除DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService)DeleteDbInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]db_tools.DbInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDbInfo 更新DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService)UpdateDbInfo(dbInfo db_tools.DbInfo) (err error) {
	err = global.GVA_DB.Save(&dbInfo).Error
	return err
}

// GetDbInfo 根据id获取DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService)GetDbInfo(id uint) (dbInfo db_tools.DbInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dbInfo).Error
	return
}

// GetDbInfoInfoList 分页获取DbInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (dbInfoService *DbInfoService)GetDbInfoInfoList(info db_toolsReq.DbInfoSearch) (list []db_tools.DbInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&db_tools.DbInfo{})
    var dbInfos []db_tools.DbInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.DbUserName != "" {
        db = db.Where("db_user_name LIKE ?","%"+ info.DbUserName+"%")
    }
    if info.DbPassword != "" {
        db = db.Where("db_password = ?",info.DbPassword)
    }
    if info.DbUrl != "" {
        db = db.Where("db_url LIKE ?","%"+ info.DbUrl+"%")
    }
    if info.DriverClassName != "" {
        db = db.Where("driver_class_name LIKE ?","%"+ info.DriverClassName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dbInfos).Error
	return  dbInfos, total, err
}
