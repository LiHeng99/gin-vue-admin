package tableInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TableInfosModelRouter struct {
}

// InitTableInfosModelRouter 初始化 TableInfosModel 路由信息
func (s *TableInfosModelRouter) InitTableInfosModelRouter(Router *gin.RouterGroup) {
	tableInfoModelRouter := Router.Group("tableInfoModel").Use(middleware.OperationRecord())
	tableInfoModelRouterWithoutRecord := Router.Group("tableInfoModel")
	var tableInfoModelApi = v1.ApiGroupApp.TableInfosApiGroup.TableInfosModelApi
	{
		tableInfoModelRouter.POST("createTableInfosModel", tableInfoModelApi.CreateTableInfosModel)             // 新建TableInfosModel
		tableInfoModelRouter.DELETE("deleteTableInfosModel", tableInfoModelApi.DeleteTableInfosModel)           // 删除TableInfosModel
		tableInfoModelRouter.DELETE("deleteTableInfosModelByIds", tableInfoModelApi.DeleteTableInfosModelByIds) // 批量删除TableInfosModel
		tableInfoModelRouter.PUT("updateTableInfosModel", tableInfoModelApi.UpdateTableInfosModel)              // 更新TableInfosModel
	}
	{
		tableInfoModelRouterWithoutRecord.GET("findTableInfosModel", tableInfoModelApi.FindTableInfosModel)       // 根据ID获取TableInfosModel
		tableInfoModelRouterWithoutRecord.GET("getTableInfosModelList", tableInfoModelApi.GetTableInfosModelList) // 获取TableInfosModel列表
	}
}
