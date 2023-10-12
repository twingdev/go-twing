package crypto

import (
	"go.dedis.ch/kyber/v3"

	"go.dedis.ch/kyber/v3/sign/schnorr"
	"go.dedis.ch/kyber/v3/suites"
)

type Signer struct {
	schnorr.Suite
	privKey   kyber.Scalar
	Signature []byte
}

func (s *Signer) Sign(msg []byte) {
	suite := suites.MustFind("schnorr")
	sig, err := schnorr.Sign(suite, s.privKey, msg)
	if err != nil {
		panic(err)
	}
	s.Signature = sig

}

func (s *Signer) Verify(pub, msg, sig []byte) error {
	return schnorr.VerifyWithChecks(s, pub, msg, sig)
}

func NewSignature(privKey kyber.Scalar, msg []byte) []byte {
	s := &Signer{privKey: privKey, Signature: nil}
	s.Sign(msg)
	return s.Signature
}

func VerifySignature(pubKey, msg, sig []byte) error {
	suite := suites.MustFind("schnorr")
	s := &Signer{Suite: suite}
	return s.Verify(pubKey, msg, sig)
}
