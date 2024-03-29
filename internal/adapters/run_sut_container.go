package adapters

import (
	"context"
	"testing"

	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// RunSUTContainer creates a container of the data generator service
// from its Dockerfile and starts it. It is a test helper using a
// shared context.
func RunSUTContainer(t testing.TB, ctx context.Context, natsServerURI string) (*testcontainers.Container, func(), error) {
	t.Helper()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../",
			PrintBuildLog: true,
		},
		Env: map[string]string{
			"NATS_SERVER_URI": natsServerURI,
		},
		Networks: []string{
			"test-network",
		},
		NetworkAliases: map[string][]string{
			"test-network": {"data-generator"},
		},
		WaitingFor: wait.ForLog("INFO: The data generator service initialized successfully."),
	}

	sut, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("could not start data generator container: %s", err)
	}

	cleanupFunc := func() {
		if err := sut.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}

	return &sut, cleanupFunc, nil
}
