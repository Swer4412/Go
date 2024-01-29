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

	// Sottoscrizione per ricevedere le notifiche degli ordini processati
	nc.Subscribe("shippingProcessed", func(msg *nats.Msg) {
		shippingID := string(msg.Data)
		fmt.Printf("Received shipping processed notification: %s\n", shippingID)

		// Simula il pagamento dell'ordine
		processPayment(shippingID)
	})

	log.Println("Payment Service is running. Waiting for order processed notifications...")
	select {}
}

func processPayment(shippingID string) {
	// Simulazione del pagamento
	time.Sleep(2 * time.Second)
	fmt.Printf("Payment processed for order: %s\n", shippingID)
}
