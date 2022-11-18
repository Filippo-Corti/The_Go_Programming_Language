package main

import "fmt"

type Data struct {
	g, m, a int
}

func main() {
	sliceDiAnni := []int{1986, 2003, 2022, 2023, 2035, 2053}
	sliceDiNatali := getNatali(sliceDiAnni)
	fmt.Println("Natali degli anni della slice: ", sliceDiNatali)
	fmt.Println("Giorni della settimana corrispondenti: ", giorniDellaSettimana(sliceDiNatali))
}

func getNatali(s []int) (natali []Data) {
	for _, anno := range s {
		natali = append(natali, Data{g: 25, m: 12, a: anno})
	}
	return
}

func giorniDellaSettimana(date []Data) (giorni []string) {
	for _, data := range date {
		giorni = append(giorni, dayOfTheWeek(data))
	}
	return
}

/* Calcolo Giorno della Settimana */

func dayOfTheWeek(date Data) string {
	giorniDellaSettimana := [...]string{"Lunedì", "Martedì", "Mercoledì", "Giovedì", "Venerdì", "Sabato", "Domenica"}
	epochDay := 3 // 1-1-1970 era un giovedì (3)
	dayOfTheWeek := (daysFromEpoch(date.g, date.m, date.a) - 1 + epochDay) % 7
	return giorniDellaSettimana[dayOfTheWeek]
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
