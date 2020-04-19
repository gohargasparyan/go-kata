package v1

func Chop(findMe int, numbers []int) int {
	return chop(findMe, numbers, 0)
}

func chop(findMe int, numbers []int, slicePositionInInitialArray int) int {
	length := len(numbers)
	if length > 1 {
		median := length / 2
		if numbers[median] < findMe {
			return chop(findMe, numbers[median:], slicePositionInInitialArray + median)
		} else if numbers[median] > findMe {
			return chop(findMe, numbers[:median], slicePositionInInitialArray)
		} else {
			return median + slicePositionInInitialArray
		}

	}

	return checkIsAtIdx(numbers, 0, findMe, slicePositionInInitialArray)
}

func checkIsAtIdx(numbers []int, idx int, findMe, slicePositionInInitialArray int) int {
	if idx == len(numbers) || numbers[idx] != findMe {
		return -1
	} else {
		return idx + slicePositionInInitialArray
	}
}


