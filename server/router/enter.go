package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/syncData"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/tableInfos"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Db_tools   db_tools.RouterGroup
	TableInfos tableInfos.RouterGroup
	SyncData   syncData.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
