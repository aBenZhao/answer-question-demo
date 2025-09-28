package main

import (
	"answer-question-demo/answer_six"
	"fmt"
)

// 主函数，程序的入口点
func main() {
	//题目一：判断一个整数是否是回文数（考察：数字操作、条件判断、循环）
	//answer_one.RunPalindromic(0)

	//题目二：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效（字符串处理、栈的使用）
	//answer_two.RunValidParentheses("{}")

	// 题目三：编写一个函数来查找字符串数组中的最长公共前缀（字符串处理、循环嵌套）
	//prefix := answer_three.RunlongestCommonPrefix([]string{"ab", "a"})
	//fmt.Println(prefix)

	// 题目四：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
	//answer_four.IncrementIntegerArr([]int64{9, 9, 0, 1, 9})

	// 题目五：给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
	//fmt.Println(answer_five.DeleteDuplicateItems([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}))

	// 题目六：合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
	fmt.Println(answer_six.MergeIntervals([][]int{{4, 7}, {1, 4}}))
}
