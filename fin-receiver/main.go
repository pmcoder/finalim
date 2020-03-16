package main

import (
	"encoding/json"
	"fmt"
	nats "github.com/nats-io/nats.go"
	zmq "github.com/pebbe/zmq4"
	"github.com/rs/zerolog/log"
)

func main() {
	//Listen on socket with zmq
	//TODO move this to an interfaced struct
	log.Info().Msg("Starting fin receiver service")
	listener, err := zmq.NewSocket(zmq.REP)
	if err != nil {
		panic(fmt.Sprintf("Could not create zmq socket: %s", err))
	}
	defer listener.Close()

	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(fmt.Sprintf("Could not connect to NATS %s", err))
	}

	listener.Bind("tcp://*:5555")
	log.Info().Msgf("Listening on port %d", 5555)
	for {
		msg, err := listener.Recv(0)
		if err != nil {
			log.Error().Err(err).Msg("Failed to receive message")
		}

		//TODO use interface or callback func here to handle the message
		//will probably be put on queue by type
		//for not just print it
		log.Info().Msgf("Received msg: %s", msg)

		err = enqueueMeessage(Message{}, natsConn)
		if err != nil {
			log.Error().Err(err).Msg("Failed to enqueue message")
		}

		listener.Send("Recv", 0)
		/*if err != nil {
			log.Error().Err(err)
		}*/
	}
}

func enqueueMeessage(message Message, nc *nats.Conn) error {
	log.Info().Msg("Enqueueing message")
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = nc.Publish("messages-incoming", data)
	if err != nil {
		return err
	}
	log.Info().Msg("Enqueued message")
	return nil

}

type Message struct {
	Id          string
	Content     string
	Destination string
}
