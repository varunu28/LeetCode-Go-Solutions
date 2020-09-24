package problem389

func findTheDifference(s string, t string) byte {
	counter := make([]int, 26)
	for _, c := range s {
		counter[int(c-'a')]++
	}
	for _, c := range t {
		counter[int(c-'a')]--
	}
	for i := range counter {
		if counter[i] < 0 {
			return byte(97 + i)
		}
	}
	return byte('-')
}
