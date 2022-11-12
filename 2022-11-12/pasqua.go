/* 
! DISCLAIMER
Il programma va solo per il calcolo della Pasqua. Le fasi lunari sono un po' più complesse :)
*/

package main

import "fmt"

type date struct {
	day, month, year int
}

func printDate(d date) string {
	return fmt.Sprintf("%d-%d-%d", d.day, d.month, d.year)
}

func main() {
	var anno int
	fmt.Printf("Inserire anno: ")
	fmt.Scan(&anno)
	fmt.Printf("Pasqua cade il %s\n", printDate(getPasqua(anno)))
}

func getPasqua(year int) date {
	var d date
	d = date{day: 21, month: 3, year: year}
	//Cerca Luna Piena
	for {
		giornoCicloLunare := getGiornoCicloLunare(d)
		if giornoCicloLunare == 15 {
			break
		}
		aumentaGiornoCirca(&d)
	}
	//Cerca Domenica
	for {
		aumentaGiornoCirca(&d)
		if dayOfTheWeek(d) == "Sunday" {
			break
		}
	}
	return d
}


func getGiornoCicloLunare(date date) int {
	return (date.day + getCapomese(date.month) + getPatta(date.year)) % 30
}

func getCapomese(m int) int {
	return (m - 1 + 10) % 12 + 1
}

//Patta nell'anno 1970: 22
//Primo anno che considero che ha fatto il +12: 1957
func getPatta(y int) int {
	patta := 22
	for i := 1970; i < y; i++{
		patta += 11
		if (i - 1957) % 19 == 0 {
			patta += 1
		}
	}
	return patta % 30
}

//Funziona solo per i giorni che ci fanno comodo
func aumentaGiornoCirca(date *date) {
	date.day += 1
	if date.day > 31 {
		date.day = 1
		date.month += 1
	}
}

// Monday: 0, Tuesday: 1, Wednesday: 2, Thursday: 3, Friday: 4, Saturday: 5, Sunday: 6
func dayOfTheWeek(date date) string {
	epochDay := 3 // 1-1-1970 era un giovedì (3)
	dayOfTheWeek := (daysFromEpoch(date.day, date.month, date.year) - 1 + epochDay) % 7
	switch dayOfTheWeek {
	case 0:
		return "Monday"
	case 1:
		return "Tuesday"
	case 2:
		return "Wednesday"
	case 3:
		return "Thursday"
	case 4:
		return "Friday"
	case 5:
		return "Saturday"
	case 6:
		return "Sunday"
	}
	return "ERRORE (Non so quale)"
}

func isAnnoBisestile(anno int) bool {
	return (anno%100 == 0 && anno%400 == 0) || (anno%100 != 0 && anno%4 == 0)
}

func giorniInMese(mese, anno int) int {
	if mese == 2 && isAnnoBisestile(anno) {
		return 29
	}
	if mese == 2 {
		return 28
	}
	if mese == 4 || mese == 6 || mese == 9 || mese == 11 {
		return 30
	}
	return 31
}

func daysFromEpoch(g, m, a int) int {
	var totale int
	//Conto Anni
	for i := 1970; i < a; i++ {
		if isAnnoBisestile(i) {
			totale += 366
		} else {
			totale += 365
		}
	}
	//Conto Mesi
	for i := 1; i < m; i++ {
		totale += giorniInMese(i, a)
	}
	//Conto Giorni
	totale += g
	//Restituisco il totale
	return totale
}
