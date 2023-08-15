package nats

import (
	"fmt"

	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
	"github.com/testcontainers/testcontainers-go"
)

type DataGenerator struct {
	NatsServer *nats_server.NatsContainer
	NatsClient *testcontainers.Container
}

func (dg *DataGenerator) GenerateData() ([]byte, error) {
	fmt.Printf("%v", dg.NatsServer)
	return []byte{}, nil
}
