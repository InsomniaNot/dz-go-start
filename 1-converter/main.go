package main

import (
	"fmt"
)

const (
	USDToEUR = 0.96
	USDToRUB = 91.15
)

func main() {
	EURToRUB := USDToRUB / USDToEUR
	fmt.Printf("Курс EUR к RUB: %.2f\n", EURToRUB)
}
