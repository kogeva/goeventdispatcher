package goeventdispatcher

type EventDispatcher struct {
	listeners map[string]map[*Listener]struct{}
}

func NewEventDispatcher() EventDispatcherInterface {
	return &EventDispatcher{
		listeners: make(map[string]map[*Listener]struct{}),
	}
}

func (e *EventDispatcher) Dispatch(eventName string, event interface{}) {
	listeners := e.GetListeners(eventName)

	for listenerPointer := range listeners {
		listener := *listenerPointer
		listener(event)
	}
}

func (e *EventDispatcher) AddListener(eventName string, listener *Listener) {
	listeners, ok := e.listeners[eventName]

	if !ok {
		listeners = make(map[*Listener]struct{})
		e.listeners[eventName] = listeners
	}

	listeners[listener] = struct{}{}
}

func (e *EventDispatcher) GetListeners(eventName string) map[*Listener]struct{} {
	listeners, ok := e.listeners[eventName]

	if !ok {
		return nil
	}

	return listeners
}

func (e *EventDispatcher) RemoveListener(eventName string, listener *Listener) {
	listeners, ok := e.listeners[eventName]

	if !ok {
		return
	}

	if _, ok := listeners[listener]; ok {
		delete(listeners, listener)
	}
}
