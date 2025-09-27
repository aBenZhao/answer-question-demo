package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
import (
	"answer-question-demo/answer_three"
	"fmt"
)

// 主函数，程序的入口点
func main() {
	//题目一：判断一个整数是否是回文数（考察：数字操作、条件判断、循环）
	//answer_one.RunPalindromic(0)

	//题目二：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效（字符串处理、栈的使用）
	//answer_two.RunValidParentheses("{}")

	// 题目三：编写一个函数来查找字符串数组中的最长公共前缀（字符串处理、循环嵌套）
	prefix := answer_three.RunlongestCommonPrefix([]string{"ab", "a"})
	fmt.Println(prefix)
}
