package global

import "github.com/songzhibin97/gkit/delayed"

var (
	// Delayed 延时任务实例
	// 还可以传更多配置
	/*
		// SetCheckTime 设置检查时间
		// SetWorkerNumber 设置并发数
		// SetSingle 设置监控信号
		// SetSingleCallback 设置监控信号回调

	*/
	Delayed = delayed.NewDispatchingDelayed()
)
