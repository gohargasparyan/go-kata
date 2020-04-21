package v1

func Chop(findMe int, numbers []int) int {
	return chop(findMe, numbers, 0, len(numbers)-1)
}

func chop(findMe int, numbers []int, startIdx, endIdx int) int {
	if endIdx <= startIdx {
		return checkIsAtIdx(numbers, startIdx, findMe)
	}

	median := (startIdx + endIdx) / 2
	if numbers[median] == findMe {
		return median
	} else if numbers[median] < findMe {
		startIdx = median + 1
	} else {
		endIdx = median - 1
	}

	return chop(findMe, numbers, startIdx, endIdx)
}

func checkIsAtIdx(numbers []int, idx int, findMe int) int {
	if idx == len(numbers) || numbers[idx] != findMe {
		return -1
	} else {
		return idx
	}
}
