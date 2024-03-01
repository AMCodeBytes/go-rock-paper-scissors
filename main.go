package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int

	fmt.Println("-----   Rock, Paper, Scissors   -----")

	for {
		fmt.Println("1. Play")
		fmt.Println("2. Quit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// code
		case 2:
			os.Exit(0)
		default:
			panic("Invalid input entered.")
		}

		choice = 0
	}
}
