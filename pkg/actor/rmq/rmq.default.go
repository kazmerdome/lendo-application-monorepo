package rmq

import (
	"log"

	"github.com/streadway/amqp"
)

type rmq struct {
	uri        string
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRmq(uri string) *rmq {
	return &rmq{
		uri: uri,
	}
}

func (r *rmq) Connect() {
	AMQPconn, err := amqp.Dial(r.uri)
	if err != nil {
		log.Panic(err)
	}

	chnl, err := AMQPconn.Channel()
	if err != nil {
		log.Panic(err)
	}

	r.channel = chnl
	r.connection = AMQPconn

	log.Print("connected to the AMQP")
}

func (r *rmq) Disconnect() {
	log.Println("closing rmq connection!")

	r.connection.Close()

	log.Println("successfully closed!")
}

func (r *rmq) GetChannel() *amqp.Channel {
	return r.channel
}
