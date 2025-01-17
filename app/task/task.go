package task

import "go-template/internal/crontab"

func Tasks() []crontab.TaskInterface {
	return []crontab.TaskInterface{
		&FooTask{},
	}
}
