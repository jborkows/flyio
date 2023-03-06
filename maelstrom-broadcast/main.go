package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	n.Handle("topology", Topology(n))
	n.Handle("broadcast", Broadcast(n))
	n.Handle("read", Read(n))
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
