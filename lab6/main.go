package main

import (
	"fmt"
	"zad6/ui"
)

func main() {
  console := ui.NewConsole()
    
  fmt.Println("Starting Stock Market Analysis Tool...")
  console.Start()
}
