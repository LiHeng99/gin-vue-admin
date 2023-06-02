package table_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TableInfoRouter struct {
}

// InitTableInfoRouter 初始化 TableInfo 路由信息
func (s *TableInfoRouter) InitTableInfoRouter(Router *gin.RouterGroup) {
	able_infoRouter := Router.Group("able_info").Use(middleware.OperationRecord())
	able_infoRouterWithoutRecord := Router.Group("able_info")
	var able_infoApi = v1.ApiGroupApp.Table_infoApiGroup.TableInfoApi
	{
		able_infoRouter.POST("createTableInfo", able_infoApi.CreateTableInfo)   // 新建TableInfo
		able_infoRouter.DELETE("deleteTableInfo", able_infoApi.DeleteTableInfo) // 删除TableInfo
		able_infoRouter.DELETE("deleteTableInfoByIds", able_infoApi.DeleteTableInfoByIds) // 批量删除TableInfo
		able_infoRouter.PUT("updateTableInfo", able_infoApi.UpdateTableInfo)    // 更新TableInfo
	}
	{
		able_infoRouterWithoutRecord.GET("findTableInfo", able_infoApi.FindTableInfo)        // 根据ID获取TableInfo
		able_infoRouterWithoutRecord.GET("getTableInfoList", able_infoApi.GetTableInfoList)  // 获取TableInfo列表
	}
}
