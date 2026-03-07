package main

import (
	"context"
	"log"

	"KVStore/internal/network"
)

func main() {
	ctx := context.Background()

	t, err := network.NewLibp2pTransport(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Start(); err != nil {
		log.Fatal(err)
	}

	select {}
}
