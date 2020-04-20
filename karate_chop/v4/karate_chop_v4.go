package v1

import "fmt"

func Chop(findMe int, numbers []int) int {
	startIdx := 0
	endIdx := len(numbers) - 1
	i := 0

	for {
		if endIdx <= startIdx {
			break
		}
		i++

		median := (startIdx + endIdx) / 2

		if numbers[median] < findMe {
			startIdx = median + 1
		} else if numbers[median] > findMe {
			endIdx = median -1
		} else {
			fmt.Println(i)
			return median
		}
	}

	fmt.Println(i)

	if len(numbers) == 0 || numbers[startIdx] != findMe {
		return -1
	}

	return startIdx
}
