package crypto

import (
	"go.dedis.ch/kyber/v3/suites"
	"hash"
)

func Hasher() hash.Hash {
	find, err := suites.Find("ed25519")
	if err != nil {
		return nil
	}
	return find.Hash()
}

func GetHash(value []byte) []byte {
	h := Hasher()
	h.Write(value)
	return h.Sum(nil)
}
