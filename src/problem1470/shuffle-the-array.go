package problem1470

func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	start := 0
	end := n
	for i := 0; i < 2*n; i += 2 {
		ans[i] = nums[start]
		ans[i+1] = nums[end]
		start++
		end++
	}
	return ans
}
