package main

import (
	"fmt"
	"math"
)

const (
	USDToEUR = 0.96
	USDToRUB = 91.15
)

func main() {
	EURToRUB := USDToRUB / USDToEUR
	EURToRUB = math.Round(EURToRUB*100) / 100

	fmt.Println("Курс EUR к RUB:", EURToRUB)
}
