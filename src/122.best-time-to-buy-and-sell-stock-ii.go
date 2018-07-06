package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max := 0

	for i := 0; i < len(prices) - 1; {
		j := i + 1

		tmax := prices[i]

		for ; j < len(prices) ; j++ {
			if prices[j] >= tmax {
				tmax = prices[j]
			} else {
				break
			}
		}

		if j >= (i + 1) && (j - 1) < len(prices) && prices[j - 1] > prices[i] {
			max = max + prices[j - 1] - prices[i]
			i = j
		} else {
			i++
		}
	}

	return max
}

func main() {
	var arr []int
	arr = []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(arr))

	arr = []int{1,2,3,4,5}
	fmt.Println(maxProfit(arr))

	arr = []int{7,6,4,3,1}
	fmt.Println(maxProfit(arr))
}