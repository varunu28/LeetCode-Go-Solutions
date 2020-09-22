package problem1512

func numIdenticalPairs(nums []int) int {
	counter := map[int]int{}
	pairs := 0
	for _, num := range nums {
		val, ok := counter[num]
		if ok {
			pairs += val
			counter[num] = counter[num] + 1
		} else {
			counter[num] = 1
		}
	}
	return pairs
}
