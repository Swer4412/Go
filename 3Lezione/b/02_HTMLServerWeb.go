package main

import (
    "fmt"
    "log"
    "net/http"
)

// func handlerHTML scrive del codice HTML. Il codice HTML è sempre un testo.
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,`<html>
	<body>
		Stringa con HTML! Ti piace <b><i><u><span style="color:red;">%s</span></u></i></b>?
	</body>
</html>
`, r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handlerHTML)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
