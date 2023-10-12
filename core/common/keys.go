package common

import (
	pb "github.com/twingdev/go-twing/core/protos"
)

type TwingKey pb.Key

func (t *TwingKey) Compare(otherKey *TwingKey) bool {
	return t.Public == otherKey.Public
}

func (t *TwingKey) GetType() pb.KeyTypes {
	return t.Type
}
