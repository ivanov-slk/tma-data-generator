package main_test

import (
	"context"
	"testing"

	"github.com/ivanov-slk/tma-data-generator/internal/adapters"
	nats_client "github.com/ivanov-slk/tma-data-generator/internal/adapters/nats"
	"github.com/ivanov-slk/tma-data-generator/test/specifications"
	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
)

func TestNatsClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	natsServer, natsCleanup, err := nats_server.RunContainer(t, ctx)
	if err != nil {
		t.Fatalf("could not initalize nats server: %s", err)
	}
	defer natsCleanup()

	sut, sutCleanup, err := adapters.StartSUTContainer(t, ctx)
	if err != nil {
		t.Fatalf("could not initialize sut: %s", err)
	}
	defer sutCleanup()

	driver := nats_client.DataGenerator{NatsServer: natsServer, NatsClient: sut}
	specifications.GenerateDataSpecification(t, &driver)
}
