package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Repository struct {
	*maelstrom.KV
}

func PrintInfo(msg string) {
	fmt.Fprintf(os.Stderr, "%v %s\n", time.Now(), msg)
}

var key = "key"

func (self *Repository) addDelta(delta int) {
	if delta == 0 {
		return
	}

	self.Write(context.Background(), fmt.Sprintf("%v AAAAA", rand.Float64()), "dupa")

	value, err := self.tryReading()
	if err != nil {
		go func() {
			time.Sleep(time.Duration(50) * time.Millisecond)
			self.addDelta(delta)
		}()
		return
	}

	err = self.KV.CompareAndSwap(context.Background(), key, value, value+delta, true)
	if err != nil {
		go func() {
			time.Sleep(time.Duration(50) * time.Millisecond)
			self.addDelta(delta)
		}()
		return
	}

}
func (self *Repository) tryReading() (int, error) {
	value, err := self.KV.ReadInt(context.Background(), key)
	if rpc, ok := err.(*maelstrom.RPCError); ok {
		if rpc.Code == maelstrom.KeyDoesNotExist {
			err = self.KV.CompareAndSwap(context.Background(), key, 0, 0, true)
			if err != nil {
				return 0, err
			} else {
				return 0, nil
			}
		}
	}
	return value, err
}

func (self *Repository) read() int {
	value, err := self.tryReading()
	if err != nil {
		return 0
	}
	return value
}
