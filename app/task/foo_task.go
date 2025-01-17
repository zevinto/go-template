package task

import "fmt"

type FooTask struct{}

func (t *FooTask) Spec() string {
	return "@every 3s"
}

func (t *FooTask) Fn() func() {
	return func() {
		fmt.Println("foo task executed")
	}
}
