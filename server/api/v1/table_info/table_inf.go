package table_info

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/table_info"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    table_infoReq "github.com/flipped-aurora/gin-vue-admin/server/model/table_info/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type TableInfoApi struct {
}

var able_infoService = service.ServiceGroupApp.Table_infoServiceGroup.TableInfoService


// CreateTableInfo 创建TableInfo
// @Tags TableInfo
// @Summary 创建TableInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body table_info.TableInfo true "创建TableInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /able_info/createTableInfo [post]
func (able_infoApi *TableInfoApi) CreateTableInfo(c *gin.Context) {
	var able_info table_info.TableInfo
	err := c.ShouldBindJSON(&able_info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "TableName":{utils.NotEmpty()},
        "TableSchema":{utils.NotEmpty()},
    }
	if err := utils.Verify(able_info, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := able_infoService.CreateTableInfo(&able_info); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTableInfo 删除TableInfo
// @Tags TableInfo
// @Summary 删除TableInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body table_info.TableInfo true "删除TableInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /able_info/deleteTableInfo [delete]
func (able_infoApi *TableInfoApi) DeleteTableInfo(c *gin.Context) {
	var able_info table_info.TableInfo
	err := c.ShouldBindJSON(&able_info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := able_infoService.DeleteTableInfo(able_info); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTableInfoByIds 批量删除TableInfo
// @Tags TableInfo
// @Summary 批量删除TableInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TableInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /able_info/deleteTableInfoByIds [delete]
func (able_infoApi *TableInfoApi) DeleteTableInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := able_infoService.DeleteTableInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTableInfo 更新TableInfo
// @Tags TableInfo
// @Summary 更新TableInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body table_info.TableInfo true "更新TableInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /able_info/updateTableInfo [put]
func (able_infoApi *TableInfoApi) UpdateTableInfo(c *gin.Context) {
	var able_info table_info.TableInfo
	err := c.ShouldBindJSON(&able_info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "TableName":{utils.NotEmpty()},
          "TableSchema":{utils.NotEmpty()},
      }
    if err := utils.Verify(able_info, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := able_infoService.UpdateTableInfo(able_info); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTableInfo 用id查询TableInfo
// @Tags TableInfo
// @Summary 用id查询TableInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query table_info.TableInfo true "用id查询TableInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /able_info/findTableInfo [get]
func (able_infoApi *TableInfoApi) FindTableInfo(c *gin.Context) {
	var able_info table_info.TableInfo
	err := c.ShouldBindQuery(&able_info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reable_info, err := able_infoService.GetTableInfo(able_info.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reable_info": reable_info}, c)
	}
}

// GetTableInfoList 分页获取TableInfo列表
// @Tags TableInfo
// @Summary 分页获取TableInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query table_infoReq.TableInfoSearch true "分页获取TableInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /able_info/getTableInfoList [get]
func (able_infoApi *TableInfoApi) GetTableInfoList(c *gin.Context) {
	var pageInfo table_infoReq.TableInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := able_infoService.GetTableInfoInfoList(pageInfo); err != nil {
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
