package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

var connection *nats.Conn

func SetConnection(nc *nats.Conn) {
	connection = nc
}

type Nats struct{}

func (n Nats) Publish(subject string, data []byte) {
	if err := connection.Publish(subject, data); err != nil {
		fmt.Println(err)
	}
	fmt.Println("publish message")
}
