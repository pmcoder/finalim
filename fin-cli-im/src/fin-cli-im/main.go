package main

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	//Listen on socket with zmq
	//TODO move this to an interfaced struct

	client, err := zmq.NewSocket(zmq.REQ)
	if err != nil {
		panic(fmt.Sprintf("Could not create zmq socket: %s", err))
	}
	defer client.Close()
	client.Connect("tcp://localhost:5555")

	input := os.Args[1]

	client.Send(input, 0)
	log.Info().Str("message", input).Msg("Message sent")
	msg, err := client.Recv(0)
	if err != nil {
		log.Error().Err(err).Msg("Error receiving acknowledgement")
	}

	log.Info().Msgf("Received acknowledgement: %s", msg)

}
