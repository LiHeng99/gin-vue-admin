package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/utils/canal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	db_toolsReq "github.com/flipped-aurora/gin-vue-admin/server/model/db_tools/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DbInfoApi struct {
}

var dbInfoService = service.ServiceGroupApp.Db_toolsServiceGroup.DbInfoService

// CreateDbInfo 创建DbInfo
// @Tags DbInfo
// @Summary 创建DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body db_tools.DbInfo true "创建DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dbInfo/createDbInfo [post]
func (dbInfoApi *DbInfoApi) CreateDbInfo(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindJSON(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"DbName":          {utils.NotEmpty()},
		"DbUserName":      {utils.NotEmpty()},
		"DbPassword":      {utils.NotEmpty()},
		"DbUrl":           {utils.NotEmpty()},
		"DriverClassName": {utils.NotEmpty()},
	}
	if err := utils.Verify(dbInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dbInfoService.CreateDbInfo(&dbInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDbInfo 删除DbInfo
// @Tags DbInfo
// @Summary 删除DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body db_tools.DbInfo true "删除DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dbInfo/deleteDbInfo [delete]
func (dbInfoApi *DbInfoApi) DeleteDbInfo(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindJSON(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dbInfoService.DeleteDbInfo(dbInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDbInfoByIds 批量删除DbInfo
// @Tags DbInfo
// @Summary 批量删除DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dbInfo/deleteDbInfoByIds [delete]
func (dbInfoApi *DbInfoApi) DeleteDbInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dbInfoService.DeleteDbInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDbInfo 更新DbInfo
// @Tags DbInfo
// @Summary 更新DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body db_tools.DbInfo true "更新DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dbInfo/updateDbInfo [put]
func (dbInfoApi *DbInfoApi) UpdateDbInfo(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindJSON(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"DbName":          {utils.NotEmpty()},
		"DbUserName":      {utils.NotEmpty()},
		"DbPassword":      {utils.NotEmpty()},
		"DbUrl":           {utils.NotEmpty()},
		"DriverClassName": {utils.NotEmpty()},
	}
	if err := utils.Verify(dbInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dbInfoService.UpdateDbInfo(dbInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDbInfo 用id查询DbInfo
// @Tags DbInfo
// @Summary 用id查询DbInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query db_tools.DbInfo true "用id查询DbInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dbInfo/findDbInfo [get]
func (dbInfoApi *DbInfoApi) FindDbInfo(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindQuery(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redbInfo, err := dbInfoService.GetDbInfo(dbInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redbInfo": redbInfo}, c)
	}
}

// GetDbInfoList 分页获取DbInfo列表
// @Tags DbInfo
// @Summary 分页获取DbInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query db_toolsReq.DbInfoSearch true "分页获取DbInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dbInfo/getDbInfoList [get]
func (dbInfoApi *DbInfoApi) GetDbInfoList(c *gin.Context) {
	var pageInfo db_toolsReq.DbInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := dbInfoService.GetDbInfoInfoList(pageInfo); err != nil {
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
func (DbInfoApi *DbInfoApi) LinkDbUrl(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindQuery(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redbInfo, err := dbInfoService.LinkDbUrlFunc(dbInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redbInfo": redbInfo}, c)
	}
}
func (DbInfoApi *DbInfoApi) SaveDataBase(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindQuery(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redbInfo, err := dbInfoService.SaveDataBase(dbInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		canal.SaveDataBaseFunc(redbInfo)
		response.OkWithData(gin.H{"redbInfo": redbInfo}, c)
	}
}
