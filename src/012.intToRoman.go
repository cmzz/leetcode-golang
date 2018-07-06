package main

import "fmt"

func intToRoman(num int) string {
	result := ""
	nums := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romans := [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for true {
		m := 1
		j := 0

		for i := 0; i < 13; i++  {
			if nums[i] > num {
				break
			}

			j = i
			m = nums[i]
		}

		if num / m == 0 {
			break
		}

		if num / m > 0 {
			for p := 0; p < num / m; p++ {
				result = result + romans[j]
			}

			num = num % m
		}
	}

	return result
}

func main() {
	n := 2
	fmt.Println(n, intToRoman(n))

	n = 3
	fmt.Println(n, intToRoman(n))

	n = 4
	fmt.Println(n, intToRoman(n))

	n = 9
	fmt.Println(n, intToRoman(n))

	n = 58
	fmt.Println(n, intToRoman(n))

	n = 1994
	fmt.Println(n, intToRoman(n))
}

