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

	// Sottoscrizione per ricevedere le notifiche degli ordini richiesti
	nc.Subscribe("newOrder", func(msg *nats.Msg) {
		orderID := string(msg.Data)
		fmt.Printf("Received new order request: %s\n", orderID)

		// Simula la richiesta dell'ordine
		processOrder(orderID)

		// Notifica che l'ordine è stato processato
		nc.Publish("orderProcessed", []byte(orderID))
	})

	log.Println("Order Service is running. Waiting for new orders...")
	select {}
}

func processOrder(orderID string) {
	// Simulazione dell'esecuzione dell'ordine
	time.Sleep(5 * time.Second)
	fmt.Printf("Order processed: %s\n", orderID)
}
