package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":8080"
	ip   = "127.0.0.1"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Errore 404 (not found).", http.StatusNotFound) // http.StatusNotFound è una costante per indicare l'errore 404
		return
	}

	// In base al metodo HTTP sceglie come operare
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Ricevuta una richiesta GET\n")
		m := r.URL.Query() //name:youssef
		for k, v := range m {
			fmt.Fprintf(w, "%s: %s\n", k, v)
			// attenzione a controllare la correttezza dei parametri ricevuti!!
		}
	default:
		http.Error(w, "Errore 405 (Method Not Allowed). È supportato solo il metodo GET.\n", http.StatusMethodNotAllowed) // http.StatusMethodNotAllowed è una costante per indicare l'errore 405
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Printf("Server HTTP GET in partenza...\n")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// Per effettuare una richiesta GET aprire da browser http://127.0.0.1:8080/?name=Marco&surname=Crobu
// oppure da terminale (con comando curl)
// curl -si "http://localhost:8080/?name=Marco&surname=Crobu"
// curl -vsi "http://localhost:8080/?name=Marco&surname=Crobu"
