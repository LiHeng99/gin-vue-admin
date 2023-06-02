package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	Db_toolsServiceGroup db_tools.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
