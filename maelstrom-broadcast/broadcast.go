package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastRequest struct {
	Type    string  `json:"type,omitempty"`
	Message Message `json:"message"`
}

type BroadcastResponse struct {
	Type string `json:"type,omitempty"`
}

func Broadcast(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body BroadcastRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		StoreInstance().Save(body.Message)

		response := BroadcastResponse{
			Type: "broadcast_ok",
		}

		return n.Reply(msg, response)
	}
}
