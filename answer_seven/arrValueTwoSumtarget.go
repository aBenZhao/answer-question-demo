package answer_seven

// ArrValueTwoSumTarget 函数用于在整数数组中查找两个数的和等于目标值，并返回这两个数的索引
// 参数:
//
//	valueSlice: 整数切片，包含需要查找的数字
//	sumTarget: 目标和，需要找到的两个数的和等于此值
//
// 返回值:
//
//	[]int: 包含两个索引的切片，这两个索引对应的元素之和等于目标值
//	       如果未找到满足条件的元素，则返回 [0]
func ArrValueTwoSumTarget(valueSlice []int, sumTarget int) []int {
	// 遍历数组中的每个元素作为第一个数
	//for i := 0; i < len(valueSlice); i++ {
	//	first := valueSlice[i] // 获取第一个数
	//	// 遍历数组中的每个元素作为第二个数
	//	for j := 1; j < len(valueSlice); j++ {
	//		second := valueSlice[j] // 获取第二个数
	//		// 检查两个数的和是否等于目标值
	//		if sumTarget == first+second {
	//			// 如果找到满足条件的两个数，返回它们的索引
	//			return []int{i, j}
	//		}
	//	}
	//}

	firstPoint := 0
	for i := 1; i < len(valueSlice); i++ {
		first := valueSlice[firstPoint] // 获取第一个数
		second := valueSlice[i]         // 获取第二个数
		// 检查两个数的和是否等于目标值
		if sumTarget == first+second {
			// 如果找到满足条件的两个数，返回它们的索引
			return []int{firstPoint, i}
		}
		if i == len(valueSlice)-1 {
			firstPoint++
			i = firstPoint
		}
	}
	// 如果遍历完整个数组都没有找到满足条件的两个数，返回 [0]
	return []int{0}
}
