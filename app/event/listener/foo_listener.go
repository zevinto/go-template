package listener

import (
	"fmt"
	appEvent "go-template/app/event/event"
	"go-template/internal/event"
)

type FooListener struct{}

func (f FooListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

func (f FooListener) Process(e event.EventInterface) {
	fmt.Println("foo listener process event:", e, e.Name())
}
