package main_test

import (
	"context"
	"testing"

	"github.com/ivanov-slk/tma-data-generator/internal/adapters"
	nats_client "github.com/ivanov-slk/tma-data-generator/internal/adapters/nats"
	"github.com/ivanov-slk/tma-data-generator/test/specifications"
	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func TestNatsClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	// Create network for the containers
	// aliases := []string{"alias1", "alias2", "alias3"}
	networkName := "ttt"
	newNetwork, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           networkName,
			CheckDuplicate: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		require.NoError(t, newNetwork.Remove(ctx))
	})

	// TODO: this is a test driver's job
	natsServer, natsCleanup, err := nats_server.RunNATSContainer(t, ctx)
	if err != nil {
		t.Fatalf("could not initalize nats server: %s", err)
	}
	defer natsCleanup()

	// TODO: this is a test driver's job
	// sut, sutCleanup, err := adapters.RunSUTContainer(t, ctx, strings.Replace(natsServer.URI, "localhost", "nats-server", 1))
	// sut, sutCleanup, err := adapters.RunSUTContainer(t, ctx, natsServer.URI)
	sut, sutCleanup, err := adapters.RunSUTContainer(t, ctx, "http://nats-server:4222")
	if err != nil {
		t.Fatalf("could not initialize sut: %s", err)
	}
	defer sutCleanup()

	// TODO: Keep only domain-specific code in the acceptance test.
	driver := nats_client.DataGenerator{NatsServer: natsServer, NatsClient: sut}
	specifications.GenerateDataSpecification(t, &driver)
}
