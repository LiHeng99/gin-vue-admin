package service

type ServiceGroup struct {
	DelayedTaskService
}

var ServiceGroupApp = new(ServiceGroup)
