package problem229

func majorityElement(nums []int) []int {
	countOne := 0
	countTwo := 0
	candidateOne := -1
	candidateTwo := -1
	for _, num := range nums {
		if candidateOne == num {
			countOne++
		} else if candidateTwo == num {
			countTwo++
		} else if countOne == 0 {
			candidateOne = num
			countOne++
		} else if countTwo == 0 {
			candidateTwo = num
			countTwo++
		} else {
			countOne--
			countTwo--
		}
	}
	result := []int{}
	countOne = 0
	countTwo = 0
	for _, num := range nums {
		if candidateOne == num {
			countOne++
		}
		if candidateTwo == num {
			countTwo++
		}
	}
	if countOne > len(nums)/3 {
		result = append(result, candidateOne)
	}
	if countTwo > len(nums)/3 {
		result = append(result, candidateTwo)
	}
	return result
}
