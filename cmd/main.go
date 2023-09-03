package main

import (
	"context"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log.Println("INFO: Initializing the data generator service...")
	natsURI, found := os.LookupEnv("NATS_SERVER_URI")
	if !found {
		log.Fatal("ERROR: NATS server URI not set.")
	}

	log.Printf("INFO: Will connect to NATS server at %s\n", natsURI)
	nc, err := nats.Connect(natsURI)
	if err != nil {
		log.Fatalf("ERROR: failed to connect to nats server: %s", err)
	}
	defer nc.Close()

	ctx := context.Background()

	js, _ := jetstream.New(nc)

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TESTSTREAM",
		Subjects: []string{"generated-data"},
	})
	log.Println("INFO: The data generator service initialized successfully.")

	js.Publish(ctx, "generated-data", []byte("hello message"))
	log.Println("INFO: Message produced.")
}
