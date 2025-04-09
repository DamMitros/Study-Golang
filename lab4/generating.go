package main

import "math/rand"

func GenerateRandomOrder(id int) Order {
	customerNames := []string{"Justyna Steczkowska", "Krzysztof Krawczyk", "Marek Grechuta", "Anna Jantar", "Doda", "Blanka"}
	items := []string{"Guitar", "Piano", "Drums", "Violin", "Flute", "Saxophone", "Trumpet", "Harmonica"}

	numItems := rand.Intn(3) + 1
	randomItems := make([]string, 0, numItems)

	for range numItems {
		randomItems = append(randomItems, items[rand.Intn(len(items))])
	}
	
	return Order{
		ID: id,
		CustomerName: customerNames[rand.Intn(len(customerNames))],
		Items: randomItems,
		TotalAmount: 100.0 + rand.Float64()*900.0,
	}
}