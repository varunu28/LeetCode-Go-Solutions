package problem495

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := len(timeSeries)
	if n == 0 || duration <= 0 {
		return 0
	}
	totalTime := 0
	currTime := 0
	for i := 0; i < n; i++ {
		if currTime > 0 {
			if currTime <= timeSeries[i] {
				totalTime += duration
			} else {
				totalTime += timeSeries[i] - timeSeries[i-1]
			}
		}
		currTime = timeSeries[i] + duration
	}
	return totalTime + currTime - timeSeries[n-1]
}
