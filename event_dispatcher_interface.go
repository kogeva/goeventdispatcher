package goeventdispatcher

type Event interface {
	GetName() string
}

type Listener func(event interface{})

type Dispatcher interface {
	Dispatch(eventName string, event interface{})
}

type EventDispatcherInterface interface {
	Dispatcher
	AddListener(eventName string, listener *Listener)
	GetListeners(eventName string) map[*Listener]struct{}
	RemoveListener(eventName string, listener *Listener)
}
