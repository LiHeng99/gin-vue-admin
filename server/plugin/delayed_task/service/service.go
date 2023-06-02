package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/delayed_task/global"
	"github.com/songzhibin97/gkit/delayed"
)

type DelayedTaskService struct{}

func (e *DelayedTaskService) AddDelayedTask(delayed delayed.Delayed) (err error) {
	global.Delayed.AddDelayed(delayed)
	return nil
}
