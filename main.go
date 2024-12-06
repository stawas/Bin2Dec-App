package main

import (
	"fmt"
	"math"
)

func main() {
	println("Binary to Decimal App")
	println("Enter Binary 1-8 length")
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	decimal, err := binaryToDecimal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("=================")

	fmt.Println("Binary: ", input)
	fmt.Println("Decimal: ", decimal)	
}

func binaryToDecimal(binary int) (int, error) {
	
	if (binary / 100_000_000) != 0 {
		return 0, fmt.Errorf(ERROR_NUMBER_OVER_8_LENGTH)
	}

	var decimal int = 0
	var digit int = 0
	var power float64 = 10

	for i := 0; i < 8; i++ {
		digit = int(math.Round(math.Mod(float64(binary)/power, 1) * 10))
		if digit != 0 && digit != 1 {
			return 0, fmt.Errorf(ERROR_NUMBER_NOT_0_OR_1)
		}
		decimal += int(digit) * IntPow(2, i)
		power *= 10
	}

	return decimal, nil
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
