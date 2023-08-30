package main_test

import (
	"context"
	"strings"
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

	// TODO: this is a test driver's job
	natsServer, natsCleanup, err := nats_server.RunNATSContainer(t, ctx)
	if err != nil {
		t.Fatalf("could not initalize nats server: %s", err)
	}
	defer natsCleanup()

	// TODO: this is a test driver's job
	sut, sutCleanup, err := adapters.RunSUTContainer(t, ctx, strings.Replace(natsServer.URI, "localhost", "host.containers.internal", 1))
	if err != nil {
		t.Fatalf("could not initialize sut: %s", err)
	}
	defer sutCleanup()

	// TODO: Keep only domain-specific code in the acceptance test.
	driver := nats_client.DataGenerator{NatsServer: natsServer, NatsClient: sut}
	specifications.GenerateDataSpecification(t, &driver)
}
