package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_LINES int32 = 3
var MAX_BET int32 = 100
var MIN_BET int32 = 1

var ROWS int32 = 3
var COLS int32 = 3

var symbolCount = map[string]int32{
	"A": 2,
	"B": 4,
	"C": 6,
	"D": 8,
}

var symbolValue = map[string]int32{
	"A": 5,
	"B": 4,
	"C": 3,
	"D": 2,
}

func getSlotMachine() (columns [][]string) {
	var allSymbols []string
	for symbol := range symbolCount {
		allSymbols = append(allSymbols, symbol)
	}

	// var columns [][]string
	for i := 0; i < int(COLS); i++ {
		var column []string
		currentSymbols := allSymbols[:]
		for j := 0; j < int(ROWS); j++ {
			s := rand.NewSource(time.Now().Unix())
			r := rand.New(s) // initialize local pseudorandom generator
			in := r.Intn(len(currentSymbols))
			value := currentSymbols[in]

			currentSymbols[in] = currentSymbols[len(currentSymbols)-1] // Copy last element to index i.
			currentSymbols = currentSymbols[:len(currentSymbols)-1]
			column = append(column, value)
		}
		columns = append(columns, column)
	}

	fmt.Println(columns)
	return columns
}

func deposit() int32 {
	var amount int32
	for {
		fmt.Print("What would you like to deposit? $")
		fmt.Scan(&amount)

		if amount > 0 {
			break
		} else {
			fmt.Println("Amount must be greater than 0.")
		}
	}

	return amount
}

func getNumberOfLines() int32 {
	var lines int32
	for {
		fmt.Printf("Enter the number of lines to bet on (1-%d)?", MAX_LINES)
		fmt.Scan(&lines)

		if 1 <= lines && lines <= int32(MAX_LINES) {
			break
		} else {
			fmt.Println("Please enter a valid number of lines.")
		}
	}

	return lines
}

func getBet() int32 {
	var amount int32
	for {
		fmt.Print("What would you like to bet on each line? $")
		fmt.Scan(&amount)

		if MIN_BET <= amount && amount <= MAX_BET {
			break
		} else {
			fmt.Printf("Amount must be between %d - %d.", MIN_BET, MAX_BET)
		}
	}

	return amount
}

func main() {
	// getSlotMachine()
	balance := deposit()
	lines := getNumberOfLines()

	for {
		bet := getBet()
		totalBet := bet * lines

		if totalBet > balance {
			fmt.Printf("You do not have enough to bet that amount, your current balance is: %d", balance)
		} else {
			break
		}
		fmt.Printf("You are betting %d on %d lines. Total bet is equal to: %d", bet, lines, totalBet)
	}

}
