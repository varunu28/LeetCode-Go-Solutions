package problem1365

func smallerNumbersThanCurrent(nums []int) []int {
	counter := make([]int, 101)
	for _, num := range nums {
		counter[num]++
	}
	sum := 0
	preSum := make([]int, 101)
	for i := range preSum {
		preSum[i] = sum
		sum += counter[i]
	}
	ans := make([]int, len(nums))
	for i := range nums {
		ans[i] = preSum[nums[i]]
	}
	return ans
}
