package v1

func Chop(findMe int, numbers []int) int {
	startIdx := 0
	endIdx := len(numbers) - 1

	for {
		if endIdx <= startIdx {
			break
		}
		median := (startIdx + endIdx) / 2
		if numbers[median] < findMe {
			startIdx = median + 1
		} else if numbers[median] > findMe {
			endIdx = median -1
		} else {
			return median
		}
	}

	return checkIsAtIdx(numbers, startIdx, findMe)
}

func checkIsAtIdx(numbers []int, startIdx int, findMe int) int {
	if len(numbers) == 0 || numbers[startIdx] != findMe {
		return -1
	}

	return startIdx
}
