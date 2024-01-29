package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Genera un ID univoco per l'ordine (per semplicità qui utilizza un timestamp)
	orderID := fmt.Sprintf("Order_%d", time.Now().Unix())

	// Pubblica la richiesta per un ordine
	err = nc.Publish("newOrder", []byte(orderID))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Order placed successfully. Order ID: %s\n", orderID)
}
