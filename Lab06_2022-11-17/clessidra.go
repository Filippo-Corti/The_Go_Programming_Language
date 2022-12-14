package main

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"time"
)

type Clessidra struct {
	altezza int
	stato int
}

func disegnaClessidra(clessidra Clessidra) {
	disegnaClessidraSopra(clessidra.altezza / 2, (clessidra.altezza / 2) - clessidra.stato)
	disegnaClessidraSotto(clessidra.altezza / 2, clessidra.stato)
}

func disegnaClessidraSopra(altezza, livelloPienezza int) {
	larghezza := altezza * 2 + 1
	stampaCaratteri("━", larghezza, true)
	for i := 0; i < altezza; i++ {
		if livelloPienezza <= altezza - i - 1 {
			disegnaClessidraRiga(i, "\\", "/", larghezza - ((i + 1) * 2), " ")
		} else {
			disegnaClessidraRiga(i, "\\", "/", larghezza - ((i + 1) * 2), "*")
		}
	}
}

func disegnaClessidraSotto(altezza, livelloPienezza int) {
	larghezza := altezza * 2 + 1
	for i := altezza - 1; i >= 0; i-- {
		if livelloPienezza > i {
			disegnaClessidraRiga(i, "/", "\\", larghezza - ((i + 1) * 2), "*")
		} else {
			disegnaClessidraRiga(i, "/", "\\", larghezza - ((i + 1) * 2), " ")
		}
	}
	stampaCaratteri("━", larghezza, true)
}

// contenuto corrisponde alla stringa contenuta fra i due bordi
func disegnaClessidraRiga(spaziIniziali int, bordoSx string, bordoDx string, lunghezzaContenuto int, contenuto string, ) {
	stampaCaratteri(" ", spaziIniziali, false)
	stampaCaratteri(bordoSx, 1, false)
	stampaCaratteri(contenuto, lunghezzaContenuto, false)
	stampaCaratteri(bordoDx, 1, true)
}

func stampaCaratteri(s string, n int, a_capo bool) {
	fmt.Print(strings.Repeat(s, n))
	//fmt.Print(strings.Repeat(s + " ", n)) //Per Clessidre Larghe
	if a_capo {
		fmt.Println()
	}
}

func main() {
	var clessidra Clessidra
	fmt.Print("Quanto tempo vuoi contare? ")
	fmt.Scan(&clessidra.altezza)
	clessidra.altezza *= 2
	contoAllaRovescia(&clessidra)
}

func contoAllaRovescia(clessidra *Clessidra) {
	for i := 0; i <= clessidra.altezza / 2; i++ {
		cancella()
		disegnaClessidra(*clessidra)
		attendi(1)
		clessidra.stato++
	}
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	cmd.Run()
}
 
func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}
 