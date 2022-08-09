package natsStreaming

import (
	"encoding/json"
	"fmt"
	"log"

	"L0/data/model"
	"L0/data/pg"
	"L0/data/redis"

	"github.com/nats-io/stan.go"
)

type NatsReceiver struct {
	sc stan.Conn

	sub stan.Subscription
}

func NewNatsReceiver() (*NatsReceiver, error) {
	sc, err := stan.Connect("test-cluster", "myID")
	if err != nil {
		log.Fatal(err)
	}

	return &NatsReceiver{
		sc: sc,
	}, nil
}

func (n *NatsReceiver) Start() (err error) {
	n.sub, err = n.sc.Subscribe("order", processOrder, stan.StartWithLastReceived())
	return err
}

func (n *NatsReceiver) Stop() {
	// Unsubscribe
	err := n.sub.Unsubscribe()
	if err != nil {
		fmt.Println(err)
	}

	// Close connection
	err = n.sc.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func processOrder(m *stan.Msg) {
	var orderModel model.Order

	err := json.Unmarshal(m.Data, &orderModel)
	if err != nil {
		fmt.Printf("error while unmarshaling order: %s\n", err)
		return
	}

	err = redis.Set(orderModel.OrderUid, string(m.Data))
	if err != nil {
		fmt.Printf("error while caching order: %s\n", err)
	}

	err = pg.AddOrder(orderModel)
	if err != nil {
		fmt.Printf("error while saving to postgres: %s\n", err)
	}

}
