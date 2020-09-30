package problem41

func firstMissingPositive(nums []int) int {
	n := len(nums)
	idx := 0
	for idx < n {
		if nums[idx] > 0 && nums[idx] < n && nums[nums[idx]] != nums[idx] {
			swap(nums, idx, nums[idx])
		} else {
			idx++
		}
	}
	idx = 1
	for idx < n && nums[idx] == idx {
		idx++
	}
	if n == 0 || idx < n {
		return idx
	}
	if nums[0] == idx {
		return idx + 1
	}
	return idx
}

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
