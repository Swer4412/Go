// Importo i pacchetti necessari
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// Definisco le costanti per l'host e la porta del server
const (
	HOST = "127.0.0.1"
	PORT = "8080"
)

func main() {
	// Creo una connessione TCP con il server
	conn, err := net.Dial("tcp", HOST+":"+PORT)
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		os.Exit(1)
	}
	// Chiudo la connessione alla fine del programma
	defer conn.Close()
	// Creo un canale per sincronizzare le goroutine
	done := make(chan struct{})
	// Creo una goroutine per leggere i messaggi dal server e stamparli a schermo
	go func() {
		// Copio i dati dalla connessione allo standard output
		io.Copy(os.Stdout, conn)
		// Invio un segnale sul canale
		done <- struct{}{}
	}()
	// Creo una goroutine per leggere l'input dall'utente e inviarlo al server
	go func() {
		// Copio i dati dallo standard input alla connessione
		io.Copy(conn, os.Stdin)
		// Invio un segnale sul canale
		done <- struct{}{}
	}()
	// Attendo che una delle due goroutine termini
	<-done
	// Stampo un messaggio di fine
	fmt.Println("Il client termina.")
}
