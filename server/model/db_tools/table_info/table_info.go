package table_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TableName 结构体
type TableInfo struct {
	global.GVA_MODEL
	TableName string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;size:255;"`
}

// TableName TableName 表名
func (TableInfo) TableInfo() string {
	return "table_info"
}
