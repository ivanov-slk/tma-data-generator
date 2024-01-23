package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/ivanov-slk/tma-data-generator/pkg/generator"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log.Println("INFO: Initializing the data generator service...")
	natsURI, found := os.LookupEnv("NATS_SERVER_URI")
	if !found {
		slog.Error("NATS server URI not set.")
		os.Exit(1)
	}

	nc, err := nats.Connect(natsURI)
	if err != nil {
		slog.Error("failed to connect to nats server:", "error", err)
		os.Exit(1)
	}
	defer nc.Close()

	ctx := context.Background()

	js, _ := jetstream.New(nc)

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TMA",
		Subjects: []string{"generated-data"},
		Storage:  jetstream.MemoryStorage,
	})
	log.Println("INFO: The data generator service initialized successfully.") // TODO: can't use slog because of the test

	for {
		byteStats, err := generator.JsonizeTemperatureStats(generator.Generate())
		if err != nil {
			slog.Error("failed generating message bytes:", "error", err)
		}

		if nc.Status() != nats.CONNECTED {
			continue
		}

		if _, err := js.Publish(ctx, "generated-data", byteStats); err != nil {
			slog.Error("failed publishing a message:", "error", err)
		}
		time.Sleep(60 * time.Second)
	}
}
