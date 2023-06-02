package db_tools

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
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
