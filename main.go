package main

import (
	"github.com/twingdev/go-twing/core"
	_ "github.com/twingdev/go-twing/core"
)

func main() {
	ds := core.NewDatastore()
	ds.DB()
}
