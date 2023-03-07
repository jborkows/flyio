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

		saved := StoreInstance().Save(body.Message)
		if saved {
			destinations := GossipTo(msg.Dest)
			for i := 0; i < len(destinations); i++ {
				n.RPC(destinations[i], BroadcastRequest{
					Type:    "broadcast",
					Message: body.Message,
				}, func(msg maelstrom.Message) error {
					return nil
				})

			}
		}
		response := BroadcastResponse{
			Type: "broadcast_ok",
		}

		return n.Reply(msg, response)
	}
}
