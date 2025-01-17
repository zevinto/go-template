package test

import (
	"go-template/app/event/event"
	"go-template/internal/global"
	"testing"
)

func TestEvent(t *testing.T) {
	global.EventDispatcher.Dispatch(&event.FooEvent{})
}
