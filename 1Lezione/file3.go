package main

import "fmt"

func file3() {
	s := []int{}
	for {
		var input int
		fmt.Print("Inserisci un valore (0 per uscire): ")
		fmt.Scanf("%d\n", &input)
		if input == 0 {
			break
		}
		s = append(s, input)
	}

	n := []int{}
	for i := len(s)-1; i >= 0; i-- {
        n = append(n, s[i])
    }
	fmt.Print(n)

}