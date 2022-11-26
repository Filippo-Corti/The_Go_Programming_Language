package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const DIM = 4

func main() {
	rand.Seed(time.Now().UnixNano())
	parole := leggiParole()
	mischiaParole(&parole)
	griglia := [DIM][DIM]string{}
	griglia, _ = nuovaParola(griglia, parole, 0, "oriz") //Parto dalla prima orizzontale
	stampaGriglia(griglia)
}

func nuovaParola(griglia [DIM][DIM]string, parole []string, start int, verso string) ([DIM][DIM]string, bool) {
	var success bool
	c := 0
	contaFallimenti := 0
	if start == DIM {
		return griglia, true //Ho finito lo spazio da coprire
	}
	for {
		//cancella()
		//stampaGriglia(*griglia)
		//1. Inserisco la parola trovata
		for !verificaValiditaParola(griglia, parole[c], start, verso) {
			c++
			if c >= len(parole) - contaFallimenti {
				return griglia, false
			}
		}
		backupGriglia := griglia 
		backupParola := parole[c]
		inserisciInGriglia(&griglia, parole[c], start, verso) 
		parole = append(parole[:c], parole[c + 1:]...)
		//2. Lancio la ricorsione
		if verso == "oriz" {
			griglia, success = nuovaParola(griglia, parole, start, "vert")
		} else {
			griglia, success = nuovaParola(griglia, parole, start + 1, "oriz")
		}
		//3. Controllo che tutto ciò che sia avvenuto dopo di me sia andato bene, in tal caso ho finito anche io
		if success {
			return griglia, true
		} else {
			//4. Se non è andato tutto bene devo cancellare la parola e ricominciare
			griglia = backupGriglia
			parole = append(parole, backupParola)
			contaFallimenti++
			c = 0
			//fmt.Println(contaFallimenti)
			//cancellaParolaDallaGriglia(griglia, parole[c], start, verso)
		}
	}
}

func leggiParole() (parole []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parole = append(parole, line)
	}
	return
}

func mischiaParole(s *[]string) {
	for i := 0; i < len(*s); i++ {
		r := rand.Intn(len(*s) - i) + i
		(*s)[i], (*s)[ r] = (*s)[r], (*s)[i]
	}
}

func stampaGriglia(g [DIM][DIM]string) {
	for i := 0; i < DIM; i++ {
		for j := 0; j < DIM; j++ {
			fmt.Print(g[j][i])
		}
		fmt.Println()
		//fmt.Println()
	}
}


func inserisciInGriglia(griglia *[DIM][DIM]string, parola string, start int, verso string) {
	if verso == "vert" {
		for i := 0; i < len(parola); i++ {
			griglia[start][i] = string(parola[i])
		}
	} else if verso == "oriz" {
		for i := 0; i < len(parola); i++ {
			griglia[i][start] = string(parola[i])
		}
	}
}

func cancellaParolaDallaGriglia(griglia *[DIM][DIM]string, parola string, start int, verso string) {
	if verso == "vert" {
		for i := 0; i < len(parola); i++ {
			griglia[start][i] = ""
		}
	} else if verso == "oriz" {
		for i := 0; i < len(parola); i++ {
			griglia[i][start] = ""
		}
	}
}

func verificaValiditaParola(griglia [DIM][DIM]string, parola string, start int, verso string) bool {
	if verso == "vert" {
		for i := 0; i < len(parola); i++ {
			if griglia[start][i] != string(parola[i]) && griglia[start][i] != "" {
				return false
			}
		}
	} else if verso == "oriz" {
		for i := 0; i < len(parola); i++ {
			if griglia[i][start] != string(parola[i]) && griglia[i][start] != "" {
				return false
			}
		}
	}
	return true
}

func cancella() {
	fmt.Println("%cc", 27)
}
 
func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}
 