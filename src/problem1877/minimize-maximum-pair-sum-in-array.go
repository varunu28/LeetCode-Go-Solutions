package problem1877

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	max_sum := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		curr_sum := nums[i] + nums[j]
		if curr_sum > max_sum {
			max_sum = curr_sum
		}
	}
	return max_sum
}
