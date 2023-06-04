package tableInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos"
	tableInfosReq "github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos/request"
)

type TableInfosModelService struct {
}

// CreateTableInfosModel 创建TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) CreateTableInfosModel(tableInfoModel *tableInfos.TableInfosModel) (err error) {
	err = global.GVA_DB.Create(tableInfoModel).Error
	return err
}

// DeleteTableInfosModel 删除TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) DeleteTableInfosModel(tableInfoModel tableInfos.TableInfosModel) (err error) {
	err = global.GVA_DB.Delete(&tableInfoModel).Error
	return err
}

// DeleteTableInfosModelByIds 批量删除TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) DeleteTableInfosModelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]tableInfos.TableInfosModel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTableInfosModel 更新TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) UpdateTableInfosModel(tableInfoModel tableInfos.TableInfosModel) (err error) {
	err = global.GVA_DB.Save(&tableInfoModel).Error
	return err
}

// GetTableInfosModel 根据id获取TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) GetTableInfosModel(id uint) (tableInfoModel tableInfos.TableInfosModel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tableInfoModel).Error
	return
}

// GetTableInfosModelInfoList 分页获取TableInfosModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (tableInfoModelService *TableInfosModelService) GetTableInfosModelInfoList(info tableInfosReq.TableInfosModelSearch) (list []tableInfos.TableInfosModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&tableInfos.TableInfosModel{})
	var tableInfoModels []tableInfos.TableInfosModel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TableName != "" {
		db = db.Where("table_name LIKE ?", "%"+info.TableName+"%")
	}
	if info.TableComment != "" {
		db = db.Where("table_comment LIKE ?", "%"+info.TableComment+"%")
	}
	if info.TableSchema != "" {
		db = db.Where("table_schema = ?", info.TableSchema)
	}
	if info.TableType != "" {
		db = db.Where("table_type = ?", info.TableType)
	}
	if info.Engine != "" {
		db = db.Where("engine = ?", info.Engine)
	}
	if info.TableRows != nil {
		db = db.Where("table_rows <> ?", info.TableRows)
	}
	if info.DataLength != nil {
		db = db.Where("data_length <> ?", info.DataLength)
	}
	if info.CreateTime != "" {
		db = db.Where("create_time <> ?", info.CreateTime)
	}
	if info.DbId != 0 {
		db = db.Where("db_id = ?", info.DbId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&tableInfoModels).Error
	return tableInfoModels, total, err
}
