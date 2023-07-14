package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"

	"product/repository/broker"
)

type NatsConfig struct {
}

func (nc NatsConfig) ConnectBroker(uri string) {
	ncon, err := nats.Connect(uri)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to Nats")
	SetConnection(ncon)
	broker.SetRepository(Nats{})
}
