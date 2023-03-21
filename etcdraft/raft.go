// This is a test base package.
package etcdraft

import (
	"go.etcd.io/etcd/raft/v3"
)

func MyRaft() {
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}
	//raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})

	peers := []raft.Peer{{ID: 0x01}}
	raft.StartNode(c, peers)

}
