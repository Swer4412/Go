// Importo i pacchetti necessari
package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Definisco una costante per la dimensione della griglia di gioco
const GridSize = 7

// Definisco una struttura per rappresentare una cella della griglia
type Cell struct {
	row int    // la riga della cella
	col int    // la colonna della cella
	val string // il valore della cella, vuoto o "X" o "O"
}

// Definisco una struttura per rappresentare un giocatore
type Player struct {
	conn net.Conn // la connessione del giocatore
	name string   // il nome del giocatore
	sym  string   // il simbolo del giocatore, "X" o "O"
}

// Definisco una struttura per rappresentare il gioco
type Game struct {
	grid   [GridSize][GridSize]Cell // la griglia di gioco
	player [2]Player                // i due giocatori
	turn   int                      // il turno di gioco, 0 o 1
}

// Creo una funzione per inizializzare il gioco
func NewGame(p1, p2 net.Conn) *Game {
	// Creo un nuovo gioco
	g := &Game{}
	// Inizializzo la griglia con celle vuote
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			g.grid[i][j] = Cell{i, j, " "}
		}
	}
	// Inizializzo i giocatori con i nomi e i simboli
	g.player[0] = Player{p1, "Player 1", "X"}
	g.player[1] = Player{p2, "Player 2", "O"}
	// Inizializzo il turno a 0
	g.turn = 0
	// Restituisco il gioco
	return g
}

// Creo una funzione per stampare la griglia di gioco
func (g *Game) PrintGrid() {
	// Creo una variabile per la griglia in formato stringa
	var grid string
	// Aggiungo una riga di separazione
	grid += strings.Repeat("-", GridSize*4+1) + "\n"
	// Per ogni riga della griglia
	for i := 0; i < GridSize; i++ {
		// Aggiungo una barra verticale
		grid += "|"
		// Per ogni colonna della griglia
		for j := 0; j < GridSize; j++ {
			// Aggiungo il valore della cella e una barra verticale
			grid += fmt.Sprintf(" %s |", g.grid[i][j].val)
		}
		// Aggiungo un carattere di fine riga
		grid += "\n"
		// Aggiungo una riga di separazione
		grid += strings.Repeat("-", GridSize*4+1) + "\n"
	}
	// Invio la griglia al giocatore di turno
	fmt.Fprintln(g.player[g.turn].conn, grid)
}


// Creo una funzione per controllare se una cella è valida
func (g *Game) IsValidCell(row, col int) bool {
	// Controllo se la riga e la colonna sono entro i limiti della griglia
	if row < 0 || row >= GridSize || col < 0 || col >= GridSize {
		return false
	}
	// Controllo se la cella è vuota
	if g.grid[row][col].val != " " {
		return false
	}
	// Altrimenti la cella è valida
	return true
}

// Creo una funzione per controllare se una mossa è valida
func (g *Game) IsValidMove(col int) bool {
	// Controllo se la colonna è entro i limiti della griglia
	if col < 0 || col >= GridSize {
		return false
	}
	// Controllo se la cella più in alto della colonna è vuota
	if g.grid[0][col].val != " " {
		return false
	}
	// Altrimenti la mossa è valida
	return true
}

// Creo una funzione per eseguire una mossa
func (g *Game) MakeMove(col int) {
	// Ottengo il simbolo del giocatore di turno
	sym := g.player[g.turn].sym
	// Partendo dalla cella più in basso della colonna
	for i := GridSize - 1; i >= 0; i-- {
		// Se la cella è vuota
		if g.grid[i][col].val == " " {
			// Assegno il simbolo del giocatore alla cella
			g.grid[i][col].val = sym
			// Esco dal ciclo
			break
		}
	}
}

// Creo una funzione per controllare se il gioco è finito
func (g *Game) IsGameOver() bool {
	// Controllo se c'è una sequenza di quattro simboli uguali in orizzontale, verticale o diagonale
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			// Ottengo il valore della cella corrente
			val := g.grid[i][j].val
			// Se il valore è diverso da vuoto
			if val != " " {
				// Controllo se c'è una sequenza orizzontale
				if j+3 < GridSize && val == g.grid[i][j+1].val && val == g.grid[i][j+2].val && val == g.grid[i][j+3].val {
					return true
				}
				// Controllo se c'è una sequenza verticale
				if i+3 < GridSize && val == g.grid[i+1][j].val && val == g.grid[i+2][j].val && val == g.grid[i+3][j].val {
					return true
				}
				// Controllo se c'è una sequenza diagonale in alto a destra
				if i-3 >= 0 && j+3 < GridSize && val == g.grid[i-1][j+1].val && val == g.grid[i-2][j+2].val && val == g.grid[i-3][j+3].val {
					return true
				}
				// Controllo se c'è una sequenza diagonale in basso a destra
				if i+3 < GridSize && j+3 < GridSize && val == g.grid[i+1][j+1].val && val == g.grid[i+2][j+2].val && val == g.grid[i+3][j+3].val {
					return true
				}
			}
		}
	}
	// Se non c'è nessuna sequenza, il gioco non è finito
	return false
}

// Creo una funzione per controllare se il gioco è in pareggio
func (g *Game) IsDraw() bool {
	// Controllo se tutte le celle sono piene
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			// Se c'è una cella vuota, il gioco non è in pareggio
			if g.grid[i][j].val == " " {
				return false
			}
		}
	}
	// Se tutte le celle sono piene, il gioco è in pareggio
	return true
}


// Creo una funzione per gestire il gioco
func (g *Game) Play() {
	// Invio un messaggio di benvenuto ai giocatori
	g.SendMessage("Benvenuto al gioco di forza 4!", 0)
	g.SendMessage("Benvenuto al gioco di forza 4!", 1)
	// Fino a quando il gioco non è finito o in pareggio
	for !g.IsGameOver() && !g.IsDraw() {
		// Stampo la griglia di gioco
		g.PrintGrid()
		// Ottengo il giocatore di turno
		p := g.player[g.turn]
		// Invio un messaggio al giocatore di turno per chiedere la sua mossa
		g.SendMessage(fmt.Sprintf("%s, è il tuo turno. Inserisci il numero di colonna da 1 a %d:", p.name, GridSize), g.turn)
		// Leggo la risposta del giocatore
		input, err := g.ReadInput()
		// Se c'è un errore
		if err != nil {
			// Invio un messaggio di errore e termino il gioco
			g.SendMessage("Si è verificato un errore nella lettura dell'input. Il gioco termina.", 0)
			g.SendMessage("Si è verificato un errore nella lettura dell'input. Il gioco termina.", 1)
			return
		}
		// Converto l'input in un intero
		col, err := strconv.Atoi(input)
		// Se c'è un errore
		if err != nil {
			// Invio un messaggio di errore e continuo il ciclo
			g.SendMessage("Input non valido. Inserisci un numero intero.", g.turn)
			continue
		}
		// Sottraggo 1 alla colonna per adattarla all'indice della griglia
		col--
		// Controllo se la mossa è valida
		if g.IsValidMove(col) {
			// Eseguo la mossa
			g.MakeMove(col)
			// Cambio il turno
			g.turn = 1 - g.turn
		} else {
			// Invio un messaggio di errore e continuo il ciclo
			g.SendMessage("Mossa non valida. Inserisci una colonna libera.", g.turn)
			continue
		}
	}
	// Stampo la griglia finale
	g.PrintGrid()
	// Controllo se il gioco è finito
	if g.IsGameOver() {
		// Ottengo il vincitore
		winner := g.player[1-g.turn]
		looser := g.player[g.turn]
		// Invio un messaggio di congratulazioni al vincitore
		g.SendMessage(fmt.Sprintf("Complimenti, %s! Hai vinto il gioco!", winner.name), 1-g.turn)
		g.SendMessage(fmt.Sprintf("Peccato, %s! Hai perso il gioco!", looser.name), g.turn)
	} else {
		// Invio un messaggio di pareggio
		g.SendMessage("Il gioco è finito in pareggio. Nessuno ha vinto.", 0)
		g.SendMessage("Il gioco è finito in pareggio. Nessuno ha vinto.", 1)
	}
	// Invio un messaggio di saluto ai giocatori
	g.SendMessage("Il gioco è finito!", 0)
	g.SendMessage("Il gioco è finito!", 1)
}



// Creo una funzione per inviare un messaggio a un giocatore
func (g *Game) SendMessage(msg string, to int) {
	// Scrivo il messaggio sulla connessione del giocatore destinatario
	fmt.Fprintln(g.player[to].conn, msg)
}


// Creo una funzione per leggere l'input del giocatore di turno
func (g *Game) ReadInput() (string, error) {

	var input string
	// Leggo una riga dalla connessione del giocatore di turno
	_, err := fmt.Fscanln(g.player[g.turn].conn, &input)

	return input, err
}

// Creo una funzione principale
func main() {
	// Creo un listener su una porta TCP
	ln, err := net.Listen("tcp", ":8080")
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		return
	}
	// Chiudo il listener alla fine del programma
	defer ln.Close()
	// Stampo un messaggio di attesa
	fmt.Println("In attesa di due client...")
	// Accetto la prima connessione
	p1, err := ln.Accept()
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		return
	}
	// Stampo un messaggio di conferma
	fmt.Println("Primo client connesso.")
	// Accetto la seconda connessione
	p2, err := ln.Accept()
	// Se c'è un errore
	if err != nil {
		// Stampo l'errore e termino il programma
		fmt.Println(err)
		return
	}
	// Stampo un messaggio di conferma
	fmt.Println("Secondo client connesso.")
	// Creo un nuovo gioco con le due connessioni
	g := NewGame(p1, p2)
	// Avvio il gioco
	g.Play()
}
