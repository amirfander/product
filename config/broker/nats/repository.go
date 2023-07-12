package nats

import (
	"github.com/nats-io/nats.go"
)

var connection *nats.Conn

func SetConnection(nc *nats.Conn) {
	connection = nc
}

type Nats struct{}
