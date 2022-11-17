package main

import "fmt"
import _ "strconv"
import "math/rand"
import "time"

type Carta struct {
	valore, seme string
}

var CUORI rune = '\u2661'   //Cuori ♡
var QUADRI rune = '\u2662' //Quadri ♢
var FIORI rune = '\u2667'    //Fiori ♧
var PICCHE rune = '\u2664'   //Picche ♤

var valoriCarte [VALORI]string = [VALORI]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var semiCarte [SEMI]string = [SEMI]string{string(CUORI), string(QUADRI), string(FIORI), string(PICCHE)}

const SEMI = 4
const VALORI = 13
const CARTE = 52

func main() {
	rand.Seed(time.Now().UnixNano())
	c1 := estraiCarta()
	fmt.Println(c1)
	quattroCarte := dai4Carte()
	fmt.Println(quattroCarte)
}

func dai4Carte() (carte [4]Carta) {
	for i := 0; i < 4; i++ {
		estratta := estraiCarta()
		if containsCarta4(carte, estratta) {
			i--
		} else {
			carte[i] = estratta
		}
	}
	return
}

func estraiCarta() Carta {
	r := rand.Intn(CARTE)
	c, _ := carta(r)
	return c
}


func carta(n int) (Carta, bool) {
	if n < 1 || n > CARTE {
		return Carta{"", ""}, false 
	}
	return Carta{valore: getValoreCarta(n), seme: getSemeCarta(n)}, true
}

func getValoreCarta(n int) string {
	return valoriCarte[n % 13]
}

func getSemeCarta(n int) string {
	return semiCarte[n / 13]
}

func containsCarta4(carte [4]Carta, c Carta) bool {
	for _, x := range carte {
		if x == c {
			return true
		}
	}
	return false
}