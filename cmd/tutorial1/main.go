package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import ("fmt"
)

func main() {
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nMemory location of thing1 is: %p", &thing1)
	var result [5]float64 = square(&thing1) //Passo indirizzo di memoria
	fmt.Printf("\nThe result is : %v", result)
	fmt.Printf("\nThe value of thing1 is : %v", thing1)
}

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nMemory location of thing2 is: %p", thing2)
	for i := range *thing2 {
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}