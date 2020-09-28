package problem713

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans := 0
	left := 0
	prod := 1
	for i, num := range nums {
		prod *= num
		for prod >= k {
			prod /= nums[left]
			left++
		}
		ans += i - left + 1
	}
	return ans
}
