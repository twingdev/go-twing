package protocols

import (
	"context"
	"encoding/base64"
	"go.uber.org/zap"
	"io"
)

type PID string

func (id PID) String() string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}

type Protocol struct {
	Ctx     context.Context
	Name    string
	Path    string
	Version string
}

type IProtocol interface {
	NameVer() (string, string)
	IProtocolHandler
	Logger() *zap.Logger
}

type (
	protocolHandler func(w io.Writer, r *io.Reader)
)

type IProtocolHandler interface {
	Handle(path string, handler protocolHandler)
}
