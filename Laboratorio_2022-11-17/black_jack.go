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
	// c, _ := carta(25)
	// fmt.Println(c, getValoreBJ(c))
	mazzo := mazzoPocker()
	fmt.Println("Mazzo:")
	fmt.Println(mazzo)
	mischia(&mazzo)	
	fmt.Println("Mazzo mischiato:")
	fmt.Println(mazzo)
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
