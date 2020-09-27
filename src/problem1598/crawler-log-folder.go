package problem1598

func minOperations(logs []string) int {
	currPos := 0
	for _, log := range logs {
		if log == "../" {
			if currPos > 0 {
				currPos--
			}
		} else if log != "./" {
			currPos++
		}
	}
	return currPos
}
