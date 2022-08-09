package main

import (
	"log"

	"L0/transport"
	"L0/transport/natsStreaming"
)

func main() {

	nats, err := natsStreaming.NewNatsReceiver()
	if err != nil {
		log.Fatalln(err)
	}

	err = nats.Start()
	if err != nil {
		log.Fatalln(err)
	}
	defer nats.Stop()

	transport.NewServer()
}
