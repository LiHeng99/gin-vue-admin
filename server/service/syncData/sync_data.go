package syncData

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/model/syncData"
	syncDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/syncData/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos"
	db_toolsDb "github.com/flipped-aurora/gin-vue-admin/server/service/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/canal"
	"strconv"
	"strings"
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
func (syncDataTaskService *SyncDataModelService) DeleteSyncDataModel(syncDataTask syncData.SyncDataModel) (err error) {
	err = global.GVA_DB.Delete(&syncDataTask).Error
	return err
}

// DeleteSyncDataModelByIds 批量删除SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService) DeleteSyncDataModelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]syncData.SyncDataModel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSyncDataModel 更新SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService) UpdateSyncDataModel(syncDataTask syncData.SyncDataModel) (err error) {
	err = global.GVA_DB.Save(&syncDataTask).Error
	return err
}

// GetSyncDataModel 根据id获取SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService) GetSyncDataModel(id uint) (syncDataTask syncData.SyncDataModel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&syncDataTask).Error
	return
}

func (syncDataTaskService *SyncDataModelService) ExecuteSyncData(id uint) (syncDataTask syncData.SyncDataModel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&syncDataTask).Error

	var sourceDb db_tools.DbInfo
	sourceDb.ID = syncDataTask.SourceDbId
	err = global.GVA_DB.Where("id = ?", sourceDb.ID).First(&sourceDb).Error
	sourceDbLink := db_toolsDb.OpenDb(sourceDb)

	var targetDb db_tools.DbInfo
	targetDb.ID = syncDataTask.TargetDbId
	err = global.GVA_DB.Where("id = ?", targetDb.ID).First(&targetDb).Error
	targetDbLink := db_toolsDb.OpenDb(targetDb)

	var tableInfoList []tableInfos.TableInfosModel
	db := global.GVA_DB.Model(&tableInfos.TableInfosModel{})
	db = db.Where("id IN (" + syncDataTask.TableInfoId + ")")
	err = db.Find(&tableInfoList).Error

	for index := range tableInfoList {
		result := canal.SelectDataByTime(targetDb, "uiDataTim", tableInfoList[index].TableName)
		query := fmt.Sprintf("select * from %s where uiDataTim > %s", tableInfoList[index].TableName, strconv.FormatInt(int64(result), 10))
		rows, _ := sourceDbLink.Query(query)

		defer rows.Close()
		// 获取查询结果的列信息
		columns, _ := rows.Columns()
		// 准备一个与列数量对应的切片，用于存储查询结果的字段值
		values := make([]interface{}, len(columns))
		// 准备一个与列数量对应的切片，用于存储字段值的指针
		valuePointers := make([]interface{}, len(columns))
		for i := range columns {
			valuePointers[i] = &values[i]
		}
		if !canal.SelectTableExist(targetDb, tableInfoList[index].TableName) {
			createSql := canal.SelectCreateSql(sourceDb, tableInfoList[index].TableName)
			if createSql != "" {
				result, _ := targetDbLink.Exec(createSql)
				fmt.Println(result)
			}
		}
		// 遍历查询结果
		for rows.Next() {
			// 将查询结果的字段值扫描到指定的变量中
			err := rows.Scan(valuePointers...)
			if err != nil {
				fmt.Println(err)
			}
			// 插入数据到目标表
			_, err = targetDbLink.Exec("INSERT INTO "+tableInfoList[index].TableName+"("+strings.Join(columns, ", ")+") VALUES ("+strings.Repeat("?, ", len(columns)-1)+"?)", values...)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	// 检查遍历结果时是否出现错误
	//err = rows.Err()
	//if err != nil {
	//	return
	//}

	return
}

// GetSyncDataModelInfoList 分页获取SyncDataModel记录
// Author [piexlmax](https://github.com/piexlmax)
func (syncDataTaskService *SyncDataModelService) GetSyncDataModelInfoList(info syncDataReq.SyncDataModelSearch) (list []syncData.SyncDataModel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&syncData.SyncDataModel{})
	var syncDataTasks []syncData.SyncDataModel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TaskName != "" {
		db = db.Where("task_name LIKE ?", "%"+info.TaskName+"%")
	}
	if info.SourceDbId != 0 {
		db = db.Where("source_db_id = ?", info.SourceDbId)
	}
	if info.SyncType != nil {
		db = db.Where("sync_type = ?", info.SyncType)
	}
	if info.SyncDirection != nil {
		db = db.Where("sync_direction = ?", info.SyncDirection)
	}
	if info.Readme != "" {
		db = db.Where("readme LIKE ?", "%"+info.Readme+"%")
	}
	if info.TableInfoId != "" {
		db = db.Where("table_info_id LIKE ?", "%"+info.TableInfoId+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&syncDataTasks).Error
	return syncDataTasks, total, err
}
