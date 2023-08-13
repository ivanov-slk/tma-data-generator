package main_test

import (
	"context"
	"testing"

	nats_client "github.com/ivanov-slk/tma-data-generator/internal/adapters/nats"
	"github.com/ivanov-slk/tma-data-generator/test/specifications"
	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
)

func TestNatsClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()
	natsServer, cleanupFunc, err := nats_server.RunContainer(ctx)
	if err != nil {
		t.Fatalf("could not initalize nats server: %s", err)
	}
	defer cleanupFunc()

	driver := nats_client.DataGenerator{NatsServer: natsServer}
	specifications.GenerateDataSpecification(t, &driver)

}
