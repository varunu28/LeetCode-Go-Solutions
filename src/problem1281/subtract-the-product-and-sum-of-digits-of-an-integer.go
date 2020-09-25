package problem1281

func subtractProductAndSum(n int) int {
	if n == 0 {
		return 0
	}
	prod := 1
	sum := 0
	for n > 0 {
		rem := n % 10
		prod *= rem
		sum += rem
		n /= 10
	}
	return prod - sum
}
