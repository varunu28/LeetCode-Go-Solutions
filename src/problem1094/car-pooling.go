package problem1094

import "sort"

func carPooling(trips [][]int, capacity int) bool {
	costMap := map[int]int{}
	for _, trip := range trips {
		_, ok := costMap[trip[1]]
		if !ok {
			costMap[trip[1]] = 0
		}
		costMap[trip[1]] += trip[0]
		_, ok = costMap[trip[2]]
		if !ok {
			costMap[trip[2]] = 0
		}
		costMap[trip[2]] -= trip[0]
	}
	keys := make([]int, len(costMap))
	for k := range costMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		capacity -= costMap[k]
		if capacity < 0 {
			return false
		}
	}
	return true
}
