package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type TableInfoRouter struct {
}

func (s *TableInfoRouter) InitTableInfoRouter(Router *gin.RouterGroup) {
	tableInfoRouter := Router.Group("tableInfo").Use(middleware.OperationRecord())
	tableInfoRouterWithoutRecord := Router.Group("tableInfo")
	var tableInfoApi = v1.ApiGroupApp.Db_toolsApiGroup.TableInfoApi
	{
		tableInfoRouter.POST("createDbInfo", tableInfoApi.GetTableInfoList) // 新建DbInfo
	}
	{
		tableInfoRouterWithoutRecord.GET("getTableInfoList", tableInfoApi.GetTableInfoList) // 获取DbInfo列表
	}
}
