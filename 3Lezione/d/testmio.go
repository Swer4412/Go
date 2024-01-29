package main

import (
	"fmt"
	"log"
	"net/http"
)

// func handlerHTML scrive del codice HTML. Il codice HTML è sempre un testo.
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/1" {
		fmt.Fprint(w, "<p>Pagina 1</p><a href='/2'>2</a>")
		return
	} else if r.URL.Path == "/2" {
		fmt.Fprint(w, "<p>Pagina 2</p><a href='/1'>1</a>")
		return
	} else if r.URL.Path == "/" {
		fmt.Fprint(w, "<a href='/1'>1</a>&nbsp<a href='/2'>2</a>")
	}
}

func main() {
	http.HandleFunc("/", handlerHTML)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
