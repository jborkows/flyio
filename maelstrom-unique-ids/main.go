package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Request struct {
	Type  string `json:"type,omitempty"`
	MsgId int64  `json:"msg_id,omitempty"`
}

type Response struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type Id = string

var counter uint64

func generateId(msg maelstrom.Message) Id {
	hostname, error := os.Hostname()
	if error != nil {
		panic("AAAAA")
	}
	atomic.AddUint64(&counter, 1)
	return fmt.Sprintf("%s-%d-%v-%v", hostname, counter, msg.Src, msg.Dest)
}

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body Request
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		log.Printf("Got %v", body)

		response := Response{
			Type: "generate_ok",
			ID:   generateId(msg),
		}

		return n.Reply(msg, response)
		// Echo the original message back with the updated message type.
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
