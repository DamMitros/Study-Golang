package main

import (
	"fmt"
	"os"

	"lab5/api"
	"lab5/ui"
)

func main() {
	fmt.Println("Loading ZTM Gdańsk stops data...")
	client := api.NewZTMClient()
	stopsData, err := client.LoadStops()
	if err != nil {
		fmt.Printf("Error loading stops: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded %d stops successfully!\n", len(stopsData.Stops))
	console := ui.NewConsoleUI(client)
	console.Stops = stopsData.Stops

	for {
		choice := console.ShowMainMenu()
		switch choice {
			case 1:
				console.HandleStopSearch()
			case 2:
				console.HandleParallelMonitoring()
			case 3:
				console.HandleShowRoutes()
			case 4:
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
		}
	}
}