package network

import (
	"context"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	identity "github.com/twingdev/go-twing-identity"
)

type Network struct {
	identity *identity.Identity
	ctx      context.Context
	peers    map[string]*Node
	Node     *Node
}

type Node struct {
	ID              peer.ID
	Host            host.Host
	ListenersLocal  *Listeners
	ListenersRemote *Listeners
	Streams         *StreamRegistry
	peerStore       pstore.Peerstore
}

func NewNode(id peer.ID, peerHost host.Host, peerStore pstore.Peerstore) *Node {
	return &Node{
		ID:              id,
		Host:            peerHost,
		peerStore:       peerStore,
		ListenersLocal:  newListenersLocal(),
		ListenersRemote: newListenersRemote(peerHost),
		Streams: &StreamRegistry{
			Streams:     map[uint64]*Stream{},
			ConnManager: peerHost.ConnManager(),
			conns:       map[peer.ID]int{},
			streams:     map[string]*Stream{},
		},
	}
}

func (p2p *Node) CheckProtoExists(proto string) bool {
	protos := p2p.Host.Mux().Protocols()

	for _, p := range protos {
		if string(p) != proto {
			continue
		}
		return true
	}
	return false
}
