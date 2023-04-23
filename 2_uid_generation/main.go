package main

import (
	"encoding/json"
	"fly-maelstrom/utils"
	"log"
	"sync"

	// "github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var once sync.Once
var nodeId int64
var lock sync.Mutex
var sequence int64

func main() {
	n := maelstrom.NewNode()
	n.Handle("generate", func(msg maelstrom.Message) error {
		once.Do(func() {
			nodeId = int64(utils.GetNodeIDInt(msg.Dest))
			nodeId++
		})
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = utils.GenerateTypeOk("generate")
		var seq int64 = 0
		lock.Lock()
		sequence++
		seq = sequence
		lock.Unlock()
		uniqueID := utils.GetUniqueID(nodeId, seq)
		// could use google's uuid generator but that's no challenge then.
		// body["id"] = uuid.New().String()
		body["id"] = uniqueID
		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
