package main_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

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
	// networkName := "data-generator-network"
	// net, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
	// 	NetworkRequest: testcontainers.NetworkRequest{
	// 		Driver: "bridge",
	// 		Name:   networkName,
	// 		// CheckDuplicate: true,
	// 		// Attachable:     true,
	// 	},
	// })
	// if err != nil {
	// 	t.Fatalf("could not create network: %v", err)
	// }
	// defer func() {
	// 	_ = net.Remove(ctx)
	// }()

	// TODO: this is a test driver's job
	natsServer, natsCleanup, err := nats_server.RunNATSContainer(t, ctx)
	if err != nil {
		t.Fatalf("could not initalize nats server: %s", err)
	}
	defer natsCleanup()
	// natsnet, _ := natsServer.Networks(ctx)
	// fmt.Printf("\nnats networks: %v", natsnet)
	fmt.Printf("\nnats uri: %s", natsServer.URI)

	time.Sleep(5 * time.Second)

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
