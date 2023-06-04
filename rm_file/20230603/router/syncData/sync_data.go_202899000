package syncData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SyncDataModelRouter struct {
}

// InitSyncDataModelRouter 初始化 SyncDataModel 路由信息
func (s *SyncDataModelRouter) InitSyncDataModelRouter(Router *gin.RouterGroup) {
	syncDataTaskRouter := Router.Group("syncDataTask").Use(middleware.OperationRecord())
	syncDataTaskRouterWithoutRecord := Router.Group("syncDataTask")
	var syncDataTaskApi = v1.ApiGroupApp.SyncDataApiGroup.SyncDataModelApi
	{
		syncDataTaskRouter.POST("createSyncDataModel", syncDataTaskApi.CreateSyncDataModel)   // 新建SyncDataModel
		syncDataTaskRouter.DELETE("deleteSyncDataModel", syncDataTaskApi.DeleteSyncDataModel) // 删除SyncDataModel
		syncDataTaskRouter.DELETE("deleteSyncDataModelByIds", syncDataTaskApi.DeleteSyncDataModelByIds) // 批量删除SyncDataModel
		syncDataTaskRouter.PUT("updateSyncDataModel", syncDataTaskApi.UpdateSyncDataModel)    // 更新SyncDataModel
	}
	{
		syncDataTaskRouterWithoutRecord.GET("findSyncDataModel", syncDataTaskApi.FindSyncDataModel)        // 根据ID获取SyncDataModel
		syncDataTaskRouterWithoutRecord.GET("getSyncDataModelList", syncDataTaskApi.GetSyncDataModelList)  // 获取SyncDataModel列表
	}
}
