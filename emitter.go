package main

import (
	"fmt"
	"sync"
)

// Emitter 结构体
type Emitter struct {
	listeners map[string][]func(interface{})
	mutex     sync.Mutex
}

// 创建一个新的 Emitter
func NewEmitter() *Emitter {
	return &Emitter{
		listeners: make(map[string][]func(interface{})),
	}
}

// 注册一个事件监听器
func (e *Emitter) On(event string, listener func(interface{})) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.listeners[event] = append(e.listeners[event], listener)
}

// 触发一个事件
func (e *Emitter) Emit(event string, data interface{}) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if listeners, found := e.listeners[event]; found {
		for _, listener := range listeners {
			go listener(data)
		}
	}
}

// 取消注册一个事件监听器
func (e *Emitter) Off(event string, listenerToRemove func(interface{})) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if listeners, found := e.listeners[event]; found {
		for i, listener := range listeners {
			if fmt.Sprintf("%p", listener) == fmt.Sprintf("%p", listenerToRemove) {
				e.listeners[event] = append(listeners[:i], listeners[i+1:]...)
				break
			}
		}
	}
}


var emitter = NewEmitter()

// func main() {
// 	emitter := NewEmitter()

// 	listener := func(data interface{}) {
// 		fmt.Println("Received:", data)
// 	}

// 	emitter.On("event1", listener)
// 	emitter.Emit("event1", "Hello, World!")

// 	emitter.Off("event1", listener)
// 	emitter.Emit("event1", "This will not be received")
// }
