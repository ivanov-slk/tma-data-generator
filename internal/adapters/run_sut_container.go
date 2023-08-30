package adapters

import (
	"context"
	"testing"

	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// type TestLogConsumer struct {
// 	Msgs []string
// }

// func (g *TestLogConsumer) Accept(l testcontainers.Log) {
// 	g.Msgs = append(g.Msgs, string(l.Content))
// }

// RunSUTContainer creates a container of the data generator service
// from its Dockerfile and starts it. It is a test helper using a
// shared context.
func RunSUTContainer(t testing.TB, ctx context.Context, natsServerURI string) (*testcontainers.Container, func(), error) {
	t.Helper()

	// networkName := "data-generator-network"
	// testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
	// 	NetworkRequest: testcontainers.NetworkRequest{
	// 		Driver:         "host",
	// 		Name:           networkName,
	// 		CheckDuplicate: true,
	// 		Attachable:     true,
	// 	},
	// })

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../",
			PrintBuildLog: true,
		},
		Env: map[string]string{
			"NATS_SERVER_URI": natsServerURI,
		},
		// Networks: []string{
		// 	networkName,
		// },
		WaitingFor: wait.ForLog("INFO: The data generator service initialized successfully."),
	}

	sut, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderPodman,
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("could not start data generator container: %s", err)
	}

	// g := TestLogConsumer{
	// 	Msgs: []string{},
	// }

	// sut.FollowOutput(&g) // must be called before StarLogProducer
	// err = sut.StartLogProducer(ctx)
	// if err != nil {
	// 	log.Fatalf("error creating sut logger: %s", err)
	// }

	cleanupFunc := func() {
		if err := sut.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}
	// sutnet, _ := sut.Networks(ctx)
	// fmt.Printf("sut networks: %v", sutnet)

	return &sut, cleanupFunc, nil
}
