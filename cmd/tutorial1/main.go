package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import ("fmt"
)

func main() {
	//[string] è il tipo della chiave, uint8 è il tipo del valore

	
	
	test := []string{"a", "b", "c"}

	for indice, elemento:= range test {
		fmt.Println(indice, elemento)
	}
}