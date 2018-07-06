package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	max := 0

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices) ; j++ {
			if (prices[j] - prices[i]) > max {
				max = prices[j] - prices[i]
			}
		}
	}

	return max
}

func main() {
	var arr []int
	arr = []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(arr))

	arr = []int{7,6,4,3,1}
	fmt.Println(maxProfit(arr))

	arr = []int{1, 2}
	fmt.Println(maxProfit(arr))

	arr = []int{2, 4, 1}
	fmt.Println(maxProfit(arr))

	arr = []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(arr))

	arr = []int{2,1,2,0,1}
	fmt.Println(maxProfit(arr))

	arr = []int{2,1,2,1,0,1,2}
	fmt.Println(maxProfit(arr))

	arr = []int{3,2,6,5,0,3	}
	fmt.Println(maxProfit(arr))
}