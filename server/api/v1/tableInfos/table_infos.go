package tableInfos

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos"
	tableInfosReq "github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TableInfosModelApi struct {
}

var tableInfoModelService = service.ServiceGroupApp.TableInfosServiceGroup.TableInfosModelService

// CreateTableInfosModel 创建TableInfosModel
// @Tags TableInfosModel
// @Summary 创建TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tableInfos.TableInfosModel true "创建TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tableInfoModel/createTableInfosModel [post]
func (tableInfoModelApi *TableInfosModelApi) CreateTableInfosModel(c *gin.Context) {
	var tableInfoModel tableInfos.TableInfosModel
	err := c.ShouldBindJSON(&tableInfoModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TableName": {utils.NotEmpty()},
		"DbId":      {utils.NotEmpty()},
	}
	if err := utils.Verify(tableInfoModel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tableInfoModelService.CreateTableInfosModel(&tableInfoModel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTableInfosModel 删除TableInfosModel
// @Tags TableInfosModel
// @Summary 删除TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tableInfos.TableInfosModel true "删除TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tableInfoModel/deleteTableInfosModel [delete]
func (tableInfoModelApi *TableInfosModelApi) DeleteTableInfosModel(c *gin.Context) {
	var tableInfoModel tableInfos.TableInfosModel
	err := c.ShouldBindJSON(&tableInfoModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tableInfoModelService.DeleteTableInfosModel(tableInfoModel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTableInfosModelByIds 批量删除TableInfosModel
// @Tags TableInfosModel
// @Summary 批量删除TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tableInfoModel/deleteTableInfosModelByIds [delete]
func (tableInfoModelApi *TableInfosModelApi) DeleteTableInfosModelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tableInfoModelService.DeleteTableInfosModelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTableInfosModel 更新TableInfosModel
// @Tags TableInfosModel
// @Summary 更新TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tableInfos.TableInfosModel true "更新TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tableInfoModel/updateTableInfosModel [put]
func (tableInfoModelApi *TableInfosModelApi) UpdateTableInfosModel(c *gin.Context) {
	var tableInfoModel tableInfos.TableInfosModel
	err := c.ShouldBindJSON(&tableInfoModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"TableName": {utils.NotEmpty()},
		"DbId":      {utils.NotEmpty()},
	}
	if err := utils.Verify(tableInfoModel, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tableInfoModelService.UpdateTableInfosModel(tableInfoModel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTableInfosModel 用id查询TableInfosModel
// @Tags TableInfosModel
// @Summary 用id查询TableInfosModel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tableInfos.TableInfosModel true "用id查询TableInfosModel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tableInfoModel/findTableInfosModel [get]
func (tableInfoModelApi *TableInfosModelApi) FindTableInfosModel(c *gin.Context) {
	var tableInfoModel tableInfos.TableInfosModel
	err := c.ShouldBindQuery(&tableInfoModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retableInfoModel, err := tableInfoModelService.GetTableInfosModel(tableInfoModel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retableInfoModel": retableInfoModel}, c)
	}
}

// GetTableInfosModelList 分页获取TableInfosModel列表
// @Tags TableInfosModel
// @Summary 分页获取TableInfosModel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tableInfosReq.TableInfosModelSearch true "分页获取TableInfosModel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tableInfoModel/getTableInfosModelList [get]
func (tableInfoModelApi *TableInfosModelApi) GetTableInfosModelList(c *gin.Context) {
	var pageInfo tableInfosReq.TableInfosModelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tableInfoModelService.GetTableInfosModelInfoList(pageInfo); err != nil {
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
