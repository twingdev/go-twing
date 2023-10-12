package common

type Metadata interface {
	ID() string
	GetMeta(key string, out interface{}) error
	MetaType() (mType string)
}
