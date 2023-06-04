package syncData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/syncData"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    syncDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/syncData/request"
)

type SyncDataModelService struct {
}

// CreateSyncDataModel 创建SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService) CreateSyncDataModel(syncDataTask *syncData.SyncDataModel) (err error) {
	err = global.GVA_DB.Create(syncDataTask).Error
	return err
}

// DeleteSyncDataModel 删除SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService)DeleteSyncDataModel(syncDataTask syncData.SyncDataModel) (err error) {
	err = global.GVA_DB.Delete(&syncDataTask).Error
	return err
}

// DeleteSyncDataModelByIds 批量删除SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService)DeleteSyncDataModelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]syncData.SyncDataModel{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSyncDataModel 更新SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService)UpdateSyncDataModel(syncDataTask syncData.SyncDataModel) (err error) {
	err = global.GVA_DB.Save(&syncDataTask).Error
	return err
}

// GetSyncDataModel 根据id获取SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService)GetSyncDataModel(id uint) (syncDataTask syncData.SyncDataModel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&syncDataTask).Error
	return
}

// GetSyncDataModelInfoList 分页获取SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService)GetSyncDataModelInfoList(info syncDataReq.SyncDataModelSearch) (list []syncData.SyncDataModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&syncData.SyncDataModel{})
    var syncDataTasks []syncData.SyncDataModel
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TaskName != "" {
        db = db.Where("task_name LIKE ?","%"+ info.TaskName+"%")
    }
    if info.SourceDbId != nil {
        db = db.Where("source_db_id = ?",info.SourceDbId)
    }
    if info.SyncType != nil {
        db = db.Where("sync_type = ?",info.SyncType)
    }
    if info.SyncDirection != nil {
        db = db.Where("sync_direction = ?",info.SyncDirection)
    }
    if info.Readme != "" {
        db = db.Where("readme LIKE ?","%"+ info.Readme+"%")
    }
    if info.TableInfoId != "" {
        db = db.Where("table_info_id LIKE ?","%"+ info.TableInfoId+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&syncDataTasks).Error
	return  syncDataTasks, total, err
}
