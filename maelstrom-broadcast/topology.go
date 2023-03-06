package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type TopologyRequest struct {
	Type    string              `json:"type,omitempty"`
	Toplogy map[string][]string `json:"topology"`
}

type TopologyResponse struct {
	Type string `json:"type,omitempty"`
}

func Topology(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body TopologyRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		response := TopologyResponse{
			Type: "topology_ok",
		}

		return n.Reply(msg, response)
	}
}
