package problem134

func canCompleteCircuit(gas []int, cost []int) int {
	totalTankCapacity := 0
	currTankCapacity := 0
	currTankIdx := 0
	for i := 0; i < len(gas); i++ {
		totalTankCapacity += gas[i] - cost[i]
		currTankCapacity += gas[i] - cost[i]
		if currTankCapacity < 0 {
			currTankIdx = i + 1
			currTankCapacity = 0
		}
	}
	if totalTankCapacity < 0 {
		return -1
	}
	return currTankIdx
}
