// 自动生成模板TableInfosModel
package tableInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TableInfosModel 结构体
type TableInfosModel struct {
	global.GVA_MODEL
	TableName    string `json:"tableName" form:"tableName" gorm:"column:table_name;comment:表名称;size:32;"`
	TableComment string `json:"tableComment" form:"tableComment" gorm:"column:table_comment;comment:表注释;size:100;"`
	TableSchema  string `json:"tableSchema" form:"tableSchema" gorm:"column:table_schema;comment:数据库名称;size:32;"`
	TableType    string `json:"tableType" form:"tableType" gorm:"column:table_type;comment:表类型;size:32;"`
	Engine       string `json:"engine" form:"engine" gorm:"column:engine;comment:存储引擎;size:32;"`
	TableRows    *int   `json:"tableRows" form:"tableRows" gorm:"column:table_rows;comment:行数;size:10;"`
	DataLength   *int   `json:"dataLength" form:"dataLength" gorm:"column:data_length;comment:数据长度;size:10;"`
	CreateTime   string `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;size:32;"`
	DbId         uint   `json:"dbId" form:"dbId" gorm:"column:db_id;comment:数据库id;size:10;"`
}

// TableName TableInfosModel 表名
func (TableInfosModel) TableInfosModel() string {
	return "tableInfos"
}
