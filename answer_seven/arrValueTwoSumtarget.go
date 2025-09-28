package answer_seven

func ArrValueTwoSumTarget(valueSlice []int, sumTarget int) []int {
	for i := 0; i < len(valueSlice); i++ {
		first := valueSlice[i]
		for j := 1; j < len(valueSlice); j++ {
			second := valueSlice[j]
			if sumTarget == first+second {
				return []int{i, j}
			}
		}
	}
	return []int{0}
}
