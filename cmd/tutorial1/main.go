package main //Serve per indicare al compilatore che dentro questo package (cartella) c'è una funzione main da cui partire
import ("fmt"
"time"
"sync"
)
var m = sync.Mutex{}
var wg= sync.WaitGroup{} //Nuovo waitgroup
var results = []string{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1) //aumento di uno il numero di task
		go dbCall(i)
	}
	wg.Wait() //Aspetta che tutte le task eseguano Done()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000 //Tolgo randomizzazione
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The reuslt from the database:", dbData[i])
	m.Lock() //Guarda se la risorsa è già stata presa da un altra task
	results = append(results, dbData[i])
	m.Unlock() //Una volta acceduto alla risorsa, la si sblocca 
	wg.Done()
}