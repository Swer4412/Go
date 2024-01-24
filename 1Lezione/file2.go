package main

import "fmt"

func file2() {
	n := []int{}
	for {
		var input int
		fmt.Print("Inserisci un valore (0 per uscire): ")
		fmt.Scanf("%d\n", &input)
		if input == 0 {
			break
		}
		n = append(n, input)
	}

	for i, v := range n {
		if v > 0 {
			fmt.Println(i)
		}
	}

}