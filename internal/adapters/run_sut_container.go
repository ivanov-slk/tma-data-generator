package adapters

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// RunSUTContainer creates a container of the data generator service
// from its Dockerfile and starts it. It is a test helper using a
// shared context.
func RunSUTContainer(t testing.TB, ctx context.Context) (*testcontainers.Container, func(), error) {
	t.Helper()
	log.Println("1")
	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../",
			PrintBuildLog: true,
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
