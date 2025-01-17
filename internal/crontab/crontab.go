package crontab

import "github.com/robfig/cron/v3"

// Crontab 类型
type Crontab struct {
	cron *cron.Cron
}

// Init 函数用于初始化一个 crontab 实例
func Init() *Crontab {
	c := &Crontab{
		cron: cron.New(),
	}
	return c
}

// TaskInterface 任务接口
type TaskInterface interface {
	Spec() string
	Fn() func()
}

// Start 启动 crontab 服务
func (c *Crontab) Start() {
	c.cron.Start()
}

// Stop 停止 crontab 服务
func (c *Crontab) Stop() {
	c.cron.Stop()
}

// AddTask 添加任务
func (c *Crontab) AddTask(tasks ...TaskInterface) {
	if len(tasks) == 0 {
		return
	}
	for _, task := range tasks {
		_, err := c.cron.AddFunc(task.Spec(), task.Fn())
		if err != nil {
			return
		}
	}
}
