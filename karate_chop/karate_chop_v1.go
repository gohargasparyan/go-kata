package karate_chop

func Chop(findMe int, numbers []int) int {

	startIdx := 0
	endIdx := len(numbers) - 1

	for startIdx <= endIdx {
		median := (startIdx + endIdx) / 2

		if numbers[median] < findMe {
			startIdx = median + 1
		} else {
			endIdx = median - 1
		}
	}

	if startIdx == len(numbers) || numbers[startIdx] != findMe {
		return -1
	}

	return startIdx
}


