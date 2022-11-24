package main

import (
	"fmt"
	"sort"
	"os"
)

func main() {
	var slice1, slice2 []string
	var parola string
	//Crea Slice1
	if len(os.Args) < 4 {
		fmt.Println("Troppi pochi argomenti")
		os.Exit(1)
	}
	slice1 = append(slice1, os.Args[1:]...)
	fmt.Println("Slice 1: ")
	fmt.Println(slice1)
	//Crea Slice2
	for {
		fmt.Scan(&parola)
		if parola == "stop" {
			break
		}
		slice2 = append(slice2, parola)
	}
	fmt.Println("Slice 2: ")
	fmt.Println(slice2)
	//Unisci Slice
	sliceUnita := append(slice1, slice2...)
	fmt.Println("1. Slice Unita:", sliceUnita)
	//Ordina Slice
	sort.Strings(sliceUnita)
	fmt.Println("2. Slice Ordinata:", sliceUnita)
	//Elimina ultimo elemento
	sliceUnita = sliceUnita[:len(sliceUnita) - 1]
	fmt.Println("3. Slice Senza Ultimo:", sliceUnita)
	//Rimuovo elementi dal 2 al 4
	sliceUnita = append(sliceUnita[:2], sliceUnita[4 + 1:]...)
	fmt.Println("4. Slice Senza Elementi dal 2 al 4:", sliceUnita)
	//Creo nuova slice {a,b,c}
	nuovaSlice := []string{"a", "b", "c"}
	fmt.Println("5. NuovaSlice:", nuovaSlice)
	//Inserisco nuovaSlice in posizione 1 di vecchia lista
	sliceUnita = append(sliceUnita[:1], append(nuovaSlice, sliceUnita[1:]...)...)
	fmt.Println("6. Unisco la Slice a Slice Precedente:", sliceUnita)
	//Leggo nuova parola e la inserisco in posizione 2
	var p string
	fmt.Print("Una parola: ")
	fmt.Scan(&p)
	sliceUnita = append(sliceUnita[:2], append([]string{p}, sliceUnita[2:]...)...)
	fmt.Println("7. Slice Con Aggiunta di Parola In Pos 2:", sliceUnita)
	//Leggo nuova parola e la inserisco alla fine
	var p2 string
	fmt.Print("Una parola: ")
	fmt.Scan(&p2)
	sliceUnita = append(sliceUnita, p2)
	fmt.Println("8. Slice Con Aggiunta di Parola In Fondo:", sliceUnita)
	//Estendo slice di 2 elementi vuoti
	sliceUnita = append(sliceUnita, make([]string, 2)...)
	fmt.Println("9. Slice Con Aggiunta di 2 in fondo", sliceUnita)
	//Estendo in posizione 3 con una slice di 3 elementi vuoti
	sliceUnita = append(sliceUnita[:3], append(make([]string, 3), sliceUnita[3:]...)...)
	fmt.Println("10. Slice Con Aggiunta di 3 in Posizione 3", sliceUnita)
	//Copio in una nuova slice quello che ho fatto finora
	var altraSlice []string
	altraSlice = append(altraSlice, sliceUnita[:]...)
	fmt.Println("11. Copia della Slice ", altraSlice)
	//Rimuovo ultimo elemento (solo dalla copia)
	altraSlice = altraSlice[:len(altraSlice) - 1]
	fmt.Println("12. Le due slice ora sono: ", altraSlice, sliceUnita)
}