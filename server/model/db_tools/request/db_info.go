package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DbInfoSearch struct{
    db_tools.DbInfo
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
