package table_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/table_info"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    table_infoReq "github.com/flipped-aurora/gin-vue-admin/server/model/table_info/request"
)

type TableInfoService struct {
}

// CreateTableInfo 创建TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService) CreateTableInfo(able_info *table_info.TableInfo) (err error) {
	err = global.GVA_DB.Create(able_info).Error
	return err
}

// DeleteTableInfo 删除TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService)DeleteTableInfo(able_info table_info.TableInfo) (err error) {
	err = global.GVA_DB.Delete(&able_info).Error
	return err
}

// DeleteTableInfoByIds 批量删除TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService)DeleteTableInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]table_info.TableInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTableInfo 更新TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService)UpdateTableInfo(able_info table_info.TableInfo) (err error) {
	err = global.GVA_DB.Save(&able_info).Error
	return err
}

// GetTableInfo 根据id获取TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService)GetTableInfo(id uint) (able_info table_info.TableInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&able_info).Error
	return
}

// GetTableInfoInfoList 分页获取TableInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (able_infoService *TableInfoService)GetTableInfoInfoList(info table_infoReq.TableInfoSearch) (list []table_info.TableInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&table_info.TableInfo{})
    var able_infos []table_info.TableInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TableName != "" {
        db = db.Where("table_name LIKE ?","%"+ info.TableName+"%")
    }
    if info.TableComment != "" {
        db = db.Where("table_comment LIKE ?","%"+ info.TableComment+"%")
    }
    if info.TableSchema != "" {
        db = db.Where("table_schema LIKE ?","%"+ info.TableSchema+"%")
    }
    if info.TableType != "" {
        db = db.Where("table_type LIKE ?","%"+ info.TableType+"%")
    }
    if info.Engine != "" {
        db = db.Where("engine LIKE ?","%"+ info.Engine+"%")
    }
    if info.TableRows != nil {
        db = db.Where("table_rows = ?",info.TableRows)
    }
    if info.DataLength != nil {
        db = db.Where("data_length <> ?",info.DataLength)
    }
    if info.CreateTime != "" {
        db = db.Where("create_time <> ?",info.CreateTime)
    }
    if info.DbId != nil {
        db = db.Where("db_id = ?",info.DbId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["tableName"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&able_infos).Error
	return  able_infos, total, err
}
