package main

import (
	"encoding/json"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type SystemNode = string

type TopologyRequest struct {
	Type    string                      `json:"type,omitempty"`
	Toplogy map[SystemNode][]SystemNode `json:"topology"`
}

type TopologyResponse struct {
	Type string `json:"type,omitempty"`
}

var topologyStore map[SystemNode][]SystemNode = make(map[SystemNode][]SystemNode)

func GossipTo(systemNode SystemNode) []SystemNode {
	return topologyStore[systemNode]
}

func Topology(n *maelstrom.Node) func(msg maelstrom.Message) error {
	return func(msg maelstrom.Message) error {
		var body TopologyRequest
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		topologyStore = body.Toplogy
		response := TopologyResponse{
			Type: "topology_ok",
		}

		return n.Reply(msg, response)
	}
}
