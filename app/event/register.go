package event

import (
	"go-template/app/event/listener"
	"go-template/internal/event"
	"go-template/internal/global"
)

var listenerList = []event.ListenerInterface{
	listener.FooListener{},
	listener.BarListener{},
}

func init() {
	for _, l := range listenerList {
		global.EventDispatcher.Register(l)
	}
}
