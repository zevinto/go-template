package listener

import (
	"fmt"
	appEvent "go-template/app/event/event"
	"go-template/internal/event"
)

type BarListener struct{}

func (f BarListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

func (f BarListener) Process(e event.EventInterface) {
	fmt.Println("bar listener process event:", e, e.Name())
}
