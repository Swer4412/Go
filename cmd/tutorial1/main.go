package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import ("fmt"
"unicode/utf8")

func main() {
	foo()
}

func foo(numero int) (int, int) {
	return numero*2, numero/2
}