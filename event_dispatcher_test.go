package goeventdispatcher

import (
	"testing"
)

func TestAddListener(t *testing.T) {
	dispatcher := NewEventDispatcher()
	listener := Listener(func(event interface{}) {})

	dispatcher.AddListener("testEvent", &listener)

	if _, ok := dispatcher.GetListeners("testEvent")[&listener]; !ok {
		t.Errorf("Listener was not added to the event")
	}
}

func TestRemoveListener(t *testing.T) {
	dispatcher := NewEventDispatcher()
	listener := Listener(func(event interface{}) {})
	dispatcher.AddListener("testEvent", &listener)
	dispatcher.RemoveListener("testEvent", &listener)

	if _, ok := dispatcher.GetListeners("testEvent")[&listener]; ok {
		t.Errorf("Listener was not removed from the event")
	}
}

func TestDispatch(t *testing.T) {
	dispatcher := NewEventDispatcher()
	isCalled := false
	listener := Listener(func(event interface{}) {
		isCalled = true
	})
	dispatcher.AddListener("testEvent", &listener)

	if isCalled {
		t.Errorf("Listener was called before dispatch")
	}

	dispatcher.Dispatch("testEvent", nil)

	if !isCalled {
		t.Errorf("Listener was not called during dispatch")
	}
}

func TestGetListeners(t *testing.T) {
	dispatcher := NewEventDispatcher()
	listener := Listener(func(event interface{}) {})
	dispatcher.AddListener("testEvent", &listener)

	listeners := dispatcher.GetListeners("testEvent")

	if _, ok := listeners[&listener]; !ok {
		t.Errorf("GetListeners did not return the correct listeners")
	}
}

func TestGetListenersWithNoListeners(t *testing.T) {
	dispatcher := NewEventDispatcher()

	listeners := dispatcher.GetListeners("testEvent")

	if len(listeners) != 0 {
		t.Errorf("GetListeners did not return an empty map for an event with no listeners")
	}
}

func TestRemoveListenerWithNoListeners(t *testing.T) {
	dispatcher := NewEventDispatcher()
	listener := Listener(func(event interface{}) {})

	dispatcher.RemoveListener("testEvent", &listener)
}
