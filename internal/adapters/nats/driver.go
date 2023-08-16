package nats

import (
	"context"
	"log"
	"time"

	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
	"github.com/nats-io/nats.go"
	"github.com/testcontainers/testcontainers-go"
)

type DataGenerator struct {
	NatsServer *nats_server.NatsContainer
	NatsClient *testcontainers.Container
}

func (dg *DataGenerator) GenerateData() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	nc, err := nats.Connect(dg.NatsServer.URI)
	if err != nil {
		log.Fatalf("failed to connect to nats: %s", err)
	}
	defer nc.Close()

	// perform assertions
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("failed to create jetstream context: %s", err)
	}

	// add stream to nats
	if _, err = js.AddStream(&nats.StreamConfig{
		Name:     "test-data-generator",
		Subjects: []string{"generated-data"},
	}); err != nil {
		log.Fatalf("failed to add stream: %s", err)
	}

	// add subscriber to nats
	sub, err := js.SubscribeSync("generated-data", nats.Durable("worker"))
	if err != nil {
		log.Fatalf("failed to subscribe to hello: %s", err)
	}

	// wait for the message to be received
	msg, err := sub.NextMsgWithContext(ctx)
	if err != nil {
		log.Fatalf("failed to get message: %s", err)
	}

	return msg.Data, nil
}
