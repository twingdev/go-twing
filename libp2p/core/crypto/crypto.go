package crypto

import (
	crypto "github.com/libp2p/go-libp2p/core/crypto"
	pb "github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/stretchr/testify/assert"
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
	"testing"
)

var _ crypto.Ed25519PrivateKey
var _ crypto.Ed25519PublicKey
var _ crypto.PubKey

type Ed25519PrivateKey struct {
	g edwards25519.SuiteEd25519
	k kyber.Scalar
}

type Ed25519PublicKey struct {
	k kyber.Point
}

func (ed Ed25519PrivateKey) Type() pb.KeyType {
	return pb.KeyType_Ed25519
}

func (ed Ed25519PrivateKey) Raw() []byte {
	r, _ := ed.k.MarshalBinary()
	return r
}

func (ed Ed25519PrivateKey) pubKeyBytes() []byte {
	pub := ed.g.Point().Mul(ed.k, nil)
	pBytes, _ := pub.MarshalBinary()
	return pBytes
}

type PubKey interface {
	crypto.Key
	Verify(data []byte, sig []byte) (bool, error)
}

type PublicKey []byte

func (p PublicKey) Equals(key crypto.Key) bool {
	return assert.EqualValues(&testing.T{}, interface{}(p.Raw()), interface{}(key.Raw()))
}

func (p PublicKey) Raw() ([]byte, error) {
	return []byte(p), nil
}

func (p PublicKey) Type() pb.KeyType {
	return pb.KeyType_Ed25519
}

func (p PublicKey) Verify(data []byte, sig []byte) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (ed Ed25519PrivateKey) GetPublic() crypto.PubKey {
	return PublicKey(ed.pubKeyBytes())
}
