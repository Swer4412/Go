package main

import (
    "fmt"
    "log"
    "net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ciao %s!\n", r.URL.Path[1:]) //Fprintf invece che scirvere alla console,
    // scrive su w che nel nostro caso è il scrittore di risposte http 
}

func main() { 
    http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
