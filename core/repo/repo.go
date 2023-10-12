package repo

import (
	"github.com/ipfs/go-datastore"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/twingdev/go-twing/core/config"
	"io"
)

type Repo interface {
	Config() (*config.Config, error)
	SetConfig(*config.Config) error
	SetConfigKey(key string, value interface{}) error
	GetConfigKey(key string) (interface{}, error)
	Datastore() Datastore
	Keystore() interface{}
	FileManager() interface{}
	SetAPIAddr(addr ma.Multiaddr) error

	// SwarmKey returns the configured shared symmetric key for the private networks feature.
	SwarmKey() ([]byte, error)

	io.Closer
}

type Datastore interface {
	datastore.Batching
}
