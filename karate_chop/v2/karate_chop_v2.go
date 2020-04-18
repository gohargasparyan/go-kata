package v1

func Chop(findMe int, numbers []int) int {
	return chop(findMe, numbers, 0, len(numbers) - 1)
}

func chop(findMe int, numbers []int, startIdx, endIdx int) int {
	if startIdx < endIdx {
		median := (startIdx + endIdx) / 2
		if numbers[median] < findMe {
			return chop(findMe, numbers, median + 1, endIdx)
		} else if numbers[median] > findMe {
			return chop(findMe, numbers, startIdx, median - 1)
		} else {
			return checkIsAtIdx(numbers, median, findMe)
		}
	}

	return checkIsAtIdx(numbers, startIdx, findMe)
}

func checkIsAtIdx(numbers []int, idx int, findMe int) int {
	if idx == len(numbers) || numbers[idx] != findMe {
		return -1
	} else {
		return idx
	}
}


