package crypt

import (
	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

type KeyExchange struct {
	public    kyber.Point
	otherKey  kyber.Point
	challenge []byte
}

type KexKeys struct {
	Public  kyber.Point
	Private kyber.Scalar
	Group   kyber.Group
}

func NewKexKeys() (*KexKeys, error) {
	var g edwards25519.SuiteEd25519
	kg := edwards25519.NewBlakeSHA256Ed25519()

	sk := g.NewKey(kg.RandomStream())
	pub := g.Point().Pick(kg.RandomStream()).Mul(sk, nil)
	return &KexKeys{Public: pub, Private: sk, Group: &g}, nil
}

func (kk *KexKeys) GetChallenge(otherKey *KexKeys) *KeyExchange {

	chBytes, _ := kk.Group.Point().Mul(kk.Private, otherKey.Public).MarshalBinary()
	return &KeyExchange{
		public:    kk.Public,
		otherKey:  otherKey.Public,
		challenge: chBytes,
	}
}
