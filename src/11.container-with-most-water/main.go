package _1_container_with_most_water

// https://leetcode-cn.com/problems/container-with-most-water

func maxArea(height []int) int {
	var maxArea = 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			maxArea = getMaxAre(height[i], height[j], j-i, maxArea)
		}
	}

	return maxArea
}

func maxArea2(height []int) int {
	var maxArea = 0
	for i, j := 0, len(height)-1; i < j; {
		maxArea = getMaxAre(height[i], height[j], j-i, maxArea)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func getMaxAre(h1, h2, w, area int) int {
	return max(min(h1, h2)*w, area)
}

func min(m, n int) int {
	if m < n {
		return m
	}

	return n
}

func max(h1 int, h2 int) int {
	if h1 > h2 {
		return h1
	}

	return h2
}
