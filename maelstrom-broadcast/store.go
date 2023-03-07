package main

import (
	"sync"

	"golang.org/x/exp/maps"
)

type Message = int64

type Store interface {
	Save(message Message) bool
	ReadAll() []Message
}
type void struct{}

var mu sync.Mutex

var member void

type storeImpl struct {
	messages map[Message]void
}

var store *storeImpl = initValue()

func initValue() *storeImpl {
	var store *storeImpl = new(storeImpl)
	store.messages = make(map[Message]void)
	return store
}

func StoreInstance() Store {
	return store
}

func (self *storeImpl) Save(message Message) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := self.messages[message]; ok {
		return false
	} else {
		self.messages[message] = member
		return true
	}
}

func (self *storeImpl) ReadAll() []Message {
	return maps.Keys(self.messages)
}
