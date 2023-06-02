package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	Db_toolsApiGroup db_tools.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
