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
	nc.Subscribe("orderProcessed", func(msg *nats.Msg) {
		orderID := string(msg.Data)
		fmt.Printf("Received new shipping request: %s\n", orderID)

		// Simula la richiesta dell'ordine
		processShipping(orderID)

		// Notifica che l'ordine è stato processato
		nc.Publish("shippingProcessed", []byte(orderID))
	})

	log.Println("Shipping Service is running. Waiting for new shippings...")
	select {}
}

func processShipping(orderID string) {
	// Simulazione dell'esecuzione dell'ordine
	time.Sleep(5 * time.Second)
	fmt.Printf("Shipping processed: %s\n", orderID)
}
