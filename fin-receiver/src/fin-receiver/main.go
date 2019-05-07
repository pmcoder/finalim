package main

import (
	"fmt"
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

		listener.Send("Recv", 0)
		/*if err != nil {
			log.Error().Err(err)
		}*/
	}
}
