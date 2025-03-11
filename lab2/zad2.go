package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sort"
)

type City struct {
	Name string
	Country string
	CountryCode string
	Timezone string
	Population int
}

func main(){
	cities, err := ReadCities("geonames-all-cities-with-a-population-1000.csv")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Loaded ", len(cities), " cities")
	
	howmanycities := 20
	
	fmt.Println("Top ", howmanycities, " cities by population")
	sortByPopulation(cities)
	printTopCities(cities, howmanycities)

	fmt.Println("Top ", howmanycities, " cities by name")
	sortByName(cities)
	printTopCities(cities, howmanycities)
}

func ReadCities(filename string) ([]City, error){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	var cities []City

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		population, _ := strconv.Atoi(record[14]) 

		city := City{
			Name: record[1],
			Country: record[7],
			CountryCode: record[6],
			Timezone: record[16],
			Population: population,
		}

		if city.Population > 0 {
			cities = append(cities, city)
		}
	}
	return cities,nil
}

func sortByPopulation(cities []City){
	sort.Slice(cities, func(i, j int) bool {
		return cities[i].Population > cities[j].Population
	}) 
}

func sortByName(cities []City){
	sort.Slice(cities, func(i, j int) bool {
		if cities[i].Name == cities[j].Name{
			return cities[i].Population < cities[j].Population
		}
		return cities[i].Name < cities[j].Name
	}) 
}

func printTopCities(cities []City, howmanycities int){
	for i:=range howmanycities{
		// fmt.Println(i+1, ":", cities[i])
		fmt.Println(i+1, ":", cities[i].Name, "(", cities[i].Country, "), population:", cities[i].Population)
	}
}