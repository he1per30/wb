package natsStreaming

import (
	"encoding/json"
	"log"

	"L0/data/model"

	"github.com/nats-io/stan.go"
)

//func Test() {
//	sc, err := stan.Connect("test-cluster", "myID")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	b := []byte(model.TestData)
//	var r model.Order
//
//	_ = json.Unmarshal(b, &r)
//
//	// Simple Synchronous Publisher
//	sc.Publish("foo", []byte("Hello World"))
//	sc.Publish("foo", []byte("Hello World2")) // does not return until an ack has been received from NATS Streaming
//	// Simple Async Subscriber
//	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
//		fmt.Printf("Received a message: %s\n", string(m.Data))
//	}, stan.DeliverAllAvailable())
//
//	// Unsubscribe
//	sub.Unsubscribe()
//
//	// Close connection
//	sc.Close()
//}
//
//func SubscribeChan() {
//	in := make(chan string)
//	sc, err := stan.Connect("test-cluster", "myID")
//	if err != nil {
//		log.Fatal(err)
//	}
//	var orderModel model.Order
//
//	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
//		in <- string(m.Data)
//	}, stan.StartWithLastReceived())
//
//	stringData := <-in
//
//	err = json.Unmarshal([]byte(stringData), &orderModel)
//
//	if err != nil {
//		log.Fatalf("error: %s", err)
//		return
//	}
//
////	pg.AddOrder(&orderModel)
//
//	// Unsubscribe
//	sub.Unsubscribe()
//
//	// Close connection
//	sc.Close()
//}

func PublicInChanel() {
	sc, err := stan.Connect("test-cluster", "myID2")

	if err != nil {
		log.Fatal(err)
	}

	b := []byte(model.TestData)
	var r model.Order

	_ = json.Unmarshal(b, &r)

	sc.Publish("order", b)

	sc.Close()
}
