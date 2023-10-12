package core

import (
	"github.com/defval/di"
	blocks "github.com/ipfs/go-block-format"
	ipld "github.com/ipfs/go-ipld-format"
	record "github.com/libp2p/go-libp2p-record"
	"github.com/libp2p/go-libp2p/core/metrics"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	identity "github.com/twingdev/go-twing-identity"
	"github.com/twingdev/go-twing/core/repo"
)

type TwingNode struct {
	Identity peer.ID
	Repo     repo.Repo
	twingID  *identity.Identity
	Services NodeServices
}

type NodeServices struct {
	Peerstore            pstore.Peerstore          `optional:"true"` // storage for other Peer instances
	Blockstore           interface{}               // the block store (lower level)
	Filestore            interface{}               `optional:"true"` // the filestore blockstore
	BaseBlocks           blocks.BasicBlock         // the raw blockstore, no filestore wrapping
	Blocks               interface{}               // the block service, get/add blocks.
	DAG                  ipld.DAGService           // the merkle dag service, get/add objects.
	IPLDFetcherFactory   interface{}               `name:"ipldFetcher"`   // fetcher that paths over the IPLD data model
	UnixFSFetcherFactory interface{}               `name:"unixfsFetcher"` // fetcher that interprets UnixFS data
	Reporter             *metrics.BandwidthCounter `optional:"true"`
	Discovery            discovery.Service         `optional:"true"`
	FilesRoot            interface{}
	RecordValidator      record.Validator
}

func init() {
	di.SetTracer(&di.StdTracer{})
}

type Module struct {
	Di      []*di.Option
	Name    string
	Version string
}

type IModule interface {
	Version() string
	Name() string
	Load() error
	Destroy() error
}
