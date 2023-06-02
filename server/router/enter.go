package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Db_tools db_tools.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
