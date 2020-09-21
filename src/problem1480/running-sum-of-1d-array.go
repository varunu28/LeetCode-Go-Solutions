package problem1480

func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		ans[i] = sum
	}
	return ans
}
