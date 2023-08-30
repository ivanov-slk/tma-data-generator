package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/testcontainers/testcontainers-go"
)

type DataGenerator struct {
	NatsServer *nats_server.NatsContainer
	NatsClient *testcontainers.Container
}

func (dg *DataGenerator) GenerateData() ([]byte, error) {
	nc, err := nats.Connect(dg.NatsServer.URI)
	if err != nil {
		log.Fatalf("failed to connect to nats: %s", err)
	}
	defer nc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	js, _ := jetstream.New(nc)

	s, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TESTSTREAM",
		Subjects: []string{"generated-data"},
	})

	// s, _ := js.Stream(ctx, "TESTSTREAM")
	fmt.Printf("\n\n=================================%v", s)

	js.Publish(ctx, "generated-data", []byte("hello message"))
	fmt.Printf("Published hello message\n")

	c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "TESTCONSUMER",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	msgs, _ := c.Fetch(1)
	fmt.Printf("\n\n>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>%v\n\n", msgs)
	for msg := range msgs.Messages() {
		fmt.Printf("\n\n=================================%v", msg)
		msg.Ack()
		fmt.Printf("Received a JetStream message via fetch: %s\n", string(msg.Data()))
		return msg.Data(), nil
	}
	if msgs.Error() != nil {
		fmt.Println("Error during Fetch(): ", msgs.Error())
	}
	return nil, nil
}
