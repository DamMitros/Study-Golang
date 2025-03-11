package main

import (
	"fmt"
	"time"
	"math/rand/v2"
	// "golang.org/x/exp/rand"
)

// GeneratePESEL: geneuje numer PESEL
// Parametry:
// - birthDate: time.Time: reprezentacja daty urodzenia
// - płeć: znak "M" lub "K"
// Wyjscie:
//Tablica z cyframi numeru PESEL


func GenerujPESEL(birthDate time.Time, gender string) [11]int {

	// tablica zawierajaca kolejne cyfry numeru PESEL 
	var cyfryPESEL [11]int 

	// konwersja daty na dane skladowe 
	year := birthDate.Year()
	month := int(birthDate.Month())
	day := birthDate.Day()

	// losowy numer
	randomSerial := rand.IntN(900) + 100 // 3 cyfrowy losowy numer z zakresu 100-999

	rr:= year % 100

	switch{
		case year >= 1800 && year < 1900:
			month += 80
		case year >= 2000 && year < 2100:
			month += 20
		case year >= 2100 && year < 2200:
			month += 40
		case year >= 2200 && year < 2300:
			month += 60
	}

	cyfryPESEL[0] = rr / 10
	cyfryPESEL[1] = rr % 10
	cyfryPESEL[2] = month / 10
	cyfryPESEL[3] = month % 10
	cyfryPESEL[4] = day / 10
	cyfryPESEL[5] = day % 10
	cyfryPESEL[6] = randomSerial / 100
	cyfryPESEL[7] = (randomSerial % 100) / 10
	cyfryPESEL[8] = randomSerial % 10

	switch gender{
		case "M":
			cyfryPESEL[9] = rand.IntN(5) * 2 + 1
		case "K":
			cyfryPESEL[9] = rand.IntN(5) * 2 
	}

	var wagi=[10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	var suma int = 0

	// for i:=range 10{
	for i:=0; i<10; i++{
		suma += (wagi[i] * cyfryPESEL[i])%10
	}
	cyfryPESEL[10]=10-(suma%10)
	return cyfryPESEL
}

// WeryfikujPESEL: weryfikuje poprawność numeru PESEL
// Parametry:
// - cyfryPESEL: Tablica z cyframi numeru PESEL
// Wyjscie:
//zmienna bool

func WeryfikujPESEL(cyfryPESEL [11]int) bool {

	var czyPESEL bool
	var wagi=[10]int{1, 3, 7, 9, 1, 3, 7, 9, 1, 3}
	var suma int = 0

	for i:=0; i<10; i++{
		suma += (wagi[i] * cyfryPESEL[i])%10
	}

	czyPESEL = (10-(suma%10) == cyfryPESEL[10])

	return czyPESEL
}

// Przykład użycia
func main() {
	// 
	birthDate := time.Date(1980, 2, 26, 0, 0, 0, 0, time.UTC)
	// birthDate := time.Date(2005,1,14,0,0,0,0,time.UTC)
	pesel := GenerujPESEL(birthDate, "M")

	fmt.Println("Wygenerowany PESEL:", pesel)

	fmt.Println("Czy numer PESEL jest poprawny:", WeryfikujPESEL(pesel))
}