package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

func main() {
	var choice int

	fmt.Println("-----   Rock, Paper, Scissors   -----")

	for {
		displayMenu()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			play()
		case 2:
			os.Exit(0)
		default:
			panic("Invalid input entered.")
		}

		choice = 0
	}
}

func play() {
	var playerChoice int
	var playerAction string
	var computerChoice int
	var computerAction string
	var winner string = ""
	var wg sync.WaitGroup

	wg.Add(1)
	go computerPlay(&computerChoice, &computerAction, &wg)

	fmt.Println("Computer Moving...")

	wg.Add(1)
	go playerPlay(&playerChoice, &playerAction, &wg)

	wg.Wait()

	fmt.Printf("Player - %v vs %v - Computer\n", playerAction, computerAction)

	if playerChoice == 1 && computerChoice == 3 {
		winner = "Winner!"
	} else if playerChoice == 2 && computerChoice == 1 {
		winner = "Winner!"
	} else if playerChoice == 3 && computerChoice == 2 {
		winner = "Winner!"
	} else if playerChoice == computerChoice {
		winner = "Draw!"
	} else {
		winner = "Loser!"
	}

	fmt.Printf("\n\n----- %v -----\n\n", winner)
}

func playerPlay(pc *int, pa *string, wg *sync.WaitGroup) {
	defer wg.Done()
	var input string

	fmt.Println("\n\n-----   Your Move   -----")
	fmt.Println("1. Rock")
	fmt.Println("2. Paper")
	fmt.Println("3. Scissors")

	fmt.Print("Move: ")
	fmt.Scan(&input)

	*pc, _ = strconv.Atoi(input)
	*pa = getAction(*pc)
}

func computerPlay(cc *int, ca *string, wg *sync.WaitGroup) {
	defer wg.Done()

	*cc = rand.Intn(3) + 1
	*ca = getAction(*cc)
}

func getAction(chose int) string {
	switch chose {
	case 1:
		return "Rock"
	case 2:
		return "Paper"
	case 3:
		return "Scissors"
	default:
		panic("Invalid input entered")
	}
}

func displayMenu() {
	fmt.Println("1. Play")
	fmt.Println("2. Quit")
}
