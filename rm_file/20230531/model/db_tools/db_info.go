// 自动生成模板DbInfo
package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// DbInfo 结构体
type DbInfo struct {
      global.GVA_MODEL
      DbUserName  string `json:"dbUserName" form:"dbUserName" gorm:"column:db_user_name;comment:数据库名称;size:32;"`
      DbPassword  string `json:"dbPassword" form:"dbPassword" gorm:"column:db_password;comment:密码;size:32;"`
      DbUrl  string `json:"dbUrl" form:"dbUrl" gorm:"column:db_url;comment:url;size:255;"`
      DriverClassName  string `json:"driverClassName" form:"driverClassName" gorm:"column:driver_class_name;comment:数据库驱动;size:255;"`
}


// TableName DbInfo 表名
func (DbInfo) TableName() string {
  return "db_info"
}

