package rmq

import "github.com/streadway/amqp"

type Rmq interface {
	Connect()
	Disconnect()
	GetChannel() *amqp.Channel
}
