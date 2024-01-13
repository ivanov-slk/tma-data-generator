package main

import (
	"context"
	"log"
	"os"

	"github.com/ivanov-slk/tma-data-generator/pkg/generator"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log.Println("INFO: Initializing the data generator service...")
	natsURI, found := os.LookupEnv("NATS_SERVER_URI")
	if !found {
		log.Fatal("ERROR: NATS server URI not set.")
	}

	opts := nats.GetDefaultOptions()
	log.Printf("DEBUG: Will use the following options: %v", opts)
	log.Printf("DEBUG: Will connect to NATS server at %s\n", natsURI)
	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalf("ERROR: failed to connect to nats server: %s", err)
	}
	defer nc.Close()

	ctx := context.Background()

	js, _ := jetstream.New(nc)

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TMA",
		Subjects: []string{"generated-data"},
		Storage:  jetstream.MemoryStorage,
	})
	log.Println("INFO: The data generator service initialized successfully.")

	byteStats, _ := generator.JsonizeTemperatureStats(generator.Generate()) // TODO error handling

	js.Publish(ctx, "generated-data", byteStats)
	log.Println("INFO: Message produced.")

	// Sleep forever to make ArgoCD happy.
	select {}
}
