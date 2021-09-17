package _83_move_zeroes

func moveZeroes1(nums []int) {
	var zeroCount = 0
	var ans = make([]int, 0)

	// count zeros
	for _, v := range nums {
		if v == 0 {
			zeroCount++
		}
	}

	// make all the non-zero elements retain their original order
	for _, v := range nums {
		if v != 0 {
			ans = append(ans, v)
		}
	}

	// move all zero to the end
	for i := 0; i < zeroCount; i++ {
		ans = append(ans, 0)
	}

	for i, v := range ans {
		nums[i] = v
	}
}

func moveZeroes2(nums []int) {
	var j = 0
	for _, v := range nums {
		if v != 0 {
			nums[j] = v
			j++
		}
	}

	if j < len(nums) {
		for k := j; k < len(nums); k++ {
			nums[k] = 0
		}
	}
}

func moveZeroes3(nums []int) {
	var j = 0
	for i, v := range nums {
		if v != 0 {
			nums[j] = v
			j++
			if i > j {
				nums[i] = 0
			}
		}
	}
}
