package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	_ = testaSistema(0) //nuovaParola(griglia, 0, "oriz" )
}

func testaSistema(i int) bool {
	if i == 6 { //se i == DIM && verso == "vert" ho finito
		return true
	}
	for {
		fmt.Println(i) //cerca parola, se non trova entro len(parole) tentativi return false
		r := testaSistema(i + 1) //nuovaParola(griglia, 0, "oriz" ) ->  nuovaParola(griglia, 0, "vert" ) -> nuovaParola(griglia, 1, "oriz")
		if r { //se ritorna true ci fermiamo
			break //ritorno direttamente true qua ("se tutto quello che c'è dopo di me è andato bene, io vado bene")
		}
		//Altrimenti cancelliamo la parola inserita e si riparte a cercare
	}
	rand := rand.Intn(7)
	fmt.Println(i, rand > 1)
	return rand > 1 //Quando arrivo qua è perché è andato tutto bene (return true)
}