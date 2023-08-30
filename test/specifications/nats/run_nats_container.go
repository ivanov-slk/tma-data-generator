// https://github.com/testcontainers/testcontainers-go/blob/v0.22.0/examples/nats/nats.go

package nats

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

// type TestLogConsumer1 struct {
// 	Msgs []string
// }

// func (g *TestLogConsumer1) Accept(l testcontainers.Log) {
// 	g.Msgs = append(g.Msgs, string(l.Content))
// }

// NatsContainer represents the nats container type used in the module
type NatsContainer struct {
	testcontainers.Container
	URI string
}

// RunNATSContainer creates an instance of the nats container type
func RunNATSContainer(t testing.TB, ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*NatsContainer, func(), error) {
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
		Image:        "nats:alpine",
		Hostname:     "127.0.0.1",
		ExposedPorts: []string{"4222/tcp", "6222/tcp", "8222/tcp"},
		// Networks: []string{
		// 	networkName,
		// },
		Cmd:        []string{"-DV", "-js"},
		WaitingFor: wait.ForLog("Listening for client connections on 0.0.0.0:4222"),
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ProviderType:     testcontainers.ProviderPodman,
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		opt.Customize(&genericContainerReq)
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	if err != nil {
		return nil, nil, err
	}

	// g := TestLogConsumer1{
	// 	Msgs: []string{},
	// }

	// container.FollowOutput(&g) // must be called before StarLogProducer
	// err = container.StartLogProducer(ctx)
	// if err != nil {
	// 	log.Fatalf("error creating nats logger: %s", err)
	// }

	mappedPort, err := container.MappedPort(ctx, "4222/tcp")
	if err != nil {
		return nil, nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, nil, err
	}

	uri := fmt.Sprintf("nats://%s:%s", hostIP, mappedPort.Port())

	log.Printf("----------------------- %s", uri)

	natsContainer := &NatsContainer{Container: container, URI: uri}

	cleanupFunc := func() {
		if err := natsContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}

	// nc, err := nats.Connect(uri)
	// if err != nil {
	// 	log.Fatalf("failed to connect to nats: %s", err)
	// }
	// defer nc.Close()

	// js, _ := jetstream.New(nc)
	// log.Println("INFO: Created new jetstream instance.")

	// js.CreateStream(ctx, jetstream.StreamConfig{
	// 	Name:     "TESTSTREAM",
	// 	Subjects: []string{"generated-data"},
	// })
	// log.Println("INFO: The data generator service initialized successfully.")

	// js.Publish(ctx, "generated-data", []byte("hello message"))
	// log.Println("INFO: Message produced.")

	return natsContainer, cleanupFunc, nil
}
