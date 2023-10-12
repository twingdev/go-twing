package crypto

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKexKeys(t *testing.T) {
	key, err := NewKexKeys()

	assert.NoError(t, err)
	spew.Dump(key)
}
