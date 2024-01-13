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
		Name:     "TMA",
		Subjects: []string{"generated-data"},
		Storage:  jetstream.MemoryStorage,
	})

	c, _ := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "TMADataGenerator",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	msgs, _ := c.Fetch(1)
	for msg := range msgs.Messages() {
		msg.Ack()
		return msg.Data(), nil
	}
	if msgs.Error() != nil {
		fmt.Println("Error during Fetch(): ", msgs.Error())
	}
	return nil, nil
}
