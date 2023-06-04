package db_tools

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
)

type DbInfoRouter struct {
}

// InitDbInfoRouter 初始化 DbInfo 路由信息
func (s *DbInfoRouter) InitDbInfoRouter(Router *gin.RouterGroup) {
	dbInfoRouter := Router.Group("dbInfo").Use(middleware.OperationRecord())
	dbInfoRouterWithoutRecord := Router.Group("dbInfo")
	var dbInfoApi = v1.ApiGroupApp.Db_toolsApiGroup.DbInfoApi
	{
		dbInfoRouter.POST("createDbInfo", dbInfoApi.CreateDbInfo)             // 新建DbInfo
		dbInfoRouter.DELETE("deleteDbInfo", dbInfoApi.DeleteDbInfo)           // 删除DbInfo
		dbInfoRouter.DELETE("deleteDbInfoByIds", dbInfoApi.DeleteDbInfoByIds) // 批量删除DbInfo
		dbInfoRouter.PUT("updateDbInfo", dbInfoApi.UpdateDbInfo)              // 更新DbInfo
	}
	{
		dbInfoRouterWithoutRecord.GET("findDbInfo", dbInfoApi.FindDbInfo)       // 根据ID获取DbInfo
		dbInfoRouterWithoutRecord.GET("getDbInfoList", dbInfoApi.GetDbInfoList) // 获取DbInfo列表
		dbInfoRouterWithoutRecord.GET("linkDbUrl", dbInfoApi.LinkDbUrl)         // 连接数据库
		dbInfoRouterWithoutRecord.GET("saveDataBase", dbInfoApi.SaveDataBase)   // 连接数据库
	}
}
