package problem1313

func decompressRLElist(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i += 2 {
		val := nums[i+1]
		count := nums[i]
		for count > 0 {
			ans = append(ans, val)
			count--
		}
	}
	return ans
}
