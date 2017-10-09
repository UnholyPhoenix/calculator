package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	userInput := flag.String("c", "", "Calculates expression.")
	flag.Parse()

	// Splits input by space
	statment := strings.Fields(*userInput)

	// Expression parts
	firstNumber := statment[0]
	operator := statment[1]
	secondNumber := statment[2]

	firstN, err := strconv.ParseFloat(firstNumber, 64)
	if err != nil {
		fmt.Print("Prišlo je do napake. Vnesite veljaven izraz.\n")
		return
	}

	secondN, err := strconv.ParseFloat(secondNumber, 64)
	if err != nil {
		fmt.Print("Prišlo je do napake. Vnesite veljaven izraz.\n")
		return
	}

	switch {
	case operator == "+":
		fmt.Printf("\nRešitev: %.2f\n", firstN+secondN)
	case operator == "-":
		fmt.Println("-------------------------------------------")
		fmt.Printf("\nRešitev: %.2f\n", firstN-secondN)
	case operator == "*":
		fmt.Println("-------------------------------------------")
		fmt.Printf("\nRešitev: %.2f\n", firstN*secondN)
	case operator == "/" && secondN == 0.00:
		fmt.Println("-------------------------------------------")
		fmt.Println("\nRešitev: ni rešljivo")
	case operator == "/":
		fmt.Println("-------------------------------------------")
		fmt.Printf("\nRešitev: %.2f\n", firstN/secondN)
	default:
		fmt.Println("\nOperator je neveljaven. Dovoljeni operatorji so '+', '-', '*' in '/'")
	}

	fmt.Println("\n------------------------------------------------------")
	fmt.Println("Dovoljene operacije so '+', '-', '*' in '/'.          ")
	fmt.Println("Oblika izraza je sledeča: število operator število    ")
	fmt.Println("Primeri: 5 + 5, 5 - 5, 5 * 5 in 5 / 5                 ")
	fmt.Println("------------------------------------------------------")
}
