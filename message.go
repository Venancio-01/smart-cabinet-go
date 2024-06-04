package main

import "sync"

type MessageQueue struct {
	data  string
	mutex sync.Mutex
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		data: "",
	}
}

func (mq *MessageQueue) Push(data string) {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.data = mq.data + data
}


func (mq *MessageQueue) Get() string {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	return mq.data
}


func (mq *MessageQueue) Clear() {
	mq.mutex.Lock()
	defer mq.mutex.Unlock()
	mq.data = ""
}

var messageQueue = NewMessageQueue()
