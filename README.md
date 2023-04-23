## Fly.io Distributed Systems Challenges

A repo hosting solutions for [Distributed Systems](https://fly.io/dist-sys) challenges from [fly.io](https://fly.io).

Challenges-

* Echo [x] - `./maelstrom test -w echo --bin /Users/vishal/work/src/github.com/vishal1132/fly-maelstrom/1_echo/main --node-count 1 --time-limit 10 `
* Uid Generator [x] - `./maelstrom test -w unique-ids --bin /Users/vishal/work/src/github.com/vishal1132/fly-maelstrom/2_uid_generation/main --node-count 30 --time-limit 30 --availability total --nemesis partition`
* Single Node Broadcast [x] - `./maelstrom test -w broadcast --bin /Users/vishal/work/src/github.com/vishal1132/fly-maelstrom/3a_broadcast/main --node-count 1 --time-limit 20 --rate 10 `