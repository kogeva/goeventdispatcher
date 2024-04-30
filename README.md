# Go Event Dispatcher

This library provides a simple event dispatcher in Go. It allows you to add and remove listeners for specific events and dispatch events to all registered listeners.

## Installation

To install this library, you can use `go get`:

```bash
go get github.com/kogeva/go-event-dispatcher
```

## Usage

Here is a basic usage example:

```go
package main

import (
	"fmt"
	"github.com/kogeva/goeventdispatcher"
)

func main() {
	dispatcher := goeventdispatcher.NewEventDispatcher()

	// Define a listener
	listener := goeventdispatcher.Listener(func(event interface{}) {
		fmt.Println("Received event:", event)
	})

	// Add the listener to the dispatcher
	dispatcher.AddListener("myEvent", &listener)

	// Dispatch an event
	dispatcher.Dispatch("myEvent", "Hello, World!")
}
```

In this example, we create a new event dispatcher, define a listener that prints any event it receives, add the listener to the dispatcher for the event "myEvent", and then dispatch the event "myEvent" with the data "Hello, World!".

## API

### `NewEventDispatcher() EventDispatcherInterface`

Creates a new event dispatcher.

### `AddListener(eventName string, listener *Listener)`

Adds a listener for the specified event.

### `RemoveListener(eventName string, listener *Listener)`

Removes a listener for the specified event.

### `GetListeners(eventName string) map[*Listener]struct{}`

Returns all listeners for the specified event.

### `Dispatch(eventName string, event interface{})`

Dispatches an event to all listeners registered for that event.

## Testing

This library includes a suite of unit tests. After installing the library, you can run the tests with `go test`:

```bash
go test github.com/kogeva/goeventdispatcher
```

## License

This library is licensed under the MIT License. See the LICENSE file for more details.