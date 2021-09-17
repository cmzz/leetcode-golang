package _78_first_bad_version

// firstBadVersion 查找第一个出现的
// **********##########  (10个*  10个#）
// 找到第一个 #
// 如果第15是#， 那么之后的都是 #
//		缩小右边界： right = 15
// 如果第6个不是#， 那么之前的都不是 #
//		缩小做辩解： left = mid + 1
// 直到  left == right == 10
func firstBadVersion(n int) int {
	l, h := 0, n
	for l < h {
		mid := l + (h-l)/2
		if isBadVersion(mid) {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func isBadVersion(n int) bool {
	return false
}
