package answer_six

import "sort"

// MergeIntervals 合并重叠的区间
// 输入: doubleDimensionalArr - 一个包含多个区间的二维数组，每个区间由两个整数表示
// 输出: 返回合并后的区间数组
func MergeIntervals(doubleDimensionalArr [][]int) [][]int {
	// 可以先对区间数组按照区间的起始位置进行排序
	sort.Slice(doubleDimensionalArr, func(i, j int) bool {
		return doubleDimensionalArr[i][0] < doubleDimensionalArr[j][0]
	})
	// 然后使用一个切片来存储合并后的区间，遍历排序后的区间数组
	relustDoubleDimensionalArr := [][]int{doubleDimensionalArr[0]}
	for i := 1; i < len(doubleDimensionalArr); i++ {
		current := doubleDimensionalArr[i]
		last := relustDoubleDimensionalArr[len(relustDoubleDimensionalArr)-1]
		// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
		if current[0] <= last[1] {
			if current[1] >= last[1] {
				merge := []int{last[0], current[1]}
				relustDoubleDimensionalArr[len(relustDoubleDimensionalArr)-1] = merge
			} else {
				merge := []int{last[0], last[1]}
				relustDoubleDimensionalArr[len(relustDoubleDimensionalArr)-1] = merge
			}
		} else {
			// 如果没有重叠，则将当前区间添加到切片中。
			relustDoubleDimensionalArr = append(relustDoubleDimensionalArr, current)
		}
	}

	return relustDoubleDimensionalArr
}
