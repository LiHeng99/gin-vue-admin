package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/db_tools"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/syncData"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/tableInfos"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	Db_toolsServiceGroup   db_tools.ServiceGroup
	TableInfosServiceGroup tableInfos.ServiceGroup
	SyncDataServiceGroup   syncData.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
