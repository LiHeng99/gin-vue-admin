## delayed task
### 实现接口
```go
type mockDelayed struct {
	exec  int64
	index int
}

func (m mockDelayed) Do() {
	fmt.Println(m.index)
}

func (m mockDelayed) ExecTime() int64 {
	return m.exec
}

func (m mockDelayed) Identify() string {
	return "mock"
}

func nowMockDelayed(exec int64) Delayed {
	return mockDelayed{exec: exec}
}
```
### 使用
```go
    // 添加任务
    service.DelayedTaskService.AddDelayedTask(nowMockDelayed(time.Now().Add(time.Duration(i) * time.Second))
```