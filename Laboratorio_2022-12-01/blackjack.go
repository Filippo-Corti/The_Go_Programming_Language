package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const ASCII_BLACKJACK = `
 ____  _            _        _            _    
|  _ \| |          | |      | |          | |   
| |_) | | __ _  ___| | __   | | __ _  ___| | __
|  _ <| |/ _\ |/ __| |/ /   | |/ _\ |/ __| |/ /
| |_) | | (_| | (__|   < |__| | (_| | (__|   < 
|____/|_|\__,_|\___|_|\_\____/ \__,_|\___|_|\_\`

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

func main() {
	rand.Seed(time.Now().UnixNano())
	ris := playBJ()
	switch ris {
	case -1:
		fmt.Println("Hai Perso! :(")
	case 1:
		fmt.Println("Hai Vinto! :)")
	case 0:
		fmt.Println("Hai Pareggiato! :|")
	}
}

//Gestisce una partita di BlackJack. Restituisce:
// -1 se il giocatore perde
// 1 se il giocatore vince
// 0 se il giocatore pareggia
func playBJ() int {
	var pesca string
	mazzo := append(mazzoPocker(), mazzoPocker()...)
	mischia(&mazzo)
	manoGiocatore := []Carta{preleva(&mazzo), preleva(&mazzo)}
	manoTavolo := []Carta{preleva(&mazzo)}
	stampaTavolo(manoTavolo, manoGiocatore)
	//Se il giocatore ha già 21 vince
	if contaPunti(manoGiocatore) == 21 {
		return 1
	}
	//Il giocatore può pescare fino a quando non supera 21
	for {
		fmt.Print("Vuoi pescare un'altra carta? (y/n) ")
		fmt.Scan(&pesca)
		if pesca != "y" {
			break
		}
		attendi(1)
		manoGiocatore = append(manoGiocatore, preleva(&mazzo))
		stampaTavolo(manoTavolo, manoGiocatore)
		if contaPunti(manoGiocatore) > 21 {
			return -1
		}
	}
	//Il tavolo pesca 1 carta. Se il punteggio è meno di 17 ne pesca un'altra
	attendi(1)
	manoTavolo = append(manoTavolo, preleva(&mazzo))
	stampaTavolo(manoTavolo, manoGiocatore)
	if contaPunti(manoTavolo) < 17 {
		attendi(1)
		manoTavolo = append(manoTavolo, preleva(&mazzo))
		stampaTavolo(manoTavolo, manoGiocatore)
	}
	if contaPunti(manoTavolo) > 21 {
		return 1
	}
	stampaTavolo(manoTavolo, manoGiocatore)
	attendi(2)
	//Confronto i punteggi
	return valutaVittoria(manoTavolo, manoGiocatore)
} 

func valutaVittoria(manoTavolo, manoGiocatore []Carta) int {
	if contaPunti(manoGiocatore) > contaPunti(manoTavolo) {
		return 1
	} else if contaPunti(manoGiocatore) < contaPunti(manoTavolo) {
		return -1
	} else {
		return 0
	}
}

func stampaTavolo(manoTavolo, manoGiocatore []Carta) {
	cancella()
	fmt.Println(ASCII_BLACKJACK)
	fmt.Println("Tavolo:")
	fmt.Println(manoTavolo)
	fmt.Println("La tua Mano:")
	fmt.Println(manoGiocatore)
	fmt.Println("-------------------")
}

//Riceve una slice di Carta e valuta e restituisce la somma dei punteggi delle singole carte.
//Se il punteggio supera 21 ed è presente un asso nella mano, l'asso viene fatto valere 1
func contaPunti(mano []Carta) (p int) {
	for _, c := range mano {
		p += getValoreBJ(c)
	}
	if p > 21 && contieneAsso(mano) {
		p -= 10
	}
	return
}

func contieneAsso(mano []Carta) bool {
	for _, c := range mano {
		if c.valore == "A" {
			return true
		}
	}
	return false
}

//Riceve una Carta e ritorna il valore numerico corrispondente al suo valore (A -> 11)
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

//Crea una slice contenente le 52 carte del mazzo e la restituisce
func mazzoPocker() (mazzo []Carta) {
	for i := 0; i < 52; i++ {
		carta, _ := carta(i)
		mazzo = append(mazzo, carta)
	}
	return 
}

//Mischia una slice di Carta
func mischia(mazzo *[]Carta) {
	for i := 0; i < len(*mazzo); i++ {
		r := rand.Intn(len(*mazzo) - i) + i
		(*mazzo)[i], (*mazzo)[r] = (*mazzo)[r], (*mazzo)[i]
	}
}

func preleva(mazzo *[]Carta) Carta {
	carta := (*mazzo)[0]
	(*mazzo) = (*mazzo)[1:]
	return carta
}


func attendi(n_seconds float64){
	time.Sleep(time.Duration(n_seconds) * time.Second)
}

func cancella() {
	cmd := exec.Command("clear")
	cmd.Stdout=os.Stdout
	cmd.Run()
}