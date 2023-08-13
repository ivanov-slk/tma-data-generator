package nats

import (
	"fmt"

	nats_server "github.com/ivanov-slk/tma-data-generator/test/specifications/nats"
)

type DataGenerator struct {
	NatsServer *nats_server.NatsContainer
}

func (dg *DataGenerator) GenerateData() ([]byte, error) {
	fmt.Printf("%v", dg.NatsServer)
	return []byte{}, nil
}
