package main

import (
	"encoding/json"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type BroadcastRequest struct {
	Type    string  `json:"type,omitempty"`
	Message Message `json:"message"`
}

type BroadcastResponse struct {
	Type string `json:"type,omitempty"`
}

func gossip_broadcast(n *maelstrom.Node) func(body BroadcastRequest, destination SystemNode) {
	return func(body BroadcastRequest, destination SystemNode) {
		fetching_completed := make(chan bool, 1)
		go func() {
			n.RPC(destination, BroadcastRequest{
				Type:    "broadcast",
				Message: body.Message,
			}, func(msg maelstrom.Message) error {
				fetching_completed <- true
				return nil
			})
		}()
		select {
		case <-time.After(100 * time.Millisecond):
			aBody := body
			aDestination := destination
			go func() {
				timer := time.NewTimer(100 * time.Millisecond)
				<-timer.C
				gossip_broadcast(n)(aBody, aDestination)
			}()
		case <-fetching_completed:
		}
	}
}

func Broadcast(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body BroadcastRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		saved := StoreInstance().Save(body.Message)
		if saved {
			broadcaster := gossip_broadcast(n)
			destinations := GossipTo(msg.Dest)
			for _, destination := range destinations {
				broadcaster(body, destination)
			}

		}
		response := BroadcastResponse{
			Type: "broadcast_ok",
		}

		return n.Reply(msg, response)
	}
}
