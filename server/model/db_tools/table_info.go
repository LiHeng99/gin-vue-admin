package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TableName 结构体
type TableInfo struct {
	global.GVA_MODEL
	TableName    string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;size:255;"`
	TableComment string `json:"tableComment" form:"tableComment" gorm:"column:table_comment;comment:表名称;size:255;"`
	TableSchema  string `json:"tableSchema" form:"tableSchema" gorm:"column:table_schema;comment:数据库名称;size:255;"`
	TableType    string `json:"tableType" form:"tableType" gorm:"column:table_type;comment:表类型;size:255;"`
	Engine       string `json:"engine" form:"engine" gorm:"column:engine;comment:存储引擎;size:255;"`
	TableRows    int    `json:"tableRows" form:"tableRows" gorm:"column:table_rows;comment:行数;"`
	DataLength   int    `json:"dataLength" form:"dataLength" gorm:"column:data_length;comment:数据长度;"`
	CreateTime   string `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime;"`
}

// TableName TableName 表名
func (TableInfo) TableInfo() string {
	return "table_info"
}
