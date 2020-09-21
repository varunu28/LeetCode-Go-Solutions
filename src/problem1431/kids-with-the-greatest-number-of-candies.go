package problem1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	maxCandies := candies[0]
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}
	for i, candy := range candies {
		ans[i] = candy+extraCandies >= maxCandies
	}
	return ans
}
