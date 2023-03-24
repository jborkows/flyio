package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type AddRequest struct {
	Type  string `json:"type,omitempty"`
	Delta int    `json:"delta"`
}

type AddResponse struct {
	Type string `json:"type,omitempty"`
}

func Add(n *maelstrom.Node, repository *Repository) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body AddRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		repository.addDelta(body.Delta)
		response := AddResponse{
			Type: "add_ok",
		}

		return n.Reply(msg, response)
	}
}

type ReadRequest struct {
	Type string `json:"type,omitempty"`
}

type ReadResponse struct {
	Type  string `json:"type,omitempty"`
	Value int    `json:"value"`
}

func Read(n *maelstrom.Node, kv *Repository) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body ReadRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		response := ReadResponse{
			Type:  "read_ok",
			Value: kv.read(),
		}

		return n.Reply(msg, response)
	}
}

func main() {
	node := maelstrom.NewNode()
	kv := maelstrom.NewSeqKV(node)
	repository := Repository{
		kv,
	}
	node.Handle("add", Add(node, &repository))
	node.Handle("read", Read(node, &repository))
	if err := node.Run(); err != nil {
		log.Fatal(err)
	}

}
