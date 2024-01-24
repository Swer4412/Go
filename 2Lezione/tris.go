package main

import (
	"fmt"
)

func mostraCampo(posizioni [9]string) {
	for i := 0; i < len(posizioni); i++ {
		fmt.Printf(" %s ", posizioni[i])
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}

func haVinto(posizioni [9]string) bool {
	// Controlla le righe
	if posizioni[0] == posizioni[1] && posizioni[1] == posizioni[2] {
		return true
	} else if posizioni[3] == posizioni[4] && posizioni[4] == posizioni[5] {
		return true
	} else if posizioni[6] == posizioni[7] && posizioni[7] == posizioni[8] {
		return true
	}
	
	// Controlla le colonne
	if posizioni[0] == posizioni[3] && posizioni[3] == posizioni[6] {
		return true
	} else if posizioni[1] == posizioni[4] && posizioni[4] == posizioni[7] {
		return true
	} else if posizioni[2] == posizioni[5] && posizioni[5] == posizioni[8] {
		return true
	}

	// Controlla le diagonali
	if posizioni[0] == posizioni[4] && posizioni[4] == posizioni[8] {
		return true
	} else if posizioni[2] == posizioni[4] && posizioni[4] == posizioni[6] {
		return true
	}

	// Nessun caso di vittoria
	return false
}


func main() {

	var posizioni = [9] string {"1","2","3","4","5","6","7","8","9"}
	var giocatori = [2] string {"X","O"}

	//Turni
	var turno int = 0
	for {
		//Mostro il campo
		mostraCampo(posizioni)

		//Se sono stati fatti più di 8 turni, allora tutti i campi sono pieni
		if turno > 8 {
			break
		}

		//Scrivo il turno
		fmt.Println("TURNO", turno)

		//Sceglo il giocatore in base al turno
		var giocatore string = giocatori[turno % 2]
		//Scrivo il giocatore
		fmt.Println("TOCCA AD", giocatore)

		//Prendo l'input
		var input int
		fmt.Print("Scegli una posizione: ")
		fmt.Scan(&input)
		if input > 9 || input < 1 {
			fmt.Println("Posizione", input, "inesistente")
			continue
		}

		//Assegno alla posizione giusta la X o la O
		if posizioni[input-1] != giocatori[0] && posizioni[input-1] != giocatori[1] {
			posizioni[input-1] = giocatore
		} else {
			fmt.Println("Posizione", input, "già usata, scegliene un'altra!")
			continue
		}

		if haVinto(posizioni) {
			fmt.Println("\n",giocatore, "HA VINTO")
			break
		}

		//Aumento il turno
		turno++
	}

	mostraCampo(posizioni)

	fmt.Println("Fine")


}
