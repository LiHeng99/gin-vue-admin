package syncData

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/syncData"
	syncDataReq "github.com/flipped-aurora/gin-vue-admin/server/model/syncData/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SyncDataModelApi struct {
}

var syncDataTaskService = service.ServiceGroupApp.SyncDataServiceGroup.SyncDataModelService

// CreateSyncDataModel 创建SyncDataModel
// @Tags SyncDataModel
// @Summary 创建SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body syncData.SyncDataModel true "创建SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /syncDataTask/createSyncDataModel [post]
func (syncDataTaskApi *SyncDataModelApi) CreateSyncDataModel(c *gin.Context) {
	var syncDataTask syncData.SyncDataModel
	err := c.ShouldBindJSON(&syncDataTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TaskName":      {utils.NotEmpty()},
		"SourceDbId":    {utils.NotEmpty()},
		"TargetDbId":    {utils.NotEmpty()},
		"SyncType":      {utils.NotEmpty()},
		"SyncDirection": {utils.NotEmpty()},
		"TableInfoId":   {utils.NotEmpty()},
	}
	if err := utils.Verify(syncDataTask, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := syncDataTaskService.CreateSyncDataModel(&syncDataTask); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSyncDataModel 删除SyncDataModel
// @Tags SyncDataModel
// @Summary 删除SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body syncData.SyncDataModel true "删除SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /syncDataTask/deleteSyncDataModel [delete]
func (syncDataTaskApi *SyncDataModelApi) DeleteSyncDataModel(c *gin.Context) {
	var syncDataTask syncData.SyncDataModel
	err := c.ShouldBindJSON(&syncDataTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := syncDataTaskService.DeleteSyncDataModel(syncDataTask); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSyncDataModelByIds 批量删除SyncDataModel
// @Tags SyncDataModel
// @Summary 批量删除SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /syncDataTask/deleteSyncDataModelByIds [delete]
func (syncDataTaskApi *SyncDataModelApi) DeleteSyncDataModelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := syncDataTaskService.DeleteSyncDataModelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSyncDataModel 更新SyncDataModel
// @Tags SyncDataModel
// @Summary 更新SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body syncData.SyncDataModel true "更新SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /syncDataTask/updateSyncDataModel [put]
func (syncDataTaskApi *SyncDataModelApi) UpdateSyncDataModel(c *gin.Context) {
	var syncDataTask syncData.SyncDataModel
	err := c.ShouldBindJSON(&syncDataTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TaskName":      {utils.NotEmpty()},
		"SourceDbId":    {utils.NotEmpty()},
		"TargetDbId":    {utils.NotEmpty()},
		"SyncType":      {utils.NotEmpty()},
		"SyncDirection": {utils.NotEmpty()},
		"TableInfoId":   {utils.NotEmpty()},
	}
	if err := utils.Verify(syncDataTask, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := syncDataTaskService.UpdateSyncDataModel(syncDataTask); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSyncDataModel 用id查询SyncDataModel
// @Tags SyncDataModel
// @Summary 用id查询SyncDataModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query syncData.SyncDataModel true "用id查询SyncDataModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /syncDataTask/findSyncDataModel [get]
func (syncDataTaskApi *SyncDataModelApi) FindSyncDataModel(c *gin.Context) {
	var syncDataTask syncData.SyncDataModel
	err := c.ShouldBindQuery(&syncDataTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resyncDataTask, err := syncDataTaskService.GetSyncDataModel(syncDataTask.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resyncDataTask": resyncDataTask}, c)
	}
}

// GetSyncDataModelList 分页获取SyncDataModel列表
// @Tags SyncDataModel
// @Summary 分页获取SyncDataModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query syncDataReq.SyncDataModelSearch true "分页获取SyncDataModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /syncDataTask/getSyncDataModelList [get]
func (syncDataTaskApi *SyncDataModelApi) GetSyncDataModelList(c *gin.Context) {
	var pageInfo syncDataReq.SyncDataModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := syncDataTaskService.GetSyncDataModelInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 执行一次
func (syncDataTaskApi *SyncDataModelApi) ExecuteSyncData(c *gin.Context) {
	var syncDataTask syncData.SyncDataModel
	err := c.ShouldBindQuery(&syncDataTask)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resyncDataTask, err := syncDataTaskService.ExecuteSyncData(syncDataTask.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resyncDataTask": resyncDataTask}, c)
	}
}
