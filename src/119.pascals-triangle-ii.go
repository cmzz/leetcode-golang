package main

import "fmt"

/**
第n行的m个数可表示为C(n-1,m-1)（n-1下标，m-1上标），即为从n-1个不同元素中取m-1个元素的组合数。
组合数计算方法：C(n,m)=n!/[m!(n-m)!]
 */
func getRow(rowIndex int) []int {
	result := []int{}
	i := rowIndex + 1
	j := rowIndex + 1

	for j > 0 {
		s := 1

		for k := 1; k < j ; k++ {
			s = s * (i - k) / k
		}

		j--

		result = append(result, s)
	}

	return result
}

func main() {
	fmt.Println(getRow(3))
}