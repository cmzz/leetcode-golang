package main

import "fmt"

func generate(numRows int) [][]int {
	var result [][]int

	for i := 0; i< numRows; i++ {
		t := new([]int)

		for j := 0; j <= i; j++ {
			if j == 0 || j == i{
				*t = append(*t, 1)
			} else {
				*t = append(*t, result[i - 1][j - 1] + result[i - 1][j])
			}
		}

		result = append(result, *t)
	}

	return result
}

func main() {
	fmt.Println(generate(5))
}