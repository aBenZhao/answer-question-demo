package answer_five

// DeleteDuplicateItems 是一个删除切片中重复元素的函数
// 使用双指针法（快慢指针）实现原地删除重复元素
// 参数: numSlice - 一个整数切片，可能包含重复元素
// 返回值: 返回删除重复元素后的新切片的长度
func DeleteDuplicateItems(numSlice []int) int {
	// slowPoint 慢指针，指向当前不重复元素的最后一个位置
	var slowPoint int = 0
	// 使用快指针遍历整个切片
	for i := 1; i < len(numSlice); i++ {
		// fast 快指针，从第二个元素开始遍历
		var fast int = i
		// 如果快指针指向的元素与慢指针指向的元素不同
		if numSlice[slowPoint] != numSlice[fast] {
			// 慢指针向前移动一位
			slowPoint++
			// 将快指针指向的元素复制到慢指针的下一个位置
			numSlice[slowPoint] = numSlice[fast]
		}
	}
	// 返回新切片的长度（慢指针位置+1）
	return slowPoint + 1
}
