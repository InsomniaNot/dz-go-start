package main

import (
	"fmt"
	"math"
)

func main() {
	userHaigt := 1.8
	userKg := 100.0
	IMT := userKg / math.Pow(userHaigt, 2)
	fmt.Println(IMT)
}
