package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	result := 0
	if x == 0 {
		return 0
	}

	for x != 0 {
		result = result * 10 + x % 10
		x = x / 10
	}

	if math.Abs(float64(result)) > math.MaxInt32 {
		result = 0
	}

	return result
}

func main() {
	fmt.Println(reverse(-123))
}