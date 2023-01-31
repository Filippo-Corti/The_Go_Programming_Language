package main

import "fmt"
import "strings"
import "strconv"

type date struct {
	day, month, year int
}

func main() {
	var date string
	fmt.Printf("Insert date (gg-mm-aaaa): ")
	fmt.Scan(&date)
	fmt.Printf("The day %s was a %s\n", printDate(splitIntoDate(date)), dayOfTheWeek(splitIntoDate(date)))
}

func printDate(d date) string {
	return fmt.Sprintf("%d-%d-%d", d.day, d.month, d.year)
}

func splitIntoDate(d string) date {
	splitted := strings.Split(d, "-")
	day, err1 := strconv.Atoi(splitted[0])
	month, err2 := strconv.Atoi(splitted[1])
	year, err3 := strconv.Atoi(splitted[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return date{-1, -1, -1}
	}
	return date{day, month, year}
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
