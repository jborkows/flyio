package main

type Message = int64

type Store interface {
	Save(message Message)
	ReadAll() []Message
}

type storeImpl struct {
	messages []Message
}

var store *storeImpl = initValue()

func initValue() *storeImpl {
	var store *storeImpl = new(storeImpl)
	store.messages = make([]Message, 0)
	return store
}

func StoreInstance() Store {
	return store
}

func (self *storeImpl) Save(message Message) {
	self.messages = append(self.messages, message)
}

func (self *storeImpl) ReadAll() []Message {
	return self.messages
}
