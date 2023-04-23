package main

import (
	"encoding/json"
	"fly-maelstrom/utils"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var lock sync.Mutex
var msgs = map[int]struct{}{}

func main() {
	n := maelstrom.NewNode()
	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		lock.Lock()
		msgs[int(body["message"].(float64))] = struct{}{}
		lock.Unlock()
		return n.Reply(msg, utils.GetTypeOkBody("broadcast"))
	})
	n.Handle("read", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		messages := make([]int, 0, len(msgs))
		for k := range msgs {
			messages = append(messages, k)
		}
		lock.Lock()
		body["messages"] = messages
		lock.Unlock()
		body["type"] = utils.GenerateTypeOk("read")
		return n.Reply(msg, body)
	})
	n.Handle("topology", func(msg maelstrom.Message) error {
		return n.Reply(msg, utils.GetTypeOkBody("topology"))
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
