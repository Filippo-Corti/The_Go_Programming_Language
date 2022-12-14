package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

/* VARIABILI GLOBALI */
var CUORI rune = '\u2661'   //Cuori ♡
var QUADRI rune = '\u2662' //Quadri ♢
var FIORI rune = '\u2667'    //Fiori ♧
var PICCHE rune = '\u2664'   //Picche ♤

const SEMI = 4
const VALORI = 13
const CARTE = 52

var VALORI_CARTE [VALORI]string = [VALORI]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var SEMI_CARTE [SEMI]string = [SEMI]string{string(CUORI), string(QUADRI), string(FIORI), string(PICCHE)}

/* CARTA */
type Carta struct {
	valore, seme string
}

func carta(n int) (Carta, bool) {
	if n < 0 || n >= CARTE {
		return Carta{"", ""}, false 
	}
	return Carta{valore: getValoreCarta(n), seme: getSemeCarta(n)}, true
}

func getValoreCarta(n int) string {
	return VALORI_CARTE[n % 13]
}

func getSemeCarta(n int) string {
	return SEMI_CARTE[n / 13]
}

/*************************************/

func main()  {
	rand.Seed(time.Now().UnixNano())
	mazzo := mazzoPocker()
	mischia(&mazzo)	
	fmt.Println("Il Mazzo è pronto e mischiato")
	playBJ(&mazzo)
}

func playBJ(mazzo *[CARTE]Carta) {
	var scelta int
	manoTavolo := mazzo[0:2]
	manoGiocatore := mazzo[2:4]
	i := 4
	fmt.Println("Carte Tavolo: ", manoTavolo, "{ Punti:", contaPunti(manoTavolo), "}")
	fmt.Println("Le tue Carte: ", manoGiocatore, "{ Punti:", contaPunti(manoGiocatore), "}")
	fmt.Print("Scegli la tua prossima mossa:\n 1 -> Mantieni \n 2-> Nuova Carta\n")
	fmt.Scan(&scelta)
	attendi(1)
	if scelta == 2 {
		manoGiocatore = append(manoGiocatore, mazzo[i])
		i++
	} 
	if contaPunti(manoTavolo) < 17 {
		manoTavolo = append(manoTavolo, mazzo[i])
		i++
	}
	fmt.Println("Le tue Carte: ", manoGiocatore, "{ Punti:", contaPunti(manoGiocatore), "}")
	attendi(2)
	fmt.Println("Carte Tavolo: ", manoTavolo, "{ Punti:", contaPunti(manoTavolo), "}")
	if valutaVittoria(manoTavolo, manoGiocatore) {
		fmt.Println("HAI VINTO!")
	} else {
		fmt.Println("HAI PERSO!")
	}
}

func contaPunti(mano []Carta) (p int) {
	for _, c := range mano {
		p += getValoreBJ(c)
	}
	return
}

func valutaVittoria(tavolo, giocatore []Carta) bool {
	if contaPunti(giocatore) > 21 {
		return false
	}
	if contaPunti(tavolo) > 21 {
		return true
	}
	return contaPunti(tavolo) <= contaPunti(giocatore) 
}

func getValoreBJ(carta Carta) int {
	switch carta.valore {
	case "A":
		return 11 
	case "J", "Q", "K":
		return 10
	default:
		val, _ := strconv.Atoi(carta.valore)
		return val
	}
}

func mazzoPocker() (mazzo [CARTE]Carta) {
	for i, _ := range mazzo {
		mazzo[i], _ = carta(i)
	}
	return 
}

func mischia(mazzo *[CARTE]Carta) {
	for i := 0; i < CARTE; i++ {
		r := rand.Intn(CARTE - i) + i
		mazzo[i], mazzo[r] = mazzo[r], mazzo[i]
	}
}

func estraiCarta() Carta {
	r := rand.Intn(CARTE)
	c, _ := carta(r)
	return c
}

func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}
 