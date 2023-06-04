// 自动生成模板SyncDataModel
package syncData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// SyncDataModel 结构体
type SyncDataModel struct {
      global.GVA_MODEL
      TaskName  string `json:"taskName" form:"taskName" gorm:"column:task_name;comment:任务名称;size:100;"`
      SourceDbId  *int `json:"sourceDbId" form:"sourceDbId" gorm:"column:source_db_id;comment:源数据库ID;size:10;"`
      TargetDbId  *int `json:"targetDbId" form:"targetDbId" gorm:"column:target_db_id;comment:目标数据ID;size:10;"`
      SyncType  *int `json:"syncType" form:"syncType" gorm:"column:sync_type;comment:同步类型增量或全量;size:1;"`
      SyncDirection  *int `json:"syncDirection" form:"syncDirection" gorm:"column:sync_direction;comment:同步方向单或双;size:1;"`
      Readme  string `json:"readme" form:"readme" gorm:"column:readme;comment:描述;size:255;"`
      TableInfoId  string `json:"tableInfoId" form:"tableInfoId" gorm:"column:table_info_id;comment:表信息ID;size:255;"`
}


// TableName SyncDataModel 表名
func (SyncDataModel) TableName() string {
  return "sync_data"
}

