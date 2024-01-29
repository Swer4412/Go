package main

// server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	connHost = "localhost" // host locale (server)
	connPort = "8080"      // porta
	connType = "tcp"       // volendo essere più specifici, usare "tcp4" (IPv4-only)
)

//Un socket e un indirizzo ip e una porta associata

func main() {
	fmt.Println("Avvio server " + connType + " su " + connHost + ":" + connPort + ".")

	// net.Listen specifica come il processo si dovrà annunciare (in che modo accetterà le connessioni)
	l, err := net.Listen(connType, ":"+connPort)
	if err != nil {
		fmt.Println("Errore connessione:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	for { // ripete all'infinito per accettare più volte le connessioni
		fmt.Println("\nServer pronto e in attesa di connessioni")

		// l.Accept fa in modo che il Listener l si blocchi in attesa di connessioni
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Errore di connessione:", err.Error())
			return
		}
		fmt.Println("Client " + conn.RemoteAddr().String() + " connesso.")

		//Main loop
		continua := true
		somma := 0
		for continua {

			// netData è un buffer (Reader bufferizzato) che viene riempito in base a ciò che conn riceverà (fino al carattere delimitatore '\n')
			netData, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			//Aspetto fine somma
			if strings.TrimSpace(string(netData)) == "0" {
				continua = false
				fmt.Fprintf(conn, "risultato: %d\n", somma)
			}

			//Estraggo valore ed errore
			addendo, _ := strconv.Atoi(netData)
			if err != nil {
				fmt.Println(err)
				return
			}

			//Aggiungo valore a somma
			somma += addendo

			// Controllo sulla richiesta di terminazione della connessione
			if strings.TrimSpace(string(netData)) == "esci" {
				continua = false // interrompi la connessione con il client
			}

			//Log del messaggio ricevuto
			fmt.Print(conn.RemoteAddr().String()+" -> ", string(netData))

			//altro modo possibile: conn.Write([]byte(stringaInvio))
		}

		//Disconnessione
		fmt.Println("Client " + conn.RemoteAddr().String() + " disconnesso.")
	}
}
