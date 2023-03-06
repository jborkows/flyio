package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type ReadRequest struct {
	Type string `json:"type,omitempty"`
}

type ReadResponse struct {
	Type     string    `json:"type,omitempty"`
	Messages []Message `json:"messages"`
}

func Read(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body ReadRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		messages := StoreInstance().ReadAll()

		response := ReadResponse{
			Messages: messages,
			Type:     "read_ok",
		}

		return n.Reply(msg, response)
	}
}
