package main

// client
import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	connHost = "127.0.0.1" // host server (se il server è in locale, cambiate con "localhost")
	connPort = "8080"          // porta
	connType = "tcp"           // tipo connessione
)

func main() {
	fmt.Println("Connessione a " + connType + " server " + connHost + ":" + connPort + " (scrivere \"esci\" per terminare)")

	// net.Dial effettua la connessione
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Errore di connessione:", err.Error())
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		testo, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Errore in creazione buffer di invio:", err.Error())
			os.Exit(1)
		}

		// Invia la stringa al server
		fmt.Fprintf(conn, testo+"\n")

		// Controlla la richiesta di chiusura della connessione
		if strings.TrimSpace(string(testo)) == "0" { // scrivendo solo "esci" si termina il programma
			fmt.Println("Client terminato.")
			return
		}

		// Ricezione del messaggio dal server
		messaggio, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Errore in creazione buffer di ricezione:", err.Error())
			os.Exit(1)
		}
		fmt.Print(conn.RemoteAddr().String() + " -> " + messaggio)
	}
}
