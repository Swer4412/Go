package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import ("fmt"
"strings")

func main() {
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune) //myRune = 97
	
	var strSlice = []string{"s","a","l","v","e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%v", strBuilder.String())
}