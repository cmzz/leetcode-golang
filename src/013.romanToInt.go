package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	result := 0
	nums := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for true {
		roman := ""
		k := 0

		for i := 0; i < 13; i++ {
			t := romans[i]

			if strings.HasPrefix(s, t) {
				roman = t
				k = i
			}
		}

		result = result + nums[k]
		s = strings.Replace(s, roman, "", 1)

		if s == "" {
			break
		}
	}

	return result
}

func main() {
	rn := "III"
	fmt.Println(rn, romanToInt(rn))

	rn = "IV"
	fmt.Println(rn, romanToInt(rn))
}