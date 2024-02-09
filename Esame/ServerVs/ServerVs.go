// Importo i pacchetti necessari
package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"strconv"
	"os"
)


const (
	HOST = "127.0.0.1"
	PORT = "8080"

	ROWS = 6
	COLS = 7
)

const (
	PLAYER1 = 'X'
	PLAYER2 = 'O'
	EMPTY   = ' '
)

type Grid struct {
	cells [ROWS][COLS]rune // Una matrice di rune per le celle
}

func NewGrid() *Grid {
	// Creo una nuova griglia
	grid := &Grid{}
	// Per ogni riga e colonna
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			// Imposto la cella corrispondente a vuota
			grid.cells[i][j] = EMPTY
		}
	}
	// Restituisco la griglia
	return grid
}


func (grid *Grid) String() string {
	// Creo una variabile di tipo stringa
	var s string
	// Per ogni riga
	for i := 0; i < ROWS; i++ {
		// Aggiungo una linea orizzontale
		s += "+---+---+---+---+---+---+---+\n"
		// Aggiungo le celle della riga
		s += fmt.Sprintf("| %c | %c | %c | %c | %c | %c | %c |\n", grid.cells[i][0], grid.cells[i][1], grid.cells[i][2], grid.cells[i][3], grid.cells[i][4], grid.cells[i][5], grid.cells[i][6])
	}
	// Aggiungo l'ultima linea orizzontale
	s += "+---+---+---+---+---+---+---+\n"
	// Aggiungo i numeri delle colonne
	s += "  1   2   3   4   5   6   7  \n"
	// Restituisco la stringa
	return s
}


func (grid *Grid) Insert(col int, symbol rune) bool {
	// Se la colonna è fuori dal range
	if col < 0 || col >= COLS {
		// Restituisco falso
		return false
	}
	// Se la colonna è piena
	if grid.cells[0][col] != EMPTY {
		// Restituisco falso
		return false
	}
	// Partendo dall'ultima riga
	for i := ROWS - 1; i >= 0; i-- {
		// Se la cella è vuota
		if grid.cells[i][col] == EMPTY {
			// Inserisco il simbolo
			grid.cells[i][col] = symbol
			// Restituisco vero
			return true
		}
	}
	// Questo punto non dovrebbe essere mai raggiunto
	return false
}


func (grid *Grid) IsFull() bool {
	// Per ogni colonna
	for j := 0; j < COLS; j++ {
		// Se la cella in alto è vuota
		if grid.cells[0][j] == EMPTY {
			// Restituisco falso
			return false
		}
	}
	// Altrimenti restituisco vero
	return true
}


func (grid *Grid) HasWon(symbol rune) bool {
	// Per ogni riga
	for i := 0; i < ROWS; i++ {
		// Per ogni colonna
		for j := 0; j < COLS; j++ {
			// Se la cella corrisponde al simbolo
			if grid.cells[i][j] == symbol {
				// Controllo le quattro direzioni possibili
				// Orizzontale
				if j+3 < COLS && grid.cells[i][j+1] == symbol && grid.cells[i][j+2] == symbol && grid.cells[i][j+3] == symbol {
					// Restituisco vero
					return true
				}
				// Verticale
				if i+3 < ROWS && grid.cells[i+1][j] == symbol && grid.cells[i+2][j] == symbol && grid.cells[i+3][j] == symbol {
					// Restituisco vero
					return true
				}
				// Diagonale ascendente
				if i+3 < ROWS && j+3 < COLS && grid.cells[i+1][j+1] == symbol && grid.cells[i+2][j+2] == symbol && grid.cells[i+3][j+3] == symbol {
					// Restituisco vero
					return true
				}
				// Diagonale discendente
				if i-3 >= 0 && j+3 < COLS && grid.cells[i-1][j+1] == symbol && grid.cells[i-2][j+2] == symbol && grid.cells[i-3][j+3] == symbol {
					// Restituisco vero
					return true
				}
			}
		}
	}
	// Altrimenti restituisco falso
	return false
}

func RandomMove() int {
	// Inizializzo il generatore di numeri casuali
	rand.Seed(time.Now().UnixNano())
	// Genero un numero tra 0 e 6
	return rand.Intn(7)
}


func main() {
	// Creo un listener TCP sulla porta specificata
	listener, err := net.Listen("tcp", ":"+PORT)
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		os.Exit(1)
	}
	// Chiudo il listener alla fine del programma
	defer listener.Close()
	// Stampo un messaggio di attesa
	fmt.Println("In attesa di connessioni...")
	// Accetto una connessione in entrata
	conn, err := listener.Accept()
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		os.Exit(1)
	}
	// Chiudo la connessione alla fine del programma
	defer conn.Close()
	// Stampo un messaggio di benvenuto
	fmt.Println("Connessione stabilita con", conn.RemoteAddr())
	// Creo una nuova griglia
	grid := NewGrid()
	// Creo una variabile per il turno del giocatore
	turn := PLAYER1
	// Creo un ciclo infinito
	for {
		// Stampo la griglia a schermo
		fmt.Fprintln(conn, grid)
		// Se il turno è del giocatore 1
		if turn == PLAYER1 {
			// Invio un messaggio di richiesta di input al client
			fmt.Fprintln(conn, "Inserisci il numero della colonna (1-7) o 'q' per uscire:")
			// Creo una variabile per l'input
			var input string
			// Leggo l'input dal client
			fmt.Fscanln(conn, &input)
			// Se l'input è 'q'
			if input == "q" {
				// Invio un messaggio di uscita al client
				fmt.Fprintln(conn, "Hai abbandonato la partita.")
				// Esco dal ciclo
				break
			}
			// Converto l'input in un intero
			col, err := strconv.Atoi(input)
			// Se c'è un errore
			if err != nil {
				// Stampo un messaggio di errore
				// Invio un messaggio di errore al client
				fmt.Fprintln(conn, "Input non valido.")
				// Salto all'inizio del ciclo
				continue
			}
			// Sottraggo 1 per adattare l'input al range 0-6
			col--
			// Provo ad inserire il disco nella colonna scelta
			ok := grid.Insert(col, PLAYER1)
			// Se non è possibile
			if !ok {
				// Stampo un messaggio di errore
				fmt.Fprintln(conn, "Colonna piena o non esistente.")
				// Salto all'inizio del ciclo
				continue
			}
			// Controllo se il giocatore 1 ha vinto
			if grid.HasWon(PLAYER1) {
				// Invio un messaggio di vittoria al client
				fmt.Fprintln(conn, "Hai vinto!")
				// Esco dal ciclo
				break
			}
			// Controllo se la griglia è piena
			if grid.IsFull() {
				// Invio un messaggio di pareggio al client
				fmt.Fprintln(conn, "Pareggio.")
				// Esco dal ciclo
				break
			}
			// Cambio il turno
			turn = PLAYER2
		} else {
			// Se il turno è del giocatore 2
			// Genero una mossa casuale
			col := RandomMove()
			// Provo ad inserire il disco nella colonna scelta
			ok := grid.Insert(col, PLAYER2)
			// Se non è possibile
			if !ok {
				// Salto all'inizio del ciclo
				continue
			}
			// Controllo se il giocatore 2 ha vinto
			if grid.HasWon(PLAYER2) {
				// Invio un messaggio di sconfitta al client
				fmt.Fprintln(conn, "Hai perso!")
				// Esco dal ciclo
				break
			}
			// Controllo se la griglia è piena
			if grid.IsFull() {
				// Invio un messaggio di pareggio al client
				fmt.Fprintln(conn, "Pareggio.")
				// Esco dal ciclo
				break
			}
			// Cambio il turno
			turn = PLAYER1
		}
	}
	// Stampo un messaggio di fine
	fmt.Println("Il server termina.")
}
