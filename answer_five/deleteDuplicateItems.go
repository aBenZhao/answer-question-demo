package answer_five

func DeleteDuplicateItems(numSlice []int) int {
	var slowPoint int = 0
	for i := 1; i < len(numSlice); i++ {
		var fast int = i
		if numSlice[slowPoint] != numSlice[fast] {
			slowPoint++
			numSlice[slowPoint] = numSlice[fast]
		}

	}
	return slowPoint + 1
}
