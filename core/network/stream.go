package network

import (
	ifconnmgr "github.com/libp2p/go-libp2p/core/connmgr"
	net "github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	ma "github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"io"
	"sync"
)

const cmgrTag = "stream-fwd"

type Stream struct {
	id uint64

	Protocol protocol.ID

	OriginAddr ma.Multiaddr
	TargetAddr ma.Multiaddr
	peer       peer.ID

	Local  manet.Conn
	Remote net.Stream

	Registry *StreamRegistry
}

type StreamRegistry struct {
	sync.Mutex

	Streams map[uint64]*Stream // map of stream id to *Stream
	conns   map[peer.ID]int
	nextID  uint64
	streams map[string]*Stream // map of remote addr to *Stream

	ifconnmgr.ConnManager
}

// close stream endpoints and deregister it
func (s *Stream) close() {
	s.Registry.Close(s)
}

// reset closes stream endpoints and deregisters it
func (s *Stream) reset() {
	s.Registry.Reset(s)
}

func (s *Stream) startStreaming() {
	go func() {
		_, err := io.Copy(s.Local, s.Remote)
		if err != nil {
			s.reset()
		} else {
			s.close()
		}
	}()

	go func() {
		_, err := io.Copy(s.Remote, s.Local)
		if err != nil {
			s.reset()
		} else {
			s.close()
		}
	}()
}

// Register registers a stream to the registry
func (r *StreamRegistry) Register(streamInfo *Stream) {
	r.Lock()
	defer r.Unlock()

	r.ConnManager.TagPeer(streamInfo.peer, cmgrTag, 20)
	r.conns[streamInfo.peer]++

	streamInfo.id = r.nextID
	r.Streams[r.nextID] = streamInfo
	r.nextID++

	r.streams[streamInfo.Local.LocalAddr().String()] = streamInfo

	streamInfo.startStreaming()
}

// Deregister deregisters stream from the registry
func (r *StreamRegistry) Deregister(streamID uint64) {
	r.Lock()
	defer r.Unlock()

	s, ok := r.Streams[streamID]
	if !ok {
		return
	}
	p := s.peer
	r.conns[p]--
	if r.conns[p] < 1 {
		delete(r.conns, p)
		r.ConnManager.UntagPeer(p, cmgrTag)
	}

	delete(r.Streams, streamID)

	delete(r.streams, s.Local.LocalAddr().String())
}

// GetStreamRemotePeerID looks up the remote's peer ID based on local open address
// Note `addr` is `RemoteAddr` from handler's context
func (r *StreamRegistry) GetStreamRemotePeerID(addr string) (peer.ID, bool) {
	r.Lock()
	defer r.Unlock()

	if s, ok := r.streams[addr]; ok {
		return s.peer, true
	}
	return "", false
}

// Close stream endpoints and deregister it
func (r *StreamRegistry) Close(s *Stream) {
	_ = s.Local.Close()
	_ = s.Remote.Close()
	s.Registry.Deregister(s.id)
}

// Reset closes stream endpoints and deregisters it
func (r *StreamRegistry) Reset(s *Stream) {
	_ = s.Local.Close()
	_ = s.Remote.Reset()
	s.Registry.Deregister(s.id)
}
