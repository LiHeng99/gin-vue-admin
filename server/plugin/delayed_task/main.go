package delayed_task

import (
	"github.com/gin-gonic/gin"
)

type DelayedTaskPlugin struct {
}

func CreateDelayedTaskPlugin() *DelayedTaskPlugin {
	return &DelayedTaskPlugin{}
}

func (*DelayedTaskPlugin) Register(group *gin.RouterGroup) {
	return
}

func (*DelayedTaskPlugin) RouterPath() string {
	return "delayed_task"
}
