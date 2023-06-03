package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tableInfos"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	_ "go.uber.org/zap"

	_ "github.com/flipped-aurora/gin-vue-admin/server/global"
	_ "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	_ "github.com/flipped-aurora/gin-vue-admin/server/model/db_tools/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	_ "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

var tableInfoService = service.ServiceGroupApp.Db_toolsServiceGroup.TableInfoService

type TableInfoApi struct {
}

func (tableInfoApi *TableInfoApi) GetTableInfoList(c *gin.Context) {
	var dbInfo db_tools.DbInfo
	err := c.ShouldBindQuery(&dbInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	redbInfo, _ := dbInfoService.GetDbInfo(dbInfo.ID)
	if list, err := tableInfoService.GetTableList(redbInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}

}

// 批量保存表信息
func (tableInfoApi *TableInfoApi) SaveTableInfoList(c *gin.Context) {
	var tableInfosModels []tableInfos.TableInfosModel
	// 解析请求中的JSON数据到tableInfosModels切片中
	if err := c.ShouldBindJSON(&tableInfosModels); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "无效的JSON数据"})
		response.FailWithMessage("无效的JSON数据", c)
	}
	// 在这里执行保存tableInfosModels的逻辑，可以调用数据库操作等
	if err := dbInfoService.SaveTableInfoList(tableInfosModels); err != nil {
		response.FailWithMessage("表信息保存失败", c)
	} else {
		response.OkWithMessage("表信息保存成功", c)
	}
	// 返回成功的响应
	//c.JSON(http.StatusOK, gin.H{"message": "表信息保存成功"})
}
