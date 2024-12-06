package main

import (
	"fmt"
	"math"
)

func main() {
	println("Binary to Decimal App")
	println("Enter Binary 1-8 length")
	var input int
	var decimal int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	if (input / 100_000_000) != 0 {
		fmt.Println("Error: input over 8 length")
		return
	}

	fmt.Println("=================")
	var digit int = 0
	var power float64 = 10

	for i := 0; i < 8; i++ {
		digit = int(math.Round(math.Mod(float64(input)/power, 1) * 10))
		// fmt.Println("Digit: ", digit)
		if digit != 0 && digit != 1 {
			fmt.Println("Error: value must contains only 0 or 1")
			return
		}
		decimal += int(digit) * IntPow(2, i)
		power *= 10
	}

	fmt.Println("Binary: ", input)
	fmt.Println("Decimal: ", decimal)
}

func IntPow(x int, y int) int {
	if y == 0 {
		return 1
	}

	if y == 1 {
		return x
	}

	sum := x

	for i := 1; i < y; i++ {
		sum *= x
	}

	return sum
}

// Test case
// 45813275
// 123
// 10111011
